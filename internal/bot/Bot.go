package tgbot

import (
	"log"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var knopkiNaKlave = botApi.NewInlineKeyboardMarkup(
	botApi.NewInlineKeyboardRow(
		botApi.NewInlineKeyboardButtonData("тыкать сюда", "пошёл нахуй долбаеб"),
		botApi.NewInlineKeyboardButtonData("не тыкай сюда", "ладно, мне что тебя обижать?"),
	),
	botApi.NewInlineKeyboardRow(
		botApi.NewInlineKeyboardButtonData("чисто чекаю", "прости что бываю злым йоу"),
	),
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
	updateChannel := bot.GetUpdatesChan(checking)

	// дебагинг
	bot.Debug = true
	log.Printf("Зашло на акк %s", bot.Self.UserName)

	for update := range updateChannel {

		if update.Message != nil {
			msg := botApi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "открыть":
				msg.ReplyMarkup = knopkiNaKlave
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

		} else if update.CallbackQuery != nil {
			callback := botApi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				log.Panic(err)
			}

			msg := botApi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
