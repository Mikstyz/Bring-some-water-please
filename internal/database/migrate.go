package database

import (
	mData "bring_some_water_please/internal/database/Migrationsdata"
	"database/sql"
	"log"
)

type Migrate struct {
	db *sql.DB
}

func NewMigrate(db *sql.DB) *Migrate {
	return &Migrate{db: db}
}

func (r *Migrate) Run() error {

	tables := []string{
		mData.UsersTable,        //пользователи
		mData.ModsTable,         //Моды
		mData.VersionsTable,     //версии
		mData.FilesTable,        //файлы модов url
		mData.LoadersTable,      //ядра модов
		mData.AssembliesTable,   //сборки
		mData.AssemblyModsTable, //моды в сборках
	}

	for _, query := range tables {
		if _, err := r.db.Exec(query); err != nil {
			log.Printf(`Ошибка при выполнении запроса:
				Error: %v
				Query: %s
				`, err, query)

			return err
		}

		log.Print("[DATA BASE][MIGRATE] - Успешное создание таблицы")
	}

	return nil
}
