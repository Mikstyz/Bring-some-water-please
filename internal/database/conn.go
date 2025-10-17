package database

import (
	"database/sql"
	"log"
	"os"
)

func Connect() *sql.DB {
	dbpath := os.Getenv("DBPATH")

	db, err := sql.Open("sqlite", dbpath)

	if err != nil {
		log.Fatalf("Error conn db\nPath:%s\nError: %v", dbpath, err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Db unavailable")
	}

	log.Println("db conn")
	return db
}
