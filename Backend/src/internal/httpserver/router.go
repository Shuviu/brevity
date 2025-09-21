package httpserver

import (
	"fmt"
	"net/http"
)

func InitializeRouter(router *http.ServeMux) {
	fmt.Println("Registering routes...")
	router.HandleFunc("/", HandleDefaultEndpoint)
	router.HandleFunc("/register", HandleRegisterShortUrlEndpoint)
	router.HandleFunc("/get", HandleGetShortUrlEndpoint)
	router.HandleFunc("/delete", HandleDeleteShortUrlEndpoint)

}
