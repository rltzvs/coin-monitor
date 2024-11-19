package config

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string
	RedisAddress   string
	RedisPassword  string
	RedisDB        int
	UpdateInterval time.Duration
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn("Не удалось загрузить .env файл", "error", err)
		return nil, err
	}

	config := &Config{
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		UpdateInterval: getEnvAsDuration("UPDATE_INTERVAL", 5*time.Minute),
		RedisAddress:   getEnv("REDIS_ADDRESS", "localhost:6379"),
		RedisPassword:  getEnv("REDIS_PASSWORD", ""),
		RedisDB:        getEnvAsInt("REDIS_DB", 0),
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

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		slog.Warn("неверный формат для %s, используется значение по умолчанию %v", key, defaultValue)
		return defaultValue
	}
	return value
}
