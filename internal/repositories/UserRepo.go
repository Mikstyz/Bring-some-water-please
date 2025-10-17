package repositories

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func isUser(r *UserRepo) bool {
	return false
}

func NewUser(r *UserRepo) bool {
	return false
}
