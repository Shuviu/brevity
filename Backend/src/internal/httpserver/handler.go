package httpserver

import (
	"brevity/internal/sqliteDatabase"
	"net/http"
	"net/url"
)

func HandleDefaultEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func HandleRegisterShortUrlEndpoint(w http.ResponseWriter, r *http.Request) {
	var params url.Values = r.URL.Query()
	longUrl := params.Get("url")

	if longUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("url is required"))
		if err != nil {
			return
		}
		return
	}

	if !sqliteDatabase.InsertNewUrl(longUrl, GenerateShortFromLongUrl(longUrl)) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleGetShortUrlEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
func HandleDeleteShortUrlEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
