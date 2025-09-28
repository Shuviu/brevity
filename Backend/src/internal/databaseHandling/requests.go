package sqliteDatabase

import "database/sql"

// InsertNewUrl inserts the provided long url with redirection hash in the database returns true if the entry was inserted successfully
func InsertNewUrl(longUrl string, shortUrl string, database *sql.DB) bool {

	// Prepare sql statement
	urlInsert, errStmt := database.Prepare("INSERT INTO url_map(long_url, url_id) VALUES($1, $2)")
	if errStmt != nil {
		return false
	}
	// Exec statement
	_, errInsert := urlInsert.Exec(longUrl, shortUrl)
	if errInsert != nil {
		return false
	}

	return true
}

// GetLongUrlFromShort Retrieves the longUrl corresponding to the redirection hash stored in the database. Returns empty string if not found
func GetLongUrlFromShort(shortUrl string, database *sql.DB) string {

	// prepare and exec sql statement
	selectStmt, errStmt := database.Prepare("SELECT long_url FROM url_map WHERE url_id=$1")
	if errStmt != nil {
		return ""
	}

	res := selectStmt.QueryRow(shortUrl)

	// fetch url from sql response
	var longUrl string
	err := res.Scan(&longUrl)
	if err != nil {
		return ""
	}

	return longUrl
}

// DeleteEntryFromShortUrl Deletes the entry corresponding to the provided redirection hash provided
func DeleteEntryFromShortUrl(shortUrl string, database *sql.DB) bool {

	// Prepare sql statement
	deleteStmt, errStmt := database.Prepare("DELETE FROM url_map WHERE url_id=?")
	if errStmt != nil {
		return false
	}

	// Exec sql
	_, errDelete := deleteStmt.Exec(shortUrl)
	if errDelete != nil {
		return false
	}
	return true
}
