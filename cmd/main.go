package main

import (
	"currency-service/internal/config"
	"currency-service/internal/logger"
	"log/slog"
)

func main() {
	// Настройка конфигурации
	config, err := config.LoadConfig()
	if err != nil {
		slog.Error("Ошибка загрузки конфигурации", "error", err)
	}

	// Настройка slog
	log := logger.NewLogger(config.LogLevel)
	log.Info("Конфигурация загружена", "config", config)

	// Подключение к Redis

	// Репозиторий и внешний API клиент

	// Создание сервиса и обработчиков

	// Запуск фонового обновления курсов

	// Настройка маршрутов Gin

	// Запуск HTTP-сервера
}
