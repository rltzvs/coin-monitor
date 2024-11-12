package redis

import (
	"currency-service/internal/entity"
	"log/slog"

	"github.com/go-redis/redis"
)

type RedisRateRepository struct {
	client *redis.Client
	logger *slog.Logger
}

func NewRedisRateRepository(client *redis.Client, logger *slog.Logger) *RedisRateRepository {
	return &RedisRateRepository{client: client, logger: logger}
}

func (r *RedisRateRepository) SaveRate(rate []entity.Rate) error {
	// Реализация сохранения курсов
	return nil
}

func (r *RedisRateRepository) GetAllRates() ([]entity.Rate, error) {
	// Реализация получения всех курсов
	return nil, nil
}

func (r *RedisRateRepository) GetRateByCryptocurrency(cryptocurrency string) (entity.Rate, error) {
	// Реализация получения курса по криптовалюте
	return entity.Rate{}, nil
}
