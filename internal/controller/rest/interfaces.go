package rest

import (
	"currency-service/internal/entity"
	"time"
)

type RateServiceInterface interface {
	UpdateRates() error
	StartRateUpdater(interval time.Duration)
	GetAllRates() ([]entity.Rate, error)
	GetRateByCryptocurrency(cryptocurrency string) (entity.Rate, error)
}
