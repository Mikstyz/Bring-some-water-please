package main

import (
	//Sc "bring_some_water_please/internal/scrper"
	//Tg "bring_some_water_please/internal/bot"
	db "bring_some_water_please/internal/database"
	m "bring_some_water_please/internal/database/migrate"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadConfig() {
	confFile := "conf.env"

	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		fmt.Printf(`
====================================================
Config file not found: %s
====================================================
`, confFile)
	}

	fmt.Print(`
====================================================
Loading config...
====================================================`)

	if err := godotenv.Load(confFile); err != nil {
		fmt.Printf(`
====================================================
Error loading %s: %v
====================================================
`, confFile, err)
	}

	fmt.Printf(`
====================================================
Config loaded: %s
====================================================`, confFile)
}

func loadDB() *sql.DB {
	fmt.Print(`
====================================================
[main.go] Подключение к базе данных
====================================================
`)

	r := db.Connect()
	migrate := m.NewMigrate(r)

	err := migrate.CreateTables()
	if err != nil {
		log.Printf(`
====================================================
[main.go] Не получилось создать недостающие таблиццы
====================================================
`)
	}

	return r
}

func main() {
	loadConfig()
	loadDB()

	//test.MigrateTest(r)
	//Tg.Tgbot()

	//Sc.View("Not Enough Animations")
}
