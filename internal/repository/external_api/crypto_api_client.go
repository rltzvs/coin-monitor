package external_api

import (
	"currency-service/internal/entity"
	"log/slog"
	"net/http"
)

type ExternalAPIClient struct {
	client *http.Client
	logger *slog.Logger
}

func NewExternalAPIClient(apiKey string, logger *slog.Logger) *ExternalAPIClient {
	return &ExternalAPIClient{client: &http.Client{}, logger: logger}
}

func (e *ExternalAPIClient) FetchRates() ([]entity.Rate, error) {
	// Запрос к внешнему API
	return nil, nil
}
