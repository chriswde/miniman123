package api

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/chriswde/miniman123/database"
)

func Resolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	encoded := strings.TrimPrefix(r.URL.Path, "/")
	decoded, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(string(decoded), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt := "SELECT `url` FROM `urls` WHERE `id` = ?;"
	result := database.Connection.QueryRow(stmt, id)

	var url string
	result.Scan(&url)

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
