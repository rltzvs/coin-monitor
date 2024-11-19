package main

import (
	"log/slog"
	"telegram-service/internal/config"
	"telegram-service/internal/logger"
	"telegram-service/internal/repository/postgres"
)

func main() {
	// Настройка конфигурации
	config, err := config.LoadConfig()
	if err != nil {
		slog.Error("Ошибка загрузки конфигурации", "error", err)
	}

	// Настройка slog
	logger := logger.NewLogger(config.LogLevel)
	logger.Info("Конфигурация загружена", "config", config)

	// Подключение к Postgres
	db, err := postgres.NewDBConnection(&config.Database)
	if err != nil {
		slog.Error("Ошибка подключения к Postgres: %v", err)
	}
	defer db.Close()

	// Инициализация Telegram бота

	// Запуск Telegram бота
}
