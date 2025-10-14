package tgbot

import (
	"log"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Tgbot(token string) {

	// объект бота + ошибки при подключении
	bot, err := botApi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//обновления в чате
	checking := botApi.NewUpdate(0)
	checking.Timeout = 60

	// канал с обновлениями идёт из bot в updates
	updateChannel, err := bot.GetUpdatesChan(checking)
	if err != nil {
		log.Panic(err)
	}

	//цикл с приёмом обновлений из канала
	for update := range updateChannel {
		if update.Message == nil {
			continue
		}

		msg := botApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		bot.Send(msg)
	}

}
