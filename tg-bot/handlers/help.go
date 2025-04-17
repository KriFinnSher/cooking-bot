package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func BB(bot *tgbotapi.BotAPI, update tgbotapi.Update, logger *zap.Logger) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID

	message := "Я помогу вам в работе с вашими маршрутами!\n\n" +
		"Вот доступные команды:\n\n" +
		"/start - Приветственное сообщение и информация о боте.\n" +
		"/get - посмотреть все маршруты с сайта.\n" +
		"/questions - Ответы на часто задаваемые вопросы.\n" +
		"/help - Показать эту справку.\n\n"

	bot.Send(tgbotapi.NewMessage(chatID, message))
}
