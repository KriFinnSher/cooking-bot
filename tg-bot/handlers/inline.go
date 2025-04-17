package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func InlineButtonHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, logger *zap.Logger) {
	if update.CallbackQuery == nil {
		return
	}

	chatID := update.CallbackQuery.Message.Chat.ID
	callbackData := update.CallbackQuery.Data

	var response string

	switch callbackData {
	case "q1":
		response = "Чтобы посмотреть ваши машруты, используйте команду /get. В каждом новом сообщении вам отобразится каждый отдельный маршрут."
	case "q2":
		response = "Автор сайта и бота — Аверинов Вадим. Он создал этот сайт и бота, чтобы помочь вам создавать и управлять рецептами."
	case "q3":
		response = "Этот сайт и бот предназначены для того, чтобы вы могли легко создавать, редактировать и управлять своими маршрутами. Просто следуйте инструкциям, и все будет готово!"
	default:
		response = "Извините, я не понял ваш запрос."
	}

	bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "Ответ отправлен"))
	bot.Send(tgbotapi.NewMessage(chatID, response))
}
