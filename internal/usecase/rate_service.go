package usecase

import (
	"currency-service/internal/entity"
	"currency-service/internal/repository/external_api"
	"currency-service/internal/repository/redis"
	"log/slog"
	"time"
)

type RateService struct {
	repo   redis.RedisRateRepository      // RateRepository
	api    external_api.ExternalAPIClient // ExternalAPIClient
	logger *slog.Logger
}

func NewRateService(repo redis.RedisRateRepository, api external_api.ExternalAPIClient, logger *slog.Logger) *RateService {
	return &RateService{
		repo:   repo,
		api:    api,
		logger: logger,
	}
}

// UpdateRates обновляет курсы валют с внешнего API и сохраняет их в хранилище.
func (s *RateService) UpdateRates() error {
	return nil
}

// StartRateUpdater запускает обновление курсов в фоновом режиме каждые N минут.
func (s *RateService) StartRateUpdater(interval time.Duration) {
}

// GetAllRates возвращает все курсы валют.
func (s *RateService) GetAllRates() ([]entity.Rate, error) {
	return s.repo.GetAllRates()
}

// GetRateByCryptocurrency возвращает курс конкретной валюты.
func (s *RateService) GetRateByCryptocurrency(cryptocurrency string) (entity.Rate, error) {
	return s.repo.GetRateByCryptocurrency(cryptocurrency)
}
