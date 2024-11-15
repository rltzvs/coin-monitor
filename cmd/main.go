package main

import (
	"currency-service/internal/config"
	"currency-service/internal/controller/rest"
	"currency-service/internal/logger"
	"currency-service/internal/repository/external_api"
	"currency-service/internal/repository/redis"
	"currency-service/internal/usecase"
	"sync"
	"time"

	"context"
	"log/slog"

	_ "currency-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Currency API
// @version 1.0
// @description API для получения курсов валют

func main() {
	// Настройка конфигурации
	config, err := config.LoadConfig()
	if err != nil {
		slog.Error("Ошибка загрузки конфигурации", "error", err)
	}

	// Настройка slog
	logger := logger.NewLogger(config.LogLevel)
	logger.Info("Конфигурация загружена", "config", config)

	// Подключение к Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	redisClient := redis.NewRedisClient(ctx, config.RedisAddress, config.RedisPassword, config.RedisDB, logger)

	// Репозиторий и внешний API клиент
	rateRepo := redis.NewRedisRateRepository(redisClient, logger)
	apiClient := external_api.NewExternalAPIClient(logger)

	// Создание сервиса и обработчиков
	rateService := usecase.NewRateService(rateRepo, apiClient, logger)
	rateHandler := rest.NewRateHandler(rateService, logger)

	// Запуск фонового обновления курсов
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		rateService.StartRateUpdater(config.UpdateInterval)
	}()
	// Настройка маршрутов Gin
	router := gin.Default()
	router.GET("/rates", func(c *gin.Context) {
		rateHandler.GetRates(c)
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск HTTP-сервера
	logger.Error("Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		logger.Error("Ошибка запуска HTTP-сервера", "error", err)
	}

	wg.Wait()
}
