package main

import (
	//Tg "bring_some_water_please/internal/bot"

	db "bring_some_water_please/internal/database"
	mig "bring_some_water_please/internal/database/migrate"

	//scr "bring_some_water_please/internal/scrper"
	//conv "bring_some_water_please/utils/converter"

	serivce "bring_some_water_please/internal/service"

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
	migrate := mig.NewMigrate(r)

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
	modname := ("fabric api")
	loadConfig()
	//loadDB()

	conn := loadDB()

	modService := serivce.NewModSerivce(conn)

	modService.DownloadMod(modname, "1.21", "fabric")

	//test.MigrateTest(r)
	//Tg.Tgbot()

	//Sc.View("Not Enough Animations")
}
