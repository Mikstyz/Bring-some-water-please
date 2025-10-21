package database

import (
	"database/sql"
)

type Migrate struct {
	db *sql.DB
}

func NewMigrate(db *sql.DB) *Migrate {
	return &Migrate{db: db}
}

func CreateTable(r *Migrate) {

}

func RemoveTable(r *Migrate) {

}
