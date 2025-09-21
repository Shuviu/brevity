package main

import (
	"brevity/internal/httpserver"
	"brevity/internal/sqliteDatabase"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	var server = http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	sqliteDatabase.InitializeDatabase()
	httpserver.InitializeRouter(router)

	fmt.Println("Starting server on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
