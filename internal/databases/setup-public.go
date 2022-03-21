// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import "github.com/jackc/pgx/v4/pgxpool"

// PostgreSQLCreateTablesPublic - создаёт таблицы для схемы public
func PostgreSQLCreateTablesPublic(dbc *pgxpool.Pool) {

	// Рецепты и список покупок

	var CreateStatements = NamedCreateStatements{
		NamedCreateStatement{
			TableName: "artworks",
			CreateStatement: `
			CREATE TABLE public.artworks
			(
				id bigserial NOT NULL,
				name character varying(100) COLLATE pg_catalog."default",
				description text COLLATE pg_catalog."default",
				file_id bigint,
				date_created timestamp with time zone,
				price bigint NOT NULL,
				currency_id bigint NOT NULL,
				type_id bigint NOT NULL,
				width double precision,
				height double precision,
				depth double precision,
				author_id bigint NOT NULL,
				terms_id bigint NOT NULL,				
				CONSTRAINT artworks_pkey PRIMARY KEY (id),
				CONSTRAINT artworks_author_id_fkey FOREIGN KEY (author_id)
					REFERENCES "references".authors (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT artworks_currency_id_fkey FOREIGN KEY (currency_id)
					REFERENCES "references".currencies (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT artworks_file_id_fkey FOREIGN KEY (file_id)
					REFERENCES "references".files (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT artworks_terms_id_fkey FOREIGN KEY (terms_id)
					REFERENCES "references".terms (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT atworks_type_id_fkey FOREIGN KEY (type_id)
					REFERENCES "references".artwork_types (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public.artworks
				OWNER to postgres;
			
			CREATE INDEX fki_artwork_file_id_fkey
				ON public.artworks USING btree
				(file_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_artworks_author_id_fkey
				ON public.artworks USING btree
				(author_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_artworks_currency_id_fkey
				ON public.artworks USING btree
				(currency_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_artworks_terms_id_fkey
				ON public.artworks USING btree
				(terms_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_atworks_type_id_fkey
				ON public.artworks USING btree
				(type_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
	}

	for _, ncs := range CreateStatements {
		PostgreSQLExecuteCreateStatement(dbc, ncs)
	}

}
