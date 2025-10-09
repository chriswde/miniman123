package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chriswde/miniman123/api"
	"github.com/chriswde/miniman123/configuration"
	"github.com/chriswde/miniman123/database"
)

func main() {
	err := database.Init("./database/db.sqlite3")
	if err != nil {
		log.Panicln(err)
	}

	err = configuration.Configuration.Init("./config.json")
	if err != nil {
		log.Panicln(err)
	}

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Handle("/", http.HandlerFunc(HandleIndex))
	router.Handle("/api/shorten", http.HandlerFunc(api.Shorten))

	log.Println(http.ListenAndServe(configuration.Configuration.Host, router))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	isIndex := strings.TrimPrefix(r.URL.Path, "/")
	if isIndex == "" {
		b, err := os.ReadFile("./html/index.html")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	} else {
		api.Resolve(w, r)
	}
}
