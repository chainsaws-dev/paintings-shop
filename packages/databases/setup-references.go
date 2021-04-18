// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import "database/sql"

// PostgreSQLCreateTablesReferences - создаёт таблицы для схемы references
func PostgreSQLCreateTablesReferences(dbc *sql.DB) {

	// Рецепты и список покупок

	var CreateStatements = NamedCreateStatements{
		NamedCreateStatement{
			TableName: "files",
			CreateStatement: `CREATE TABLE "references".files
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
			
			ALTER TABLE "references".files
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "countries",
			CreateStatement: `CREATE TABLE "references".countries
			(
				id bigserial NOT NULL,
				name character varying(60) COLLATE pg_catalog."default",
				full_name character varying(60) COLLATE pg_catalog."default",
				english character varying(100) COLLATE pg_catalog."default",
				alpha_2 character varying(2) COLLATE pg_catalog."default",
				alpha_3 character varying(3) COLLATE pg_catalog."default",
				iso character varying(3) COLLATE pg_catalog."default",
				location character varying(20) COLLATE pg_catalog."default",
				location_precise character varying(30) COLLATE pg_catalog."default",
				CONSTRAINT countries_pkey PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".countries
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "addresses",
			CreateStatement: `CREATE TABLE "references".addresses
			(
				id bigserial NOT NULL,
				index character varying(6) COLLATE pg_catalog."default",
				country_id bigint,
				city character varying(100) COLLATE pg_catalog."default",
				district character varying(100) COLLATE pg_catalog."default",
				street character varying(100) COLLATE pg_catalog."default",
				name character varying(200) COLLATE pg_catalog."default",
				CONSTRAINT addresses_pkey PRIMARY KEY (id),
				CONSTRAINT addresses_country_id_fkey FOREIGN KEY (country_id)
					REFERENCES "references".countries (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".addresses
				OWNER to postgres;
			
			CREATE INDEX fki_addresses_country_id_fkey
				ON "references".addresses USING btree
				(country_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "authors",
			CreateStatement: `CREATE TABLE "references".authors
			(
				id bigserial NOT NULL,
				first_name character varying(200) COLLATE pg_catalog."default",
				middle_name character varying(200) COLLATE pg_catalog."default",
				last_name character varying(200) COLLATE pg_catalog."default",
				bio text COLLATE pg_catalog."default",
				file_id bigint,
				country_id bigint,
				city character varying(100) COLLATE pg_catalog."default",
				eng_name character varying(200) COLLATE pg_catalog."default",
				user_id uuid NOT NULL,
				CONSTRAINT authors_pkey PRIMARY KEY (id),
				CONSTRAINT authors_country_id_fkey FOREIGN KEY (country_id)
					REFERENCES "references".countries (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT authors_file_id_fkey FOREIGN KEY (file_id)
					REFERENCES "references".files (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT authors_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".authors
				OWNER to postgres;
			
			CREATE INDEX fki_authors_country_id_fkey
				ON "references".authors USING btree
				(country_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_authors_file_id_fkey
				ON "references".authors USING btree
				(file_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_authors_user_id_fkey
				ON "references".authors USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "artwork_types",
			CreateStatement: `CREATE TABLE "references".artwork_types
			(
				id bigserial NOT NULL,
				name character varying(50) COLLATE pg_catalog."default",
				eng_name character varying(50) COLLATE pg_catalog."default",
				CONSTRAINT artwork_types_pkey PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".artwork_types
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "currencies",
			CreateStatement: `CREATE TABLE "references".currencies
			(
				id bigserial NOT NULL,
				rus_name character varying(50) COLLATE pg_catalog."default",
				eng_name character varying(50) COLLATE pg_catalog."default",
				iso_lat_3 character varying(3) COLLATE pg_catalog."default",
				iso_dig character varying(3) COLLATE pg_catalog."default",
				CONSTRAINT currencies_pkey PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".currencies
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "terms",
			CreateStatement: `CREATE TABLE "references".terms
			(
				id bigserial NOT NULL,
				delivery_time text COLLATE pg_catalog."default",
				returns text COLLATE pg_catalog."default",
				delivery_cost text COLLATE pg_catalog."default",
				name character varying(100) COLLATE pg_catalog."default",
				currency_id bigint NOT NULL,
				CONSTRAINT terms_pkey PRIMARY KEY (id),
				CONSTRAINT terms_currency_id_fkey FOREIGN KEY (currency_id)
					REFERENCES "references".currencies (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".terms
				OWNER to postgres;
			
			CREATE INDEX fki_terms_currency_id_fkey
				ON "references".terms USING btree
				(currency_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
	}

	for _, ncs := range CreateStatements {
		PostgreSQLExecuteCreateStatement(dbc, ncs)
	}

}
