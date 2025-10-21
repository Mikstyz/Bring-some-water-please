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
	m := migrate.NewMigrate(r)

	// // -------------------------
	// // 1. Создание таблиц
	// // -------------------------
	// startTest("migrate", "CreateTables")
	// if err := m.CreateTables(); err != nil {
	// 	log.Printf("[TEST][migrate][CreateTables] - ERROR: %v", err)
	// }
	// finishTest("migrate", "CreateTables")

	// // -------------------------
	// // 2. Удаление конкретной таблицы: users
	// // -------------------------
	// startTest("migrate", "DropTable users")
	// if err := m.DropTable("users"); err != nil {
	// 	log.Printf("[TEST][migrate][DropTable users] - ERROR: %v", err)
	// }
	// finishTest("migrate", "DropTable users")

	// -------------------------
	// 3. Фулл сброс базы через BackupAndResetDB
	// -------------------------
	startTest("migrate", "BackupAndResetDB")
	if err := m.BackupAndResetDB(); err != nil {
		log.Printf("[TEST][migrate][BackupAndResetDB] - ERROR: %v", err)
	}
	finishTest("migrate", "BackupAndResetDB")

	// -------------------------
	// 4. Проверка, что таблиц нет
	// -------------------------
	startTest("migrate", "CheckTablesDeleted")
	count, err := m.TablesInDb()
	if err != nil {
		log.Printf("[TEST][migrate][CheckTablesDeleted] - ERROR: %v", err)
	} else if count != 0 {
		messageTest("migrate", fmt.Sprintf("не все таблицы были удалены, осталось: %d", count))
	} else {
		messageTest("migrate", "все таблицы успешно удалены")
	}
}
