// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLTermsSelect - получает список условий продажи
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLTermsSelect(page int, limit int, dbc *sql.DB) (TermsResponse, error) {

	var result TermsResponse
	result.SaleTerms = Terms{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".terms;`

	row := dbc.QueryRow(sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {
		sqlreq = fmt.Sprintf(`SELECT 
									terms.id, 
									terms.delivery_time, 
									terms.returns, 
									terms.delivery_cost, 
									terms.name,
									terms.currency_id,
									currencies.rus_name,
									currencies.eng_name,
									currencies.iso_lat_3,
									currencies.iso_dig,
									currencies.symbol
								FROM 
									"references".terms
								LEFT JOIN
									"references".currencies
								ON 
									terms.currency_id=currencies.id	
							ORDER BY
								terms.id
							LIMIT %v OFFSET %v;`, limit, offset)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {

		var cur Term
		var curr Currency

		err = rows.Scan(&cur.ID, &cur.DeliveryTime, &cur.Returns,
			&cur.DeliveryCost, &cur.Name, &curr.ID, &curr.RusName,
			&curr.EngName, &curr.LatISO, &curr.DigISO, &curr.Symbol)

		if err != nil {
			return result, err
		}

		cur.Currency = curr

		result.SaleTerms = append(result.SaleTerms, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLSingleTermSelect - получает данные конкретных условий продажи
//
// Параметры:
//
// ID - номер условий в базе данных
//
func PostgreSQLSingleTermSelect(ID int, dbc *sql.DB) (Term, error) {

	var result Term
	result.Currency = Currency{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".terms
			WHERE 
				id=$1;`

	row := dbc.QueryRow(sqlreq, ID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows <= 0 {
		return result, ErrTermNotFound
	}

	sqlreq = `SELECT 
					terms.id, 
					terms.delivery_time, 
					terms.returns, 
					terms.delivery_cost, 
					terms.name,
					terms.currency_id,
					currencies.rus_name,
					currencies.eng_name,
					currencies.iso_lat_3,
					currencies.iso_dig,
					currencies.symbol
				FROM 
					"references".terms
				LEFT JOIN
					"references".currencies
				ON 
					terms.currency_id=currencies.id					
				WHERE 
					id=$1
				ORDER BY
					terms.id
				LIMIT 1;`

	row = dbc.QueryRow(sqlreq, ID)

	var curr Currency

	err = row.Scan(&result.ID, &result.DeliveryTime, &result.Returns,
		&result.DeliveryCost, &result.Name, &curr.ID, &curr.RusName,
		&curr.EngName, &curr.LatISO, &curr.DigISO, &curr.Symbol)

	if err != nil {
		return result, err
	}

	result.Currency = curr

	return result, nil
}

// PostgreSQLTermsChange - определяет существует ли данные условия и вызывает
// INSERT или UPDATE в зависимости от результата проверки
func PostgreSQLTermsChange(t Term, dbc *sql.DB) (Term, error) {

	found, t, err := PostgreSQLFindTerm(t, dbc)

	if err != nil {
		return t, err
	}

	if found {
		t, err = PostgreSQLTermsUpdate(t, dbc)
	} else {
		t, err = PostgreSQLTermsInsert(t, dbc)
	}

	return t, err
}

// PostgreSQLFindTerm - ищет условия продажи по ID
func PostgreSQLFindTerm(t Term, dbc *sql.DB) (bool, Term, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM 
					"references".terms 
				WHERE 
					id=$1;`

	CountRow := dbc.QueryRow(sqlreq, t.ID)

	var ItemsCount int
	err := CountRow.Scan(&ItemsCount)

	if err != nil {
		return false, t, err
	}

	if ItemsCount > 0 {
		return true, t, nil
	}

	return false, t, nil

}

// PostgreSQLTermsInsert - добавляет новые условия продажи
func PostgreSQLTermsInsert(t Term, dbc *sql.DB) (Term, error) {

	dbc.Exec("BEGIN")

	sqlreq := `INSERT INTO 
						"references".terms(delivery_time, returns, delivery_cost, name, currency_id)
						VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	row := dbc.QueryRow(sqlreq, t.DeliveryTime, t.Returns, t.DeliveryCost, t.Name, t.Currency.ID)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return t, PostgreSQLRollbackIfError(err, false, dbc)
	}

	t.ID = curid

	log.Printf("Данные о валюте сохранены в базу данных под индексом %v", curid)

	dbc.Exec("COMMIT")

	return t, nil
}

// PostgreSQLCurrenciesUpdate - обновляет существующего контрагента
func PostgreSQLTermsUpdate(t Term, dbc *sql.DB) (Term, error) {

	dbc.Exec("BEGIN")

	sqlreq := `UPDATE "references".terms		
				SET (delivery_time, returns, delivery_cost, name, currency_id) = ($1, $2, $3, $4, $5)
				WHERE id=$6;`

	_, err := dbc.Exec(sqlreq, t.DeliveryTime, t.Returns, t.DeliveryCost, t.Name, t.Currency.ID, t.ID)

	if err != nil {
		return t, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return t, nil
}

// PostgreSQLCurrenciesDelete - удаляет валюту по номеру
func PostgreSQLTermsDelete(ID int, dbc *sql.DB) error {

	dbc.Exec("BEGIN")

	sqlreq := `DELETE FROM "references".terms WHERE id=$1;`

	_, err := dbc.Exec(sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"references"."terms_id_seq"',(select COALESCE(max("id"),1) from "references"."terms")::bigint);`

	_, err = dbc.Exec(sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return nil
}
