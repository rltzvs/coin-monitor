package main

import (
	"log"
	"log/slog"
	"telegram-service/internal/config"
	"telegram-service/internal/controller/telegram"
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
		logger.Error("Ошибка подключения к Postgres: %v", err)
	}
	defer db.Close()

	// Инициализация Telegram бота
	bot, err := telegram.NewTelegramBot(config.TelegramToken)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot: %v", err)
	}

	bot.Run()
}
