package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	token := os.Getenv("huita")
	if token == "" {

		log.Fatal("Ты не поставил токен бота. Его можно узнать у @g4mace")

	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {

			switch update.Message.Command() {

			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "куку ёпта)_)))")
				bot.Send(msg)

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "пошёл нахуй пидорас нормальную команду напиши")
				bot.Send(msg)

			}

		} else {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "не, чувак, я не "+update.Message.Text)
			bot.Send(msg)

		}

	}

}
