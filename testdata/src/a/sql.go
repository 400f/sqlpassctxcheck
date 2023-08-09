package a

import (
	"database/sql"
)

func CallSqlDBMethod(db *sql.DB) {
	db.Begin()              // want `use \(\*database\/sql\.DB\)\.BeginTx instead of \(\*database\/sql\.DB\)\.Begin`
	db.Exec("SELECT 1")     // want `use \(\*database\/sql\.DB\)\.ExecContext instead of \(\*database\/sql\.DB\)\.Exec`
	db.Ping()               // want `use \(\*database\/sql\.DB\)\.PingContext instead of \(\*database\/sql\.DB\)\.Ping`
	db.Prepare("SELECT 1")  // want `use \(\*database\/sql\.DB\)\.PrepareContext instead of \(\*database\/sql\.DB\)\.Prepare`
	db.Query("SELECT 1")    // want `use \(\*database\/sql\.DB\)\.QueryContext instead of \(\*database\/sql\.DB\)\.Query`
	db.QueryRow("SELECT 1") // want `use \(\*database\/sql\.DB\)\.QueryRowContext instead of \(\*database\/sql\.DB\)\.QueryRow`
}

func CallSqlStmtMethod(stmt *sql.Stmt) {
	stmt.Exec("SELECT 1")     // want `use \(\*database\/sql\.Stmt\)\.ExecContext instead of \(\*database\/sql\.Stmt\)\.Exec`
	stmt.Query("SELECT 1")    // want `use \(\*database\/sql\.Stmt\)\.QueryContext instead of \(\*database\/sql\.Stmt\)\.Query`
	stmt.QueryRow("SELECT 1") // want `use \(\*database\/sql\.Stmt\)\.QueryRowContext instead of \(\*database\/sql\.Stmt\)\.QueryRow`
}

func CallSqlTxMethod(tx *sql.Tx, stmt *sql.Stmt) {
	tx.Exec("SELECT 1")     // want `use \(\*database\/sql\.Tx\)\.ExecContext instead of \(\*database\/sql\.Tx\)\.Exec`
	tx.Prepare("SELECT 1")  // want `use \(\*database\/sql\.Tx\)\.PrepareContext instead of \(\*database\/sql\.Tx\)\.Prepare`
	tx.Query("SELECT 1")    // want `use \(\*database\/sql\.Tx\)\.QueryContext instead of \(\*database\/sql\.Tx\)\.Query`
	tx.QueryRow("SELECT 1") // want `use \(\*database\/sql\.Tx\)\.QueryRowContext instead of \(\*database\/sql\.Tx\)\.QueryRow`
	tx.Stmt(stmt)           // want `use \(\*database\/sql\.Tx\)\.StmtContext instead of \(\*database\/sql\.Tx\)\.Stmt`
}
