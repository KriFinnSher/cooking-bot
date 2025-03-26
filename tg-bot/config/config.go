package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	TelegramToken string
	BackendURL    string
}

func LoadConfig(logger *zap.Logger) *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Не удалось загрузить .env файл, используются системные переменные")
	}

	return &Config{
		TelegramToken: getEnv("TELEGRAM_BOT_TOKEN", "", logger),
		BackendURL:    getEnv("BACKEND_URL", "", logger),
	}
}

func getEnv(key, defaultValue string, logger *zap.Logger) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logger.Warn("Используется значение по умолчанию", zap.String("key", key), zap.String("default", defaultValue))
	return defaultValue
}
