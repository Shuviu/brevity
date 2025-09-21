package httpserver

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
)

func GenerateShortFromLongUrl(longUrl string) string {
	hasher := sha256.New()
	_, err := io.WriteString(hasher, longUrl)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))[:8]
}
