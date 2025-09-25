package sqliteDatabase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase(database *sql.DB) {
	db := database

	sqlStmt :=
		`CREATE TABLE IF NOT EXISTS "url_map" (
		"url_id" VARCHAR(50) NOT NULL UNIQUE,
		"long_url" VARCHAR(50) NOT NULL,
		"creation_date" TIMESTAMP DEFAULT current_timestamp,
		PRIMARY KEY("url_id")
		);
		`
	err1 := db.Ping()
	if err1 != nil {
		log.Fatal(err1)
	}

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized successfully...")
}
