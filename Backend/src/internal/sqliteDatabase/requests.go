package sqliteDatabase

func InsertNewUrl(longUrl string, shortUrl string) bool {
	db := openSqliteDB()

	urlInsert, errStmt := db.Prepare("INSERT INTO url_map(long_url, url_id) VALUES(?, ?)")
	if errStmt != nil {
		return false
	}

	_, errInsert := urlInsert.Exec(longUrl, shortUrl)
	if errInsert != nil {
		return false
	}

	return true
}

func GetLongUrlFromShort(shortUrl string) string {
	db := openSqliteDB()
	selectStmt, errStmt := db.Prepare("SELECT long_url FROM url_map WHERE url_id=?")
	if errStmt != nil {
		return ""
	}

	res := selectStmt.QueryRow(shortUrl)
	var longUrl string
	err := res.Scan(&longUrl)
	if err != nil {
		return ""
	}

	return longUrl
}

func DeleteEntryFromShortUrl(shortUrl string) bool {
	db := openSqliteDB()
	deleteStmt, errStmt := db.Prepare("DELETE FROM url_map WHERE url_id=?")
	if errStmt != nil {
		return false
	}

	_, errDelete := deleteStmt.Exec(shortUrl)
	if errDelete != nil {
		return false
	}
	return true
}
