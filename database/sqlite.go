package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var Connection *sql.DB

func Init(dbPath string) error {
	err := createDBFileIfNotExists(dbPath)
	if err != nil {
		return err
	}

	Connection, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	err = migrate(Connection)
	if err != nil {
		return err
	}

	log.Println(Connection.Ping())
	return nil
}

func createDBFileIfNotExists(dbPath string) error {
	_, err := os.Stat(dbPath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			dbFile, err := os.Create(dbPath)
			if err != nil {
				return nil
			}
			defer dbFile.Close()

			log.Println("Database path didn't exist. Created empty file.")
		} else {
			return err
		}
	}

	return nil
}

func migrate(db *sql.DB) error {
	migrations, err := os.ReadDir("./database/migration/")
	if err != nil {
		return nil
	}

	for _, m := range migrations {
		if m.IsDir() {
			continue
		}

		cont, err := os.ReadFile("./database/migration/" + m.Name())
		if err != nil {
			return err
		}

		cmd := string(cont)
		_, err = db.Exec(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}
