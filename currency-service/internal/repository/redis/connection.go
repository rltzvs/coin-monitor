package redis

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	logger *slog.Logger
}

func NewRedisClient(ctx context.Context, addr, password string, db int, logger *slog.Logger) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Error("Ошибка подключения к Redis", "error", err)
	} else {
		logger.Info("Подключение к Redis успешно")
	}

	return &RedisClient{client: client, logger: logger}
}
