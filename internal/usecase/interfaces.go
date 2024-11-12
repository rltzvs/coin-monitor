package usecase

import "currency-service/internal/entity"

// RateRepository описывает операции с хранилищем для курса валют.
type RateRepository interface {
	SaveRate(rate []entity.Rate) error
	GetAllRates() ([]entity.Rate, error)
	GetRateByCryptocurrency(cryptocurrency string) (entity.Rate, error)
}

// ExternalAPIClient описывает операции взаимодействия с внешним API для получения курсов валют.
type ExternalAPIClient interface {
	FetchRates() ([]entity.Rate, error)
}
