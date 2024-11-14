package usecase

import (
	"context"
	"currency-service/internal/entity"
)

type RateRepository interface {
	SaveRate(ctx context.Context, rates []entity.Rate) error
	GetRates(ctx context.Context, cryptocurrencies []string) ([]entity.Rate, error)
}

type ExternalAPIClient interface {
	FetchRates() ([]entity.Rate, error)
}
