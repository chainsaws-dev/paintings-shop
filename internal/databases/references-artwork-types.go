// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLArtworkTypesSelect - получает список типов картин
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLArtworkTypesSelect(page int, limit int, dbc *pgxpool.Pool) (ArtworkTypesResponse, error) {

	var result ArtworkTypesResponse
	result.ArtTypes = ArtworkTypesList{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".artwork_types;`

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
								name, 
								eng_name 
							FROM 
								"references".artwork_types
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

		var cur ArtworkType

		err = rows.Scan(&cur.ID, &cur.Name, &cur.EngName)

		if err != nil {
			return result, err
		}

		result.ArtTypes = append(result.ArtTypes, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLSingleArtworkTypeSelect - получает данные конкретного типа картин
//
// Параметры:
//
// ID - номер валюты в базе данных
//
func PostgreSQLSingleArtworkTypeSelect(ID int, dbc *pgxpool.Pool) (ArtworkType, error) {

	var result ArtworkType

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".artwork_types
			WHERE 
				id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, ID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows <= 0 {
		return result, ErrArtTypeNotFound
	}

	sqlreq = `SELECT 
					id, 
					name, 
					eng_name
				FROM 
					"references".artwork_types
				WHERE 
					id=$1
				ORDER BY
					id
				LIMIT 1;`

	row = dbc.QueryRow(context.Background(), sqlreq, ID)

	err = row.Scan(&result.ID, &result.Name, &result.EngName)

	if err != nil {
		return result, err
	}

	return result, nil
}

// PostgreSQLArtworkTypeChange - определяет существует ли данный тип картин и вызывает
// INSERT или UPDATE в зависимости от результата проверки
func PostgreSQLArtworkTypeChange(cp ArtworkType, dbc *pgxpool.Pool) (ArtworkType, error) {

	found, cp, err := PostgreSQLFindArtworkType(cp, dbc)

	if err != nil {
		return cp, err
	}

	if found {
		cp, err = PostgreSQLArtworkTypesUpdate(cp, dbc)
	} else {
		cp, err = PostgreSQLArtworkTypesInsert(cp, dbc)
	}

	return cp, err
}

// PostgreSQLFindArtworkType - ищет тип картин по ID
func PostgreSQLFindArtworkType(cp ArtworkType, dbc *pgxpool.Pool) (bool, ArtworkType, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM 
					"references".artwork_types 
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
func PostgreSQLArtworkTypesInsert(cp ArtworkType, dbc *pgxpool.Pool) (ArtworkType, error) {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `INSERT INTO 
						"references".artwork_types(name, eng_name)
						VALUES ($1, $2) RETURNING id;`

	row := dbc.QueryRow(context.Background(), sqlreq, cp.Name, cp.EngName)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	cp.ID = curid

	log.Printf("Данные о типе картины сохранены в базу данных под индексом %v", curid)

	dbc.Exec(context.Background(), "COMMIT")

	return cp, nil
}

// PostgreSQLCurrenciesUpdate - обновляет существующего контрагента
func PostgreSQLArtworkTypesUpdate(cp ArtworkType, dbc *pgxpool.Pool) (ArtworkType, error) {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `UPDATE "references".artwork_types		
				SET (name, eng_name) = ($1, $2)
				WHERE id=$3;`

	_, err := dbc.Exec(context.Background(), sqlreq, cp.Name, cp.EngName, cp.ID)

	if err != nil {
		return cp, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return cp, nil
}

// PostgreSQLCurrenciesDelete - удаляет валюту по номеру
func PostgreSQLArtworkTypesDelete(ID int, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `DELETE FROM "references".artwork_types WHERE id=$1;`

	_, err := dbc.Exec(context.Background(), sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"references"."artwork_types_id_seq"',(select COALESCE(max("id"),1) from "references"."artwork_types")::bigint);`

	_, err = dbc.Exec(context.Background(), sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}
