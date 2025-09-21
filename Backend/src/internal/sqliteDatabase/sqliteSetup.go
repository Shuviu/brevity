package sqliteDatabase

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase() {
	db := openSqliteDB()

	sqlStmt :=
		`CREATE TABLE IF NOT EXISTS "url_map" (
		"url_id" VARCHAR NOT NULL UNIQUE,
		"long_url" VARCHAR NOT NULL,
		"creation_date" DATETIME DEFAULT current_timestamp,
		PRIMARY KEY("url_id")
		);
		`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized successfully...")
}
