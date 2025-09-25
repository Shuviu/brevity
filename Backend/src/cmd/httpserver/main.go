package main

import (
	"brevity/internal/databaseHandling"
	"brevity/internal/httpserver"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	router := http.NewServeMux()

	loadEnvironmentFile()
	dbWrapper := setupDatabaseWrapper()
	handler := setupCorsHeaders(router)

	// finish httpserver init
	var server = http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	httpserver.InitializeRouter(router, dbWrapper)

	log.Println("Listening on port 8080")
	err := server.ListenAndServe()

	if err != nil {
		return
	}
}

func loadEnvironmentFile() {
	if len(os.Args) < 2 {
		log.Fatal("Missing argument for the .env filepath")
	}

	err := godotenv.Load(os.Args[1])
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setupCorsHeaders(router http.Handler) http.Handler {
	// setup cors headers
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://brevity.shuviu.de"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)

	return handler
}

func setupDatabaseWrapper() httpserver.DbReqWrapper {
	dbType := os.Getenv("DB_TYPE")
	var dbWrapper httpserver.DbReqWrapper

	switch dbType {
	case "postgres":
		dbWrapper = httpserver.DbReqWrapper{Db: sqliteDatabase.OpenPostgresDB()}
		break

	case "sqlite":
		dbWrapper = httpserver.DbReqWrapper{Db: sqliteDatabase.OpenSqliteDB()}
		break

	default:
		log.Fatal("Unsupported database type: " + dbType)
	}

	sqliteDatabase.InitializeDatabase(dbWrapper.Db)

	return dbWrapper
}
