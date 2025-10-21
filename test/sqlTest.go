package test

import (
	migrate "bring_some_water_please/internal/database/migrate"
	"database/sql"
	"fmt"
	"log"
)

// Единотипное начало теста
func startTest(module, step string) {
	log.Printf("-------------------------------------------")
	log.Printf("[TEST][%s][%s] - START", module, step)
}

// Сообщение во время теста
func messageTest(module, step string) {
	log.Printf("[TEST][%s][%s] - SUCCESS", module, step)
	log.Printf("-------------------------------------------")
}

// Единотипное завершение теста
func finishTest(module, step string) {
	log.Printf("[TEST][%s][%s] - SUCCESS", module, step)
	log.Printf("-------------------------------------------")
}

func MigrateTest(r *sql.DB) {
	var errorInTest int32

	m := migrate.NewMigrate(r)

	// -------------------------
	// Создание таблиц
	// -------------------------
	startTest("migrate", "CreateTables")
	if err := m.CreateTables(); err != nil {
		log.Printf("[TEST][migrate][CreateTables] - ERROR: %v", err)
		errorInTest++
	}
	finishTest("migrate", "CreateTables")

	// -------------------------
	// Удаление конкретной таблицы: users
	// -------------------------
	startTest("migrate", "DropTable users")
	if err := m.DropTable("users"); err != nil {
		log.Printf("[TEST][migrate][DropTable users] - ERROR: %v", err)
		errorInTest++
	}
	finishTest("migrate", "DropTable users")

	// -------------------------
	// Вставка данных
	// -------------------------
	startTest("migrate", "InsertDataInTables")
	if err := m.InsertDataInTables(); err != nil {
		log.Printf("[TEST][migrate][InsertDataInTables] - ERROR: %v", err)
		errorInTest++
	}
	finishTest("migrate", "InsertDataInTables")

	// -------------------------
	// Фулл сброс базы через BackupAndResetDB
	// -------------------------
	startTest("migrate", "BackupAndResetDB")
	if err := m.BackupAndResetDB(); err != nil {
		log.Printf("[TEST][migrate][BackupAndResetDB] - ERROR: %v", err)
		errorInTest++
	}
	finishTest("migrate", "BackupAndResetDB")

	// -------------------------
	// Проверка, что таблиц нет
	// -------------------------
	startTest("migrate", "CheckTablesDeleted")
	count, err := m.TablesInDb()
	if err != nil {
		log.Printf("[TEST][migrate][CheckTablesDeleted] - ERROR: %v", err)
		errorInTest++
	} else if count != 0 {
		messageTest("migrate", fmt.Sprintf("не все таблицы были удалены, осталось: %d", count))
	} else {
		messageTest("migrate", "все таблицы успешно удалены")
	}

	if errorInTest != 0 {
		messageTest("migrate", fmt.Sprintf("не все тесты прошли без ошибок: %d", errorInTest))
	} else {
		messageTest("migrate", "все тесты прошли без ошибок")
	}
}
