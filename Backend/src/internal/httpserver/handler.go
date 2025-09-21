package httpserver

import (
	"brevity/internal/sqliteDatabase"
	"net/http"
	"net/url"
	"strings"
)

func HandleDefaultEndpoint(w http.ResponseWriter, r *http.Request) {
	var reqUrl string = r.URL.Path
	shortUrl := strings.Split(reqUrl, "/")[1]

	longUrl := sqliteDatabase.GetLongUrlFromShort(shortUrl)
	if longUrl == "" {
		http.NotFound(w, r)
	}

	http.Redirect(w, r, longUrl, http.StatusPermanentRedirect)
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

	var shortUrl string = GenerateShortFromLongUrl(longUrl)
	if !sqliteDatabase.InsertNewUrl(longUrl, shortUrl) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err := w.Write([]byte(shortUrl))
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
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
