package api

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/chriswde/miniman123/internal/database"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	encodedDeletionToken := strings.TrimPrefix(r.URL.Path, "/api/delete/")
	deletionToken, err := base64.RawURLEncoding.DecodeString(encodedDeletionToken)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt := "DELETE FROM `urls` WHERE `deletion_token` = ?;"
	result, err := database.Connection.Exec(stmt, string(deletionToken))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if affected != 1 {
		fmt.Fprintln(w, "Entry not found")
		return
	}

	fmt.Fprintln(w, "Entry removed")
}
