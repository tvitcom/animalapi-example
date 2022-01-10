package util

import (
	"database/sql"
)

func GetDbConn() *sql.DB {
	db, err := sql.Open("sqlite3", "apidb.sqlite3")
	PanicError(err)
	return db
}
