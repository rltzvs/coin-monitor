package rest

import (
	"context"
	"currency-service/internal/entity"
	"time"
)

type RateServiceInterface interface {
	UpdateRates() error
	StartRateUpdater(interval time.Duration)
	GetRates(context context.Context, cryptocurrencies []string) ([]entity.Rate, error)
}
