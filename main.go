package main

import (
	//Sc "bring_some_water_please/internal/scrper"
	//Tg "bring_some_water_please/internal/bot"
	db "bring_some_water_please/internal/database"
	test "bring_some_water_please/test"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadConfig() {
	confFile := "conf.env"

	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", confFile)
	}

	log.Print("Loading config...")

	if err := godotenv.Load(confFile); err != nil {
		log.Fatalf("Error loading %s: %v", confFile, err)
	}

	log.Printf("Config loaded: %s", confFile)
}

func main() {
	loadConfig()
	r := db.Connect()

	test.MigrateTest(r)
	//Tg.Tgbot()

	//Sc.View("Not Enough Animations")
}
