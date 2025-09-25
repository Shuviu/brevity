package httpserver

import (
	"fmt"
	"net/http"
)

func InitializeRouter(router *http.ServeMux, dbWrapper DbReqWrapper) {
	fmt.Println("Registering routes...")
	router.HandleFunc("/", dbWrapper.HandleDefaultEndpoint)
	router.HandleFunc("/register", dbWrapper.HandleRegisterShortUrlEndpoint)
	router.HandleFunc("/get", HandleGetShortUrlEndpoint)
	router.HandleFunc("/delete", HandleDeleteShortUrlEndpoint)
}
