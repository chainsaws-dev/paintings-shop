// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLCurrenciesSelect - получает список валют
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLCurrenciesSelect(page int, limit int, dbc *pgxpool.Pool) (CurrenciesResponse, error) {

	var result CurrenciesResponse
	result.Currencies = CurrenciesList{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".currencies;`

	row := dbc.QueryRow(context.Background(), sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {
		sqlreq = fmt.Sprintf(`SELECT 
								id, 
								rus_name, 
								eng_name, 
								iso_lat_3, 
								iso_dig, 
								symbol
							FROM 
								"references".currencies
							ORDER BY
								id
							LIMIT %v OFFSET %v;`, limit, offset)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(context.Background(), sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {

		var cur Currency

		err = rows.Scan(&cur.ID, &cur.RusName, &cur.EngName,
			&cur.LatISO, &cur.DigISO, &cur.Symbol)

		if err != nil {
			return result, err
		}

		result.Currencies = append(result.Currencies, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLSingleCurrencySelect - получает данные конкретной валюты
//
// Параметры:
//
// ID - номер валюты в базе данных
//
func PostgreSQLSingleCurrencySelect(ID int, dbc *pgxpool.Pool) (Currency, error) {

	var result Currency

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".currencies
			WHERE 
				id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, ID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows <= 0 {
		return result, ErrCurrencyNotFound
	}

	sqlreq = `SELECT 
					id, 
					rus_name, 
					eng_name, 
					iso_lat_3, 
					iso_dig, 
					symbol
				FROM 
					"references".currencies
				WHERE 
					id=$1
				ORDER BY
					id
				LIMIT 1;`

	row = dbc.QueryRow(context.Background(), sqlreq, ID)

	err = row.Scan(&result.ID, &result.RusName, &result.EngName,
		&result.LatISO, &result.DigISO, &result.Symbol)

	if err != nil {
		return result, err
	}

	return result, nil
}

// PostgreSQLCurrenciesChange - определяет существует ли данная валюта и вызывает
// INSERT или UPDATE в зависимости от результата проверки
func PostgreSQLCurrenciesChange(cp Currency, dbc *pgxpool.Pool) (Currency, error) {

	found, cp, err := PostgreSQLFindCurrency(cp, dbc)

	if err != nil {
		return cp, err
	}

	if found {
		cp, err = PostgreSQLCurrenciesUpdate(cp, dbc)
	} else {
		cp, err = PostgreSQLCurrenciesInsert(cp, dbc)
	}

	return cp, err
}

// PostgreSQLFindCurrency - ищет валюту по ID
func PostgreSQLFindCurrency(cp Currency, dbc *pgxpool.Pool) (bool, Currency, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM 
					"references".currencies 
				WHERE 
					id=$1;`

	CountRow := dbc.QueryRow(context.Background(), sqlreq, cp.ID)

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

// PostgreSQLCurrenciesInsert - добавляет новую валюту
func PostgreSQLCurrenciesInsert(cp Currency, dbc *pgxpool.Pool) (Currency, error) {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `INSERT INTO 
						"references".currencies(rus_name, eng_name, iso_lat_3, iso_dig, symbol)
						VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	row := dbc.QueryRow(context.Background(), sqlreq, cp.RusName, cp.EngName, cp.LatISO, cp.DigISO, cp.Symbol)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	cp.ID = curid

	log.Printf("Данные о валюте сохранены в базу данных под индексом %v", curid)

	dbc.Exec(context.Background(), "COMMIT")

	return cp, nil
}

// PostgreSQLCurrenciesUpdate - обновляет существующего контрагента
func PostgreSQLCurrenciesUpdate(cp Currency, dbc *pgxpool.Pool) (Currency, error) {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `UPDATE "references".currencies		
				SET (rus_name, eng_name, iso_lat_3, iso_dig, symbol) = ($1, $2, $3, $4, $5)
				WHERE id=$6;`

	_, err := dbc.Exec(context.Background(), sqlreq, cp.RusName, cp.EngName, cp.LatISO, cp.DigISO, cp.Symbol, cp.ID)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return cp, nil
}

// PostgreSQLCurrenciesDelete - удаляет валюту по номеру
func PostgreSQLCurrenciesDelete(ID int, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `DELETE FROM "references".currencies WHERE id=$1;`

	_, err := dbc.Exec(context.Background(), sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"references"."currencies_id_seq"',(select COALESCE(max("id"),1) from "references"."currencies")::bigint);`

	_, err = dbc.Exec(context.Background(), sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}
