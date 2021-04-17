// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import "database/sql"

// PostgreSQLCreateTablesPublic - создаёт таблицы для схемы public (для рецептов и списка покупок)
func PostgreSQLCreateTablesPublic(dbc *sql.DB) {

	// Рецепты и список покупок

	var CreateStatements = NamedCreateStatements{
		NamedCreateStatement{
			TableName: "Files",
			CreateStatement: `CREATE TABLE public."Files"
			(
				id bigserial NOT NULL,
				filename character varying(255) COLLATE pg_catalog."default",
				filesize bigint,
				filetype character varying(50) COLLATE pg_catalog."default",
				file_id character varying(50) COLLATE pg_catalog."default",
				preview_id character varying(50) COLLATE pg_catalog."default",
				CONSTRAINT "Files_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."Files"
				OWNER to postgres;`,
		},
	}

	for _, ncs := range CreateStatements {
		PostgreSQLExecuteCreateStatement(dbc, ncs)
	}

}
