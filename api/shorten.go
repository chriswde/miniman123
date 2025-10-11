package api

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/chriswde/miniman123/configuration"
	"github.com/chriswde/miniman123/database"
	"github.com/google/uuid"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	var isHTMXRequest = false
	hxHeader := r.Header["Hx-Request"]
	if len(hxHeader) == 1 {
		var err error
		isHTMXRequest, err = strconv.ParseBool(hxHeader[0])
		if err != nil {
			isHTMXRequest = false
		}
	}

	url := r.FormValue("url")
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		stmt := "INSERT INTO `urls` (`url`, `deletion_token`) VALUES (?, ?);"
		deletionToken := uuid.NewString()
		insert, err := database.Connection.Exec(stmt, url, deletionToken)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "%s", "Something happened Ö")
			return
		}

		id, err := insert.LastInsertId()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "%s", "Something happened Ö")
			return
		}

		shortURL := fmt.Sprintf("%s/%s", configuration.Configuration.HostAddress, base64.RawURLEncoding.EncodeToString([]byte(strconv.FormatInt(id, 10))))
		encodedDeletionToken := base64.RawURLEncoding.EncodeToString([]byte(deletionToken))

		if isHTMXRequest {
			re :=
				`<div class="alert alert-success text-center user-select-all">` +
					shortURL +
					"</div>" +
					`<div class="alert alert-danger text-center"><button class="btn btn-danger" type="button" data-bs-toggle="collapse" data-bs-target="#deletionToken" aria-expanded="false" aria-controls="deletionToken">Show deletion token</button><br><div class="collapse" id="deletionToken">` +
					encodedDeletionToken +
					"</div></div>"
			fmt.Fprintf(w, "%s", re)
		} else {
			fmt.Fprintf(w, "%s\n%s", shortURL, encodedDeletionToken)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
