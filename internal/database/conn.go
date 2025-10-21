package database

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func Connect() *sql.DB {
	dbpath := os.Getenv("DBPATH")

	db, err := sql.Open("sqlite", dbpath)

	log.Print("[Dbconnect] Подключение к базе данных...")
	log.Printf("[Dbconnect] DBPATH: %s", dbpath)

	if err != nil {
		log.Fatalf("[Dbconnect] Error conn db\nPath:%s\nError: %v", dbpath, err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("[Dbconnect] Db unavailable")
	}

	log.Println("[Dbconnect] Успешное подключение к базе данных")
	return db
}
