package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

// Response описывает структуру данных, получаемых от API
type Response struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Ingredients map[string]int `json:"ingredients"`
	RecipeText  string         `json:"recipe_text"`
}

// GetHandler обрабатывает команду /get, получает список рецептов и отправляет пользователю
func GetHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, logger *zap.Logger) {
	apiURL := "http://localhost:8080/api/recipes/all/"

	// Отправляем GET-запрос
	resp, err := http.Get(apiURL)
	if err != nil {
		logger.Error("Ошибка при запросе к API", zap.Error(err))
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка при получении маршрутов 😢")
		bot.Send(msg)
		return
	}
	defer resp.Body.Close()

	// Декодируем JSON-ответ
	var recipes []Response
	if err := json.NewDecoder(resp.Body).Decode(&recipes); err != nil {
		logger.Error("Ошибка при декодировании JSON", zap.Error(err))
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при обработке данных 😢")
		bot.Send(msg)
		return
	}

	// Если рецептов нет
	if len(recipes) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Маршрутов пока нет 😕")
		bot.Send(msg)
		return
	}

	// Формируем ответное сообщение
	var messages []string
	for _, recipe := range recipes {
		ingredients := make([]string, 0, len(recipe.Ingredients))
		for name, quantity := range recipe.Ingredients {
			ingredients = append(ingredients, fmt.Sprintf("- %s: %d", name, quantity))
		}

		message := fmt.Sprintf(
			"🍽 *%s*\n\n📋 *Интересные места:*\n%s\n\n📝 *Маршрут:*\n%s",
			recipe.Title,
			strings.Join(ingredients, "\n"),
			recipe.RecipeText,
		)
		messages = append(messages, message)
	}

	// Отправляем рецепты (по отдельности, чтобы избежать ограничения по длине)
	for _, msgText := range messages {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}
