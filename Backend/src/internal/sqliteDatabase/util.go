package sqliteDatabase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func OpenSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func OpenPostgresDB() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "admin"
		password = "egon34"
		dbname   = "mydb"
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
