package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"cooking-bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

type Request struct {
	Title       string         `json:"title,omitempty"`
	Ingredients map[string]int `json:"ingredients,omitempty"`
	RecipeText  string         `json:"recipe_text,omitempty"`
}

type UserState struct {
	Step        int
	Recipe      Request
	CurrentIngr string
}

var (
	userStates = make(map[int64]*UserState)
	mu         sync.Mutex
)

func AA(bot *tgbotapi.BotAPI, update tgbotapi.Update, cfg *config.Config, logger *zap.Logger) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	text := update.Message.Text

	mu.Lock()
	state, exists := userStates[chatID]
	if !exists {
		state = &UserState{Step: 0, Recipe: Request{Ingredients: make(map[string]int)}}
		userStates[chatID] = state
	}
	mu.Unlock()

	switch state.Step {
	case 0:
		// Шаг 1: Получаем название рецепта
		state.Recipe.Title = text
		state.Step++
		bot.Send(tgbotapi.NewMessage(chatID, "Введите ингредиенты в формате: Название - Количество. Когда закончите, напишите 'готово'"))

	case 1:
		// Шаг 2: Получаем ингредиенты
		if strings.ToLower(text) == "готово" {
			state.Step++
			bot.Send(tgbotapi.NewMessage(chatID, "Введите текст рецепта"))
		} else {
			parts := strings.SplitN(text, " - ", 2)
			if len(parts) != 2 {
				bot.Send(tgbotapi.NewMessage(chatID, "Ошибка формата. Введите в виде 'Название - Количество'"))
				return
			}
			amount, err := strconv.Atoi(parts[1])
			if err != nil {
				bot.Send(tgbotapi.NewMessage(chatID, "Ошибка: количество должно быть числом"))
				return
			}
			state.Recipe.Ingredients[parts[0]] = amount
			bot.Send(tgbotapi.NewMessage(chatID, "Добавлено! Введите следующий ингредиент или 'готово'"))
		}

	case 2:
		// Шаг 3: Получаем текст рецепта
		state.Recipe.RecipeText = text
		body, _ := json.Marshal(state.Recipe)

		resp, err := http.Post(cfg.BackendURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			logger.Error("Ошибка запроса к бэкенду", zap.Error(err))
			bot.Send(tgbotapi.NewMessage(chatID, "Ошибка соединения с сервером"))
			return
		}
		defer resp.Body.Close()

		bot.Send(tgbotapi.NewMessage(chatID, "Рецепт успешно добавлен"))
		mu.Lock()
		delete(userStates, chatID)
		mu.Unlock()
	}
}
