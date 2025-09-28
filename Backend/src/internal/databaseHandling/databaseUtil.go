package sqliteDatabase

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func OpenSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", fetchEnvVariable("SQLITE_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func OpenPostgresDB() *sql.DB {
	var (
		host      = fetchEnvVariable("POSTGRES_HOST")
		port, err = strconv.Atoi(fetchEnvVariable("POSTGRES_PORT"))
		user      = fetchEnvVariable("POSTGRES_USER")
		password  = fetchEnvVariable("POSTGRES_PASSWORD")
		dbname    = fetchEnvVariable("POSTGRES_DB")
	)

	if err != nil {
		log.Fatal("Port is not an integer!")
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitializeDatabase(database *sql.DB) {

	sqlStmt :=
		`CREATE TABLE IF NOT EXISTS "url_map" (
		"url_id" VARCHAR(50) NOT NULL UNIQUE,
		"long_url" VARCHAR(255) NOT NULL,
		"creation_date" TIMESTAMP DEFAULT current_timestamp,
		PRIMARY KEY("url_id")
		);
		`

	_, err := database.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

// fetchEnvVariable reads the value of key from the environment file. If not set a fatal error is logged
func fetchEnvVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal(key + " is not defined in environment")
	}
	return value
}
