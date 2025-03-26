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

	message := "Я помогу вам в работе с рецептами!\n\n" +
		"Вот доступные команды:\n\n" +
		"/start - Приветственное сообщение и информация о боте.\n" +
		"/create - Добавить новый рецепт.\n" +
		"/questions - Ответы на часто задаваемые вопросы.\n" +
		"/help - Показать эту справку.\n\n" +
		"Для создания рецепта используйте команду /create. Я буду пошагово собирать информацию от вас, включая название, ингредиенты и текст рецепта.\n\n" +
		"Если вам нужно получить ответы на вопросы, используйте /questions."

	bot.Send(tgbotapi.NewMessage(chatID, message))
}
