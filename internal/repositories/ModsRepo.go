package repositories

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type ModRepo struct {
	db *sql.DB
}

func newModRepo(db *sql.DB) *ModRepo {
	return &ModRepo{db: db}
}

func IsThereMod(mod string) bool {

	const Query string = ""
	return false
}

func SaveMod(mod string) error {

	const Query string = ""
	return nil
}

func LoadMod(mod string) error {

	const Query string = ""
	return nil
}
