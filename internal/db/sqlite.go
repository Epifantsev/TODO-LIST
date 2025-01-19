package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func New(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		log.Fatal("init db", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("ping db", err)
	}

	return db
}
