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

func (r *Migrate) CreateTables() error {

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

		log.Print("[SQL][MIGRATE] - Успешное создание таблицы")
	}

	return nil
}
func (r *Migrate) DropAllTables() error {
	const selectTablesQuery = `
	SELECT name FROM sqlite_master
	WHERE type='table' AND name NOT LIKE 'sqlite_%';
	`

	rows, err := r.db.Query(selectTablesQuery)
	if err != nil {
		log.Fatalf("[SQL][Drop All Tables] Ошибка при получении списка таблиц: %v", err)
		return err
	}
	defer rows.Close()

	var tableName string
	var tables []string

	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("[SQL][Drop All Tables] Ошибка при сканировании: %v", err)
			return err
		}
		tables = append(tables, tableName)
	}

	removedTables := 0
	for _, t := range tables {
		if _, err := r.db.Exec("DROP TABLE IF EXISTS " + t); err != nil {
			log.Fatalf("[SQL][Drop All Tables] Ошибка при удалении таблицы %s: %v", t, err)
			return err
		}
		removedTables++
		log.Printf("[SQL][Drop All Tables] Таблица удалена: %s", t)
	}

	log.Printf("[SQL][Drop All Tables] Успешное удаление таблиц. Всего удалено: %d", removedTables)
	return nil
}
