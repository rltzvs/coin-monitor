package usecase

import (
	"context"
	"currency-service/internal/entity"
	"log/slog"
	"time"
)

type RateService struct {
	repo   RateRepository
	api    ExternalAPIClient
	logger *slog.Logger
}

func NewRateService(repo RateRepository, api ExternalAPIClient, logger *slog.Logger) *RateService {
	return &RateService{
		repo:   repo,
		api:    api,
		logger: logger,
	}
}

func (s *RateService) UpdateRates() error {
	s.logger.Info("Запуск UpdateRates")

	rates, err := s.api.FetchRates()
	if err != nil {
		s.logger.Error("Ошибка при обновлении курсов", "error", err)
		return err
	}

	if err := s.repo.SaveRate(context.Background(), rates); err != nil {
		s.logger.Error("Ошибка при сохранении курсов", "error", err)
		return err
	}

	s.logger.Info("Курсы успешно обновлены")
	return nil
}

func (s *RateService) StartRateUpdater(interval time.Duration) {
	s.logger.Info("Запуск StartRateUpdater", "interval", interval)

	if err := s.UpdateRates(); err != nil {
		s.logger.Error("Ошибка при первоначальном обновлении курсов", "error", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if err := s.UpdateRates(); err != nil {
			s.logger.Error("Ошибка при обновлении курсов", "error", err)
		}
	}
}

func (s *RateService) GetRates(context context.Context, cryptocurrencies []string) ([]entity.Rate, error) {
	s.logger.Info("Запуск GetRates", "cryptocurrencies", cryptocurrencies)
	rates, err := s.repo.GetRates(context, cryptocurrencies)
	if err != nil {
		s.logger.Error("Ошибка при получении курсов", "error", err)
		return nil, err
	}
	return rates, err
}
