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
		stmt := "INSERT INTO `urls` (`url`) VALUES (?);"
		insert, err := database.Connection.Exec(stmt, url)
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

		s := fmt.Sprintf("%s/%s", configuration.Configuration.Host, base64.RawURLEncoding.EncodeToString([]byte(strconv.FormatInt(id, 10))))

		if isHTMXRequest {
			re := `<div class="alert alert-success text-center user-select-all" id="urlAlert">` +
				s +
				`</div>`
			fmt.Fprintf(w, "%s", re)
		} else {
			fmt.Fprintf(w, "%s", s)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
