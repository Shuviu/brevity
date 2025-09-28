package httpserver

import (
	"net/http"
)

func InitializeRouter(router *http.ServeMux, dbWrapper DbReqWrapper) {
	router.HandleFunc("/", dbWrapper.HandleDefaultEndpoint)
	router.HandleFunc("/register", dbWrapper.HandleRegisterShortUrlEndpoint)
	router.HandleFunc("/get", HandleGetShortUrlEndpoint)
	router.HandleFunc("/delete", HandleDeleteShortUrlEndpoint)
}
