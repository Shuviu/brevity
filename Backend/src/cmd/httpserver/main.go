package main

import (
	"brevity/internal/httpserver"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	var server = http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	httpserver.InitializeRouter(router)

	fmt.Println("Starting server on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
