package main

import (
	"brevity/internal/httpserver"
	"brevity/internal/sqliteDatabase"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := http.NewServeMux()

	corsHandler := cors.New(cors.Options{
		// Allow your frontend's origin
		AllowedOrigins: []string{"http://localhost:3000", "https://brevity.shuviu.de"},
		// Add other necessary headers and methods
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		// Set to true if your frontend needs to send credentials like cookies
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)
	var server = http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Starting Initialization...")
	sqliteDatabase.InitializeDatabase()
	httpserver.InitializeRouter(router)

	fmt.Println("Starting server on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
