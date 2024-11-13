package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string
	RedisAddress   string
	UpdateInterval time.Duration
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn("Не удалось загрузить .env файл", "error", err)
	}

	config := &Config{
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		UpdateInterval: getEnvAsDuration("UPDATE_INTERVAL", 5*time.Minute),
		RedisAddress:   getEnv("REDIS_ADDRESS", "localhost:6379"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return value
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		slog.Warn("неверный формат для %s, используется значение по умолчанию %v", key, defaultValue)
		return defaultValue
	}
	return value
}
