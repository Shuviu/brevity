package httpserver

import (
	"brevity/internal/sqliteDatabase"
	"database/sql"
	"net/http"
	"net/url"
	"strings"
)

type DbReqWrapper struct {
	Db *sql.DB
}

// HandleDefaultEndpoint redirects to the url provided in the request url
func (dbwrapper DbReqWrapper) HandleDefaultEndpoint(w http.ResponseWriter, r *http.Request) {
	// fetch redirection hash
	var reqUrl string = r.URL.Path
	shortUrl := strings.Split(reqUrl, "/")[1]

	// fetch long url form db and redirect
	longUrl := sqliteDatabase.GetLongUrlFromShort(shortUrl, dbwrapper.Db)
	if longUrl == "" {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("No corresponding entry for this url"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	http.Redirect(w, r, longUrl, http.StatusPermanentRedirect)
}

// HandleRegisterShortUrlEndpoint stores the provided url in the database and returns the given hash as a response.
func (dbWrapper DbReqWrapper) HandleRegisterShortUrlEndpoint(w http.ResponseWriter, r *http.Request) {
	// fetch and check the provided url
	var params url.Values = r.URL.Query()
	longUrl := params.Get("url")
	if longUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("url is required"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	// generate redirection hash and store in db
	var shortUrl string = GenerateShortFromLongUrl(longUrl)
	if !sqliteDatabase.InsertNewUrl(longUrl, shortUrl, dbWrapper.Db) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// respond with the redirection hash
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(shortUrl))
	w.Header().Set("Content-Type", "json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleGetShortUrlEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
func HandleDeleteShortUrlEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
