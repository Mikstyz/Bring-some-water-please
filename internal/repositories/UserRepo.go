package repositories

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type UserRepo struct {
	db *sql.DB
}

func newUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func IsUser(r *UserRepo, tgId uint64) bool {

	fmt.Print(tgId)

	const Query string = ""

	return false
}

func NewUser(r *UserRepo, tgId uint64) bool {

	const Query string = ""

	return false
}
