// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// region * выборка * контрагенты *

// PostgreSQLCounterpartiesSelect - получает список контрагентов
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLCountriesSelect(page int, limit int, dbc *sql.DB) (CountriesResponse, error) {

	var result CountriesResponse
	result.Countries = CountriesList{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".countries;`

	row := dbc.QueryRow(sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {
		sqlreq = fmt.Sprintf(`SELECT 
									id, 
									name,
									full_name, 
									english, 
									alpha_2, 
									alpha_3, 
									iso, 
									location, 
									location_precise
							FROM 
								"references".countries
								LIMIT %v OFFSET %v;`, limit, offset)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {

		var cur Country

		err = rows.Scan(&cur.ID, &cur.Name, &cur.FullName, &cur.English,
			&cur.Alpha2, &cur.Alpha3, &cur.ISO, &cur.Location, &cur.LocationPrecise)

		if err != nil {
			return result, err
		}

		result.Countries = append(result.Countries, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLSingleCounterpartySelect - получает данные конкретного контрагента
//
// Параметры:
//
// ID - номер запроса налогвика в базе данных
//
func PostgreSQLSingleCountrySelect(ID int, dbc *sql.DB) (Country, error) {

	var result Country

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".countries
			WHERE 
				id=$1;`

	row := dbc.QueryRow(sqlreq, ID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows <= 0 {
		return result, ErrContryNotFound
	}

	sqlreq = `SELECT 
					id, 
					name,
					full_name, 
					english, 
					alpha_2, 
					alpha_3, 
					iso, 
					location, 
					location_precise
				FROM 
					"references".countries
				WHERE 
					id=$1
				LIMIT 1;`

	row = dbc.QueryRow(sqlreq, ID)

	err = row.Scan(&result.ID, &result.Name, &result.FullName,
		&result.English, &result.Alpha2, &result.Alpha3,
		&result.ISO, &result.Location, &result.LocationPrecise)

	if err != nil {
		return result, err
	}

	return result, nil
}

// endregion

// region * изменение * контрагенты *

// PostgreSQLCountriesChange - определяет существует ли данный контрагент и вызывает
// INSERT или UPDATE в зависимости от результата проверки
func PostgreSQLCountriesChange(cp Country, dbc *sql.DB) (Country, error) {

	found, cp, err := PostgreSQLFindCountry(cp, dbc)

	if err != nil {
		return cp, err
	}

	if found {
		cp, err = PostgreSQLCountriesUpdate(cp, dbc)
	} else {
		cp, err = PostgreSQLCountriesInsert(cp, dbc)
	}

	return cp, err
}

// PostgreSQLFindCountry - ищет контрагента по ID, а затем по ИНН и КПП
func PostgreSQLFindCountry(cp Country, dbc *sql.DB) (bool, Country, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM 
					"references".countries 
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

// PostgreSQLCounterpartiesInsert - добавляет нового контрагента
func PostgreSQLCountriesInsert(cp Country, dbc *sql.DB) (Country, error) {

	dbc.Exec("BEGIN")

	sqlreq := `INSERT INTO 
						"references".countries(name, full_name, english, alpha_2, alpha_3, iso, location, location_precise)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`

	row := dbc.QueryRow(sqlreq, cp.Name, cp.FullName, cp.English,
		cp.Alpha2, cp.Alpha3, cp.ISO, cp.Location, cp.LocationPrecise)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	cp.ID = curid

	log.Printf("Данные о стране сохранены в базу данных под индексом %v", curid)

	dbc.Exec("COMMIT")

	return cp, nil
}

// PostgreSQLCounterpartiesUpdate - обновляет существующего контрагента
func PostgreSQLCountriesUpdate(cp Country, dbc *sql.DB) (Country, error) {

	dbc.Exec("BEGIN")

	sqlreq := `UPDATE "references".countries		
				SET (name, full_name, english, alpha_2, alpha_3, iso, location, location_precise) = ($1, $2, $3, $4, $5, $6, $7, $8)
				WHERE id=$9;`

	_, err := dbc.Exec(sqlreq, cp.Name, cp.FullName, cp.English,
		cp.Alpha2, cp.Alpha3, cp.ISO,
		cp.Location, cp.LocationPrecise, cp.ID)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return cp, nil
}

// endregion

// region * удаление * контрагенты *

// PostgreSQLCountriesDelete - удаляет контрагента по номеру
func PostgreSQLCountriesDelete(ID int, dbc *sql.DB) error {

	sqlreq := `SELECT 
					COUNT(*) 
				FROM 
					"references".addresses
				WHERE 
					country_id=$1;`

	CountRow := dbc.QueryRow(sqlreq, ID)

	var ItemsCount int
	err := CountRow.Scan(&ItemsCount)

	if err != nil {
		return err
	}

	if ItemsCount > 0 {
		return ErrNoDeleteIfLinksExist
	}

	dbc.Exec("BEGIN")

	sqlreq = `DELETE FROM "references".countries WHERE id=$1;`

	_, err = dbc.Exec(sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"references"."countries_id_seq"',(select COALESCE(max("id"),1) from "references"."countries")::bigint);`

	_, err = dbc.Exec(sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return nil
}

// endregion
