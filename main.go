package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chriswde/miniman123/api"
	"github.com/chriswde/miniman123/database"
)

func main() {
	err := database.Init("./database/db.sqlite3")
	if err != nil {
		log.Panicln(err)
	}

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := os.ReadFile("./html/index.html")
		w.Write(b)
	}))
	router.Handle("/api/shorten", http.HandlerFunc(api.Shorten))

	log.Println(http.ListenAndServe("127.0.0.1:80", router))
}
