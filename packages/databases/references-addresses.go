// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLAddressesSelect - получает список адресов
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLAddressesSelect(page int, limit int, dbc *sql.DB) (AddressesResponse, error) {

	var result AddressesResponse
	result.Addresses = AddressesList{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".addresses;`

	row := dbc.QueryRow(sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {
		sqlreq = fmt.Sprintf(`SELECT 
								addresses.id,
								addresses.index,
								addresses.city,
								addresses.district,
								addresses.street,
								addresses.name,
								countries.id,
								countries.name,
								countries.full_name,
								countries.english,
								countries.alpha_2,
								countries.alpha_3,
								countries.iso,
								countries.location,
								countries.location_precise
							FROM 
								"references".addresses
							LEFT JOIN
								"references".countries
							ON addresses.country_id=countries.id
							ORDER BY
								addresses.id
							LIMIT %v OFFSET %v;`, limit, offset)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {

		var cur Address

		var c Country

		err = rows.Scan(&cur.ID, &cur.Index, &cur.City,
			&cur.District, &cur.Street, &cur.Name,
			&c.ID, &c.Name, &c.FullName, &c.English,
			&c.Alpha2, &c.Alpha3, &c.ISO, &c.Location, &c.LocationPrecise)

		if err != nil {
			return result, err
		}

		cur.Country = c

		result.Addresses = append(result.Addresses, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLSingleAddressSelect - получает данные конкретного адреса
//
// Параметры:
//
// ID - номер валюты в базе данных
//
func PostgreSQLSingleAddressSelect(ID int, dbc *sql.DB) (Address, error) {

	var result Address

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".addresses
			WHERE 
				id=$1;`

	row := dbc.QueryRow(sqlreq, ID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows <= 0 {
		return result, ErrAddressNotFound
	}

	sqlreq = `SELECT 
					addresses.id,
					addresses.index,
					addresses.city,
					addresses.district,
					addresses.street,
					addresses.name,
					countries.id,
					countries.name,
					countries.full_name,
					countries.english,
					countries.alpha_2,
					countries.alpha_3,
					countries.iso,
					countries.location,
					countries.location_precise
				FROM 
					"references".addresses
				LEFT JOIN
					"references".countries
				ON addresses.country_id=countries.id
				WHERE 
					addresses.id=$1
				ORDER BY
					addresses.id
				LIMIT 1;`

	row = dbc.QueryRow(sqlreq, ID)

	var c Country

	err = row.Scan(&result.ID, &result.Index, &result.City,
		&result.District, &result.Street, &result.Name, &c.ID,
		&c.Name, &c.FullName, &c.English, &c.Alpha2, &c.Alpha3,
		&c.ISO, &c.Location, &c.LocationPrecise)

	if err != nil {
		return result, err
	}

	result.Country = c

	return result, nil
}

// PostgreSQLCurrenciesChange - определяет существует ли данный адрес и вызывает
// INSERT или UPDATE в зависимости от результата проверки
func PostgreSQLAddressChange(cp Address, dbc *sql.DB) (Address, error) {

	found, cp, err := PostgreSQLFindAddress(cp, dbc)

	if err != nil {
		return cp, err
	}

	if found {
		cp, err = PostgreSQLAddressesUpdate(cp, dbc)
	} else {
		cp, err = PostgreSQLAddressesInsert(cp, dbc)
	}

	return cp, err
}

// PostgreSQLFindAddress - ищет адрес по ID
func PostgreSQLFindAddress(cp Address, dbc *sql.DB) (bool, Address, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM 
					"references".addresses 
				WHERE 
					id=$1;`

	CountRow := dbc.QueryRow(sqlreq, cp.ID)

	var ItemsCount int
	err := CountRow.Scan(&ItemsCount)

	if err != nil {
		return false, cp, err
	}

	if ItemsCount > 0 {
		return true, cp, nil
	}

	return false, cp, nil

}

// PostgreSQLAddressesInsert - добавляет новый адрес
func PostgreSQLAddressesInsert(cp Address, dbc *sql.DB) (Address, error) {

	dbc.Exec("BEGIN")

	sqlreq := `INSERT INTO 
						"references".addresses(index, country_id, city, district, street, name)
						VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	row := dbc.QueryRow(sqlreq, cp.Index, cp.Country.ID, cp.City, cp.District, cp.Street, cp.Name)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	cp.ID = curid

	log.Printf("Данные о адресе сохранены в базу данных под индексом %v", curid)

	dbc.Exec("COMMIT")

	return cp, nil
}

// PostgreSQLCurrenciesUpdate - обновляет существующего контрагента
func PostgreSQLAddressesUpdate(cp Address, dbc *sql.DB) (Address, error) {

	dbc.Exec("BEGIN")

	sqlreq := `UPDATE "references".addresses		
				SET (index, country_id, city, district, street, name) = ($1, $2, $3, $4, $5, $6)
				WHERE id=$6;`

	_, err := dbc.Exec(sqlreq, cp.Index, cp.Country.ID, cp.City, cp.District, cp.Street, cp.Name, cp.ID)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return cp, nil
}

// PostgreSQLCurrenciesDelete - удаляет валюту по номеру
func PostgreSQLAddressesDelete(ID int, dbc *sql.DB) error {

	dbc.Exec("BEGIN")

	sqlreq := `DELETE FROM "references".addresses WHERE id=$1;`

	_, err := dbc.Exec(sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"references"."addresses_id_seq"',(select COALESCE(max("id"),1) from "references"."addresses")::bigint);`

	_, err = dbc.Exec(sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return nil
}
