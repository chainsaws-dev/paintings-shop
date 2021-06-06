// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLAuthorsSelect - получает список авторов
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLAuthorsSelect(page int, limit int, dbc *sql.DB) (AuthorsResponse, error) {

	var result AuthorsResponse
	result.ArtAuthors = AuthorsList{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".authors;`

	row := dbc.QueryRow(sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {
		sqlreq = fmt.Sprintf(`SELECT 
									authors.id, 
									authors.first_name, 
									authors.middle_name, 
									authors.last_name, 
									authors.bio, 
									authors.city, 
									authors.eng_name, 		
									authors.file_id, 
									files.filename,
									files.filesize,
									files.filetype,
									files.file_id AS picture_link,
									files.preview_id AS preview_link,		
									authors.country_id, 
									countries.name,
									countries.full_name,
									countries.english,
									countries.alpha_2,
									countries.alpha_3,
									countries.iso,
									countries.location,
									countries.location_precise,
									authors.user_id,
									users.role,
									users.email,
									users.phone,
									users.name,
									users.isadmin,
									users.confirmed,
									users.disabled,
									users.totp_active
							FROM 
									"references".authors
								LEFT JOIN
									"references".countries
								ON authors.country_id = countries.id
								LEFT JOIN
									"references".files
								ON authors.file_id = files.id
								LEFT JOIN
										secret.users 
								ON authors.user_id = users.id
							ORDER BY
								authors.id
							LIMIT %v OFFSET %v;`, limit, offset)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {

		var cur Author
		var f File
		var u User
		var c Country

		err = rows.Scan(&cur.ID, &cur.FirstName, &cur.MiddleName, &cur.LastName,
			&cur.Bio, &cur.OriginCity, &cur.EngName, &f.ID, &f.FileName, &f.FileSize,
			&f.FileType, &f.FileID, &f.PreviewID, &c.ID, &c.Name, &c.FullName, &c.English,
			&c.Alpha2, &c.Alpha3, &c.ISO, &c.Location, &c.LocationPrecise, &u.GUID, &u.Role,
			&u.Email, &u.Phone, &u.Name, &u.IsAdmin, &u.Confirmed, &u.Disabled, &u.SecondFactor)

		if err != nil {
			return result, err
		}

		cur.Photo = f
		cur.User = u
		cur.OriginCountry = c

		result.ArtAuthors = append(result.ArtAuthors, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLSingleAuthorSelect - получает данные конкретного автора
//
// Параметры:
//
// ID - номер автора в базе данных
//
func PostgreSQLSingleAuthorSelect(ID int, dbc *sql.DB) (Author, error) {

	var result Author
	result.Photo = File{}
	result.User = User{}
	result.OriginCountry = Country{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				"references".authors
			WHERE 
				id=$1;`

	row := dbc.QueryRow(sqlreq, ID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows <= 0 {
		return result, ErrAuthorNotFound
	}

	sqlreq = `SELECT 
				authors.id, 
				authors.first_name, 
				authors.middle_name, 
				authors.last_name, 
				authors.bio, 
				authors.city, 
				authors.eng_name, 		
				authors.file_id, 
				files.filename,
				files.filesize,
				files.filetype,
				files.file_id AS picture_link,
				files.preview_id AS preview_link,		
				authors.country_id, 
				countries.name,
				countries.full_name,
				countries.english,
				countries.alpha_2,
				countries.alpha_3,
				countries.iso,
				countries.location,
				countries.location_precise,
				authors.user_id,
				users.role,
				users.email,
				users.phone,
				users.name,
				users.isadmin,
				users.confirmed,
				users.disabled,
				users.totp_active
			FROM 
					"references".authors
				LEFT JOIN
					"references".countries
				ON authors.country_id = countries.id
				LEFT JOIN
					"references".files
				ON authors.file_id = files.id
				LEFT JOIN
						secret.users 
				ON authors.user_id = users.id
			WHERE 
				id=$1				
			ORDER BY
				authors.id
			LIMIT 1;`

	row = dbc.QueryRow(sqlreq, ID)

	var f File
	var u User
	var c Country

	err = row.Scan(&result.ID, &result.FirstName, &result.MiddleName, &result.LastName,
		&result.Bio, &result.OriginCity, &result.EngName, &f.ID, &f.FileName, &f.FileSize,
		&f.FileType, &f.FileID, &f.PreviewID, &c.ID, &c.Name, &c.FullName, &c.English,
		&c.Alpha2, &c.Alpha3, &c.ISO, &c.Location, &c.LocationPrecise, &u.GUID, &u.Role,
		&u.Email, &u.Phone, &u.Name, &u.IsAdmin, &u.Confirmed, &u.Disabled, &u.SecondFactor)

	if err != nil {
		return result, err
	}

	result.Photo = f
	result.User = u
	result.OriginCountry = c

	return result, nil
}

// PostgreSQLAuthorsChange - определяет существует ли данная страна и вызывает
// INSERT или UPDATE в зависимости от результата проверки
func PostgreSQLAuthorsChange(au Author, dbc *sql.DB) (Author, error) {

	found, au, err := PostgreSQLFindAuthor(au, dbc)

	if err != nil {
		return au, err
	}

	if found {
		au, err = PostgreSQLAuthorsUpdate(au, dbc)
	} else {
		au, err = PostgreSQLAuthorsInsert(au, dbc)
	}

	return au, err
}

// PostgreSQLFindAuthor - ищет страну по ID
func PostgreSQLFindAuthor(au Author, dbc *sql.DB) (bool, Author, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM 
					"references".authors 
				WHERE 
					id=$1;`

	CountRow := dbc.QueryRow(sqlreq, au.ID)

	var ItemsCount int
	err := CountRow.Scan(&ItemsCount)

	if err != nil {
		return false, au, err
	}

	if ItemsCount > 0 {
		return true, au, nil
	}

	return false, au, nil

}

// PostgreSQLAuthorsInsert - добавляет нового автора
func PostgreSQLAuthorsInsert(au Author, dbc *sql.DB) (Author, error) {

	dbc.Exec("BEGIN")

	sqlreq := `INSERT INTO 
						"references".authors(first_name, middle_name, last_name, bio, file_id, country_id, city, eng_name, user_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;`

	row := dbc.QueryRow(sqlreq, au.FirstName, au.MiddleName, au.LastName, au.Bio,
		au.Photo.FileID, au.OriginCountry.ID, au.OriginCity, au.EngName, au.User.GUID)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return au, PostgreSQLRollbackIfError(err, false, dbc)
	}

	au.ID = curid

	log.Printf("Данные об авторе сохранены в базу данных под индексом %v", curid)

	dbc.Exec("COMMIT")

	return au, nil
}

// PostgreSQLAuthorsUpdate - обновляет существующего автора
func PostgreSQLAuthorsUpdate(au Author, dbc *sql.DB) (Author, error) {

	dbc.Exec("BEGIN")

	sqlreq := `UPDATE "references".authors		
				SET (first_name, middle_name, last_name, bio, file_id, country_id, city, eng_name, user_id) = ($1, $2, $3, $4, $5, $6, $7, $8, $9)
				WHERE id=$10;`

	_, err := dbc.Exec(sqlreq, au.FirstName, au.MiddleName, au.LastName, au.Bio,
		au.Photo.FileID, au.OriginCountry.ID, au.OriginCity, au.EngName, au.User.GUID,
		au.ID)

	if err != nil {
		return au, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return au, nil
}

// PostgreSQLAuthorsDelete - удаляет страну по номеру
func PostgreSQLAuthorsDelete(ID int, dbc *sql.DB) error {

	dbc.Exec("BEGIN")

	sqlreq := `DELETE FROM "references".authors WHERE id=$1;`

	_, err := dbc.Exec(sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"references"."authors_id_seq"',(select COALESCE(max("id"),1) from "references"."authors")::bigint);`

	_, err = dbc.Exec(sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return nil
}
