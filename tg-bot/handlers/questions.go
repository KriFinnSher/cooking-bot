package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func QuestionsHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, logger *zap.Logger) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID

	msg := tgbotapi.NewMessage(chatID, "Здесь собраны ответы на самые частые вопросы. Выберите вопрос:")

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Как посмотреть мои маршруты с сайта?", "q1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кто автор сайта и бота?", "q2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Расскажи про сайт и бота", "q3"),
		),
	)

	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}
