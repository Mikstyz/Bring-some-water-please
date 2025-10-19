package database

import "database/sql"

const FilesTable string = `
CREATE TABLE IF NOT EXISTS files
`

const ModsTable string = `
CREATE TABLE IF NOT EXISTS mods
`

const LoadersTable string = `
CREATE TABLE IF NOT EXISTS loaders
`

const DataTable string = `
CREATE TABLE IF NOT EXISTS users	
`

const UserTable = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY NOT NULL,
    tgId INTEGER NOT NULL
)
`

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
