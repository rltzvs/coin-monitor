package rest

import (
	"context"
	"currency-service/internal/entity"
)

type RateServiceInterface interface {
	GetRates(context context.Context, cryptocurrencies []string) ([]entity.Rate, error)
}
