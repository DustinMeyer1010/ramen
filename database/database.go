package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("db.Ping failed:", err)
	}

}
