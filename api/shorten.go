package api

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chriswde/miniman123/database"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	stmt := "INSERT INTO `urls` (`long`) VALUES (?);"
	insert, err := database.Connection.Exec(stmt, r.FormValue("url"))
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

	s := base64.RawURLEncoding.EncodeToString([]byte(strconv.FormatInt(id, 10)))
	fmt.Fprintf(w, "%s", s)
}
