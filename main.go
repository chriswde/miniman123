package main

import (
	"log"

	"github.com/chriswde/miniman123/database"
)

func main() {
	log.Println(database.Init("./database/db.sqlite3"))
}
