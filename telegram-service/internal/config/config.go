package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database      DatabaseConfig
	TelegramToken string
	LogLevel      string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn("Не удалось загрузить .env файл", "error", err)
		return nil, err
	}

	config := &Config{
		TelegramToken: getEnv("TELEGRAM_TOKEN", ""),
		LogLevel:      getEnv("LOG_LEVEL", "info"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5436"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Dbname:   getEnv("DB_NAME", "postgres"),
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// func getEnvAsInt(key string, defaultValue int) int {
// 	valueStr := getEnv(key, "")
// 	if valueStr == "" {
// 		return defaultValue
// 	}
// 	value, err := strconv.Atoi(valueStr)
// 	if err != nil {
// 		slog.Warn("неверный формат для %s, используется значение по умолчанию %v", key, defaultValue)
// 		return defaultValue
// 	}
// 	return value
// }
