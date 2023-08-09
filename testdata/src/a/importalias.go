package a

import (
	aliasedsql "database/sql"

	aliasedsqlx "github.com/jmoiron/sqlx"
)

func CallSql(db *aliasedsql.DB) {
	db.Query("SELECT 1") // want `use \(\*database\/sql\.DB\)\.QueryContext instead of \(\*database\/sql\.DB\)\.Query`
}

func CallSqlx(db *aliasedsqlx.DB) {
	db.Queryx("SELECT 1") // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.QueryxContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.Queryx`
}
