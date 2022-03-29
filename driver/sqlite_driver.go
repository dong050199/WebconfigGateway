package driver

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	SQL *sql.DB
}

var SQLite = &SQLiteDB{}

func Connect(dbDirectory string) *SQLiteDB {
	db, err := sql.Open("sqlite3", dbDirectory)
	if err != nil {
		panic(err)
	}
	SQLite.SQL = db
	return SQLite
}
