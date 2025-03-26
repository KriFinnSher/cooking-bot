package main

import (
	"cooking-bot/config"
	"cooking-bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg := config.LoadConfig(logger)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		logger.Fatal("Ошибка при подключении к Telegram API", zap.Error(err))
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.Fatal("Ошибка при получении обновлений", zap.Error(err))
	}

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				handlers.StartHandler(bot, update, logger)
			case "create":
				handlers.AA(bot, update, cfg, logger)
			case "questions":
				handlers.QuestionsHandler(bot, update, logger)
			case "help":
				handlers.BB(bot, update, logger)
			}
		}

		if update.CallbackQuery != nil {
			handlers.InlineButtonHandler(bot, update, logger)
		}
	}
}
