package a

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func CallSqlxDBMethod(db *sqlx.DB) {
	db.Begin()                     // want `use \(\*database\/sql\.DB\)\.BeginTx instead of \(\*database\/sql\.DB\)\.Begin`
	db.Exec("SELECT 1")            // want `use \(\*database\/sql\.DB\)\.ExecContext instead of \(\*database\/sql\.DB\)\.Exec`
	db.Ping()                      // want `use \(\*database\/sql\.DB\)\.PingContext instead of \(\*database\/sql\.DB\)\.Ping`
	db.Prepare("SELECT 1")         // want `use \(\*database\/sql\.DB\)\.PrepareContext instead of \(\*database\/sql\.DB\)\.Prepare`
	db.Query("SELECT 1")           // want `use \(\*database\/sql\.DB\)\.QueryContext instead of \(\*database\/sql\.DB\)\.Query`
	db.QueryRow("SELECT 1")        // want `use \(\*database\/sql\.DB\)\.QueryRowContext instead of \(\*database\/sql\.DB\)\.QueryRow`
	db.Beginx()                    // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.BeginTxx instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.Beginx`
	db.Get(nil, "SELECT 1")        // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.GetContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.Get`
	db.MustBegin()                 // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.MustBeginTx instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.MustBegin`
	db.MustExec("SELECT 1")        // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.MustExecContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.MustExec`
	db.NamedExec("SELECT 1", nil)  // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.NamedExecContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.NamedExec`
	db.NamedQuery("SELECT 1", nil) // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.NamedQueryContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.NamedQuery`
	db.PrepareNamed("SELECT 1")    // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.PrepareNamedContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.PrepareNamed`
	db.Preparex("SELECT 1")        // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.PreparexContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.Preparex`
	db.QueryRowx("SELECT 1")       // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.QueryRowxContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.QueryRowx`
	db.Queryx("SELECT 1")          // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.QueryxContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.Queryx`
	db.Select(nil, "SELECT 1")     // want `use \(\*github.com\/jmoiron\/sqlx\.DB\)\.SelectContext instead of \(\*github.com\/jmoiron\/sqlx\.DB\)\.Select`
}

func CallSqlxNamedStmtMethod(stmt *sqlx.NamedStmt) {
	stmt.Exec(nil)        // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.ExecContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.Exec`
	stmt.Get(nil, nil)    // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.GetContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.Get`
	stmt.MustExec(nil)    // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.MustExecContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.MustExec`
	stmt.Query(nil)       // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.QueryContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.Query`
	stmt.QueryRow(nil)    // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.QueryRowContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.QueryRow`
	stmt.QueryRowx(nil)   // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.QueryRowxContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.QueryRowx`
	stmt.Queryx(nil)      // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.QueryxContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.Queryx`
	stmt.Select(nil, nil) // want `use \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.SelectContext instead of \(\*github.com\/jmoiron\/sqlx\.NamedStmt\)\.Select`
}

func CallSqlxStmtMethod(stmt *sqlx.Stmt) {
	stmt.Exec("SELECT 1")     // want `use \(\*database\/sql\.Stmt\)\.ExecContext instead of \(\*database\/sql\.Stmt\)\.Exec`
	stmt.Query("SELECT 1")    // want `use \(\*database\/sql\.Stmt\)\.QueryContext instead of \(\*database\/sql\.Stmt\)\.Query`
	stmt.QueryRow("SELECT 1") // want `use \(\*database\/sql\.Stmt\)\.QueryRowContext instead of \(\*database\/sql\.Stmt\)\.QueryRow`
	stmt.Get(nil)             // want `use \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.GetContext instead of \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.Get`
	stmt.MustExec()           // want `use \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.MustExecContext instead of \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.MustExec`
	stmt.QueryRowx()          // want `use \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.QueryRowxContext instead of \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.QueryRowx`
	stmt.Queryx()             // want `use \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.QueryxContext instead of \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.Queryx`
	stmt.Select(nil)          // want `use \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.SelectContext instead of \(\*github.com\/jmoiron\/sqlx\.Stmt\)\.Select`
}

func CallSqlxTxMethod(tx *sqlx.Tx, stmt *sql.Stmt, nstmt *sqlx.NamedStmt) {
	tx.Exec("SELECT 1")           // want `use \(\*database\/sql\.Tx\)\.ExecContext instead of \(\*database\/sql\.Tx\)\.Exec`
	tx.Prepare("SELECT 1")        // want `use \(\*database\/sql\.Tx\)\.PrepareContext instead of \(\*database\/sql\.Tx\)\.Prepare`
	tx.Query("SELECT 1")          // want `use \(\*database\/sql\.Tx\)\.QueryContext instead of \(\*database\/sql\.Tx\)\.Query`
	tx.QueryRow("SELECT 1")       // want `use \(\*database\/sql\.Tx\)\.QueryRowContext instead of \(\*database\/sql\.Tx\)\.QueryRow`
	tx.Stmt(stmt)                 // want `use \(\*database\/sql\.Tx\)\.StmtContext instead of \(\*database\/sql\.Tx\)\.Stmt`
	tx.Get(nil, "SELECT 1")       // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.GetContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.Get`
	tx.MustExec("SELECT 1")       // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.MustExecContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.MustExec`
	tx.NamedExec("SELECT 1", nil) // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.NamedExecContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.NamedExec`
	tx.NamedStmt(nstmt)           // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.NamedStmtContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.NamedStmt`
	tx.PrepareNamed("SELECT 1")   // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.PrepareNamedContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.PrepareNamed`
	tx.Preparex("SELECT 1")       // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.PreparexContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.Preparex`
	tx.QueryRowx("SELECT 1")      // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.QueryRowxContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.QueryRowx`
	tx.Queryx("SELECT 1")         // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.QueryxContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.Queryx`
	tx.Select(nil, "SELECT 1")    // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.SelectContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.Select`
	tx.Stmtx(nil)                 // want `use \(\*github.com\/jmoiron\/sqlx\.Tx\)\.StmtxContext instead of \(\*github.com\/jmoiron\/sqlx\.Tx\)\.Stmtx`
}
