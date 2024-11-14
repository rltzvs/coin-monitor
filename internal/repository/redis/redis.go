package redis

import (
	"context"
	"currency-service/internal/entity"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRateRepository struct {
	client *redis.Client
	logger *slog.Logger
}

func NewRedisRateRepository(redisClient *RedisClient, logger *slog.Logger) *RedisRateRepository {
	return &RedisRateRepository{client: redisClient.client, logger: logger}
}

func (r *RedisRateRepository) SaveRate(ctx context.Context, rates []entity.Rate) error {
	for _, rate := range rates {
		data, err := json.Marshal(rate)
		if err != nil {
			r.logger.Error("Ошибка при маршалинге курса", "cryptocurrency", rate.Name, "error", err)
			return err
		}

		err = r.client.Set(ctx, rate.Name, data, 24*time.Hour).Err()
		if err != nil {
			r.logger.Error("Ошибка при сохранении курса в Redis", "cryptocurrency", rate.Name, "error", err)
			return err
		}
		r.logger.Info("Курс успешно сохранен в Redis", "cryptocurrency", rate.Name)
	}
	return nil
}

func (r *RedisRateRepository) GetRates(ctx context.Context, cryptocurrencies []string) ([]entity.Rate, error) {
	var rates []entity.Rate

	// Если список пустой, получаем все ключи
	if len(cryptocurrencies) == 0 {
		var err error
		cryptocurrencies, err = r.client.Keys(ctx, "*").Result()
		if err != nil {
			r.logger.Error("Ошибка при получении всех ключей из Redis", "error", err)
			return nil, err
		}
	}

	for _, name := range cryptocurrencies {
		data, err := r.client.Get(ctx, name).Result()
		if err == redis.Nil {
			r.logger.Warn("Курс не найден в Redis", "cryptocurrency", name)
			continue
		} else if err != nil {
			r.logger.Error("Ошибка при получении курса из Redis", "cryptocurrency", name, "error", err)
			return nil, err
		}

		var rate entity.Rate
		if err := json.Unmarshal([]byte(data), &rate); err != nil {
			r.logger.Error("Ошибка при декодировании JSON курса", "cryptocurrency", name, "error", err)
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}
