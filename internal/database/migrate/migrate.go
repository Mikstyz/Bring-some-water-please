package migarte

import (
	"bring_some_water_please/internal/database"
	mData "bring_some_water_please/internal/database/migrate/Migrationsdata"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

type Migrate struct {
	db *sql.DB
}

func NewMigrate(db *sql.DB) *Migrate {
	return &Migrate{db: db}
}

func (r *Migrate) TablesInDb() (int32, error) {
	const selectTablesQuery = `
	SELECT name FROM sqlite_master
	WHERE type='table' AND name NOT LIKE 'sqlite_%';
	`

	rows, err := r.db.Query(selectTablesQuery)
	if err != nil {
		return 0, fmt.Errorf("ошибка при получении списка таблиц: %w", err)
	}
	defer rows.Close()

	var count int32
	for rows.Next() {
		count++
	}

	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("ошибка при обходе результатов: %w", err)
	}

	return count, nil
}

func (r *Migrate) CreateTables() error {
	fmt.Println("====================================================")

	tables := map[string]string{
		"UsersTable":        mData.UsersTable,
		"ModsTable":         mData.ModsTable,
		"AssembliesTable":   mData.AssembliesTable,
		"AssemblyModsTable": mData.AssemblyModsTable,
	}

	for name, query := range tables {
		res, err := r.db.Exec(query)
		if err != nil {
			log.Printf(`Ошибка при выполнении запроса:
	Error: %v
	Query: %s
	`, err, query)
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return fmt.Errorf("ошибка подсчета строк: %w", err)
		}

		if rowsAffected == 0 {
			log.Printf("[SQL][MIGRATE] - %s таблица уже существует", name)
		} else {
			log.Printf("[SQL][MIGRATE] - %s таблица успешно создана", name)
		}
	}

	fmt.Println("====================================================")
	return nil
}

func (r *Migrate) InsertDataInTables() error {
	inserts := []string{}

	for _, query := range inserts {
		res, err := r.db.Exec(query)
		if err != nil {
			log.Printf("Ошибка при выполнении запроса:\nError: %v\nQuery: %s", err, query)
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return fmt.Errorf("ошибка при подсчете кол-ва строк: %w", err)
		}

		if rowsAffected == 0 {
			return fmt.Errorf("не удалось вставить записи в таблицу\nQuery: %s", query)
		}

		log.Printf("Затронуто строк: %d", rowsAffected)
	}

	return nil
}

func (r *Migrate) BackupAndResetDB() error {
	dbPath := "internal/database/mods.db"
	backupPath := strings.TrimSuffix(dbPath, ".db") + "_OLD.db"

	// 1. Закрываем текущее соединение
	if r.db != nil {
		if err := r.db.Close(); err != nil {
			return fmt.Errorf("[BackupAndResetDB] ошибка при закрытии базы: %w", err)
		}
		r.db = nil
	}

	// 2. Переименовываем старый файл
	if _, err := os.Stat(dbPath); err == nil {
		if err := os.Rename(dbPath, backupPath); err != nil {
			return fmt.Errorf("[BackupAndResetDB] не удалось переименовать базу: %w", err)
		}
		log.Printf("[BackupAndResetDB] база переименована: %s → %s", dbPath, backupPath)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("[BackupAndResetDB] ошибка при проверке файла базы: %w", err)
	} else {
		log.Printf("[BackupAndResetDB] база не существует, пропускаем переименование")
	}

	// 3. Создаем новый пустой файл mods.db
	file, err := os.Create(dbPath)
	if err != nil {
		return fmt.Errorf("[BackupAndResetDB] не удалось создать новый файл базы: %w", err)
	}
	file.Close()
	log.Printf("[BackupAndResetDB] создан новый файл базы: %s", dbPath)

	// 4. Подключаемся к новой базе
	r.db = database.Connect()
	if r.db == nil {
		return fmt.Errorf("[BackupAndResetDB] ошибка при подключении к новой базе")
	}
	log.Printf("[BackupAndResetDB] успешное подключение к новой базе")

	return nil
}

func (r *Migrate) DropTable(table string) error {
	query := "DROP TABLE IF EXISTS " + table

	if _, err := r.db.Exec(query); err != nil {
		log.Fatalf("[SQL][DropTable] Ошибка при удалении таблицы %s: %v", table, err)
		return err
	}

	log.Printf("[SQL][DropTable] Таблица %s успешно удалена", table)
	return nil
}
