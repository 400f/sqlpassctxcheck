package a

import "database/sql"

func CallWithMethodReturn(get func() (bool, *sql.DB)) {
	_, db := get()
	db.Begin() // want `use \(\*database\/sql\.DB\)\.BeginTx instead of \(\*database\/sql\.DB\)\.Begin`
}
