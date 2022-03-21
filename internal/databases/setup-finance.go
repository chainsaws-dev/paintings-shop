// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import "github.com/jackc/pgx/v4/pgxpool"

// PostgreSQLCreateTablesFinance - создаёт таблицы для схемы finance
func PostgreSQLCreateTablesFinance(dbc *pgxpool.Pool) {

	// Рецепты и список покупок

	var CreateStatements = NamedCreateStatements{
		NamedCreateStatement{
			TableName: "orders",
			CreateStatement: `CREATE TABLE finance.orders
			(
				id bigserial NOT NULL,
				address_id bigint,
				total bigint,
				confirmed boolean,
				paid boolean,
				tracking_number character varying(50) COLLATE pg_catalog."default",
				currency_id bigint NOT NULL,
				user_id uuid NOT NULL,
				CONSTRAINT orders_pkey PRIMARY KEY (id),
				CONSTRAINT orders_address_id_fkey FOREIGN KEY (address_id)
					REFERENCES "references".addresses (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT orders_currency_id_fkey FOREIGN KEY (currency_id)
					REFERENCES "references".currencies (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE finance.orders
				OWNER to postgres;
			
			CREATE INDEX fki_orders_address_id_fkey
				ON finance.orders USING btree
				(address_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_orders_currency_id_fkey
				ON finance.orders USING btree
				(currency_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_orders_user_id_fkey
				ON finance.orders USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "orders_details",
			CreateStatement: `CREATE TABLE finance.orders_details
			(
				id bigserial NOT NULL,
				order_id bigint NOT NULL,
				artwork_id bigint NOT NULL,
				amount bigint NOT NULL,
				currency_id bigint NOT NULL,
				CONSTRAINT order_details_pkey PRIMARY KEY (id),
				CONSTRAINT orders_details_artwork_id_fkey FOREIGN KEY (artwork_id)
					REFERENCES public.artworks (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT orders_details_currency_id_fkey FOREIGN KEY (currency_id)
					REFERENCES "references".currencies (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT orders_details_order_id_fkey FOREIGN KEY (order_id)
					REFERENCES finance.orders (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE finance.orders_details
				OWNER to postgres;
			
			CREATE INDEX fki_orders_details_artwork_id_fkey
				ON finance.orders_details USING btree
				(artwork_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_orders_details_currency_id_fkey
				ON finance.orders_details USING btree
				(currency_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_orders_details_order_id_fkey
				ON finance.orders_details USING btree
				(order_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "orders_returns",
			CreateStatement: `CREATE TABLE finance.orders_returns
			(
				id bigserial NOT NULL,
				order_id bigint NOT NULL,
				tracking_number character varying(50) COLLATE pg_catalog."default",
				package_recieved boolean,
				money_returned boolean,
				payment_number character varying(50) COLLATE pg_catalog."default",
				payment_date timestamp with time zone,
				file_id bigint,
				request text COLLATE pg_catalog."default",
				response text COLLATE pg_catalog."default",
				payment_amount bigint,
				currency_id bigint NOT NULL,
				CONSTRAINT orders_returns_pkey PRIMARY KEY (id),
				CONSTRAINT orders_returns_currency_id_fkey FOREIGN KEY (currency_id)
					REFERENCES "references".currencies (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT orders_returns_file_id_fkey FOREIGN KEY (file_id)
					REFERENCES "references".files (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT orders_returns_order_id_fkey FOREIGN KEY (order_id)
					REFERENCES finance.orders (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE finance.orders_returns
				OWNER to postgres;
			
			CREATE INDEX fki_orders_returns_currency_id_fkey
				ON finance.orders_returns USING btree
				(currency_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_orders_returns_file_id_fkey
				ON finance.orders_returns USING btree
				(file_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_orders_returns_order_id_fkey
				ON finance.orders_returns USING btree
				(order_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "payments",
			CreateStatement: `CREATE TABLE finance.payments
			(
				id bigserial NOT NULL,
				order_id bigint NOT NULL,
				amount money,
				date timestamp with time zone,
				currency_id bigint NOT NULL,
				CONSTRAINT payments_pkey PRIMARY KEY (id),
				CONSTRAINT payments_currency_id_fkey FOREIGN KEY (currency_id)
					REFERENCES "references".currencies (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT payments_order_id_fkey FOREIGN KEY (order_id)
					REFERENCES finance.orders (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE finance.payments
				OWNER to postgres;
			
			CREATE INDEX fki_payments_currency_id_fkey
				ON finance.payments USING btree
				(currency_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_payments_order_id_fkey
				ON finance.payments USING btree
				(order_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
	}

	for _, ncs := range CreateStatements {
		PostgreSQLExecuteCreateStatement(dbc, ncs)
	}

}
