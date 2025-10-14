package tgbot

import (
	"log"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func tgbot() {

	// токен бота
	var tokenBot string = "7833805808:AAHo8zQ_VT2tkbrbwu2-ud0OKzdkgDovpCs"

	// объект бота + ошибки при подключении
	bot, err := botApi.NewBotAPI(tokenBot)
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
