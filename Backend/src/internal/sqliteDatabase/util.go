package sqliteDatabase

import (
	"database/sql"
	"log"
)

func openSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
