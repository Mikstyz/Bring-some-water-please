package repositories

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type ModRepo struct {
	db *sql.DB
}

func NewModRepo(db *sql.DB) *ModRepo {
	return &ModRepo{db: db}
}

func IsThereMod(mod string) bool {

	return false
}

func SaveMod(mod string) error {
	return nil
}

func LoadMod(mod string) error {
	return nil
}
