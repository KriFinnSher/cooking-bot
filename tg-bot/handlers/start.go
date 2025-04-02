package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func StartHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, logger *zap.Logger) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID

	message := "Добро пожаловать! Я бот для добавления рецептов.\n\n" +
		"Для начала работы используйте команду /get, чтобы посмотреть ваши рецепты. Я использую внешнее апи со своего сайта.\n\n" +
		"Если нужна помощь, напишите /help.\n\n" +
		"Если у вас есть вопросы, используйте команду /questions для получения ответов на самые частые вопросы."

	bot.Send(tgbotapi.NewMessage(chatID, message))
}
