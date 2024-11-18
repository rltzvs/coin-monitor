package external_api

import (
	"currency-service/internal/entity"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

const baseURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=1h"

type ExternalAPIClient struct {
	client  *http.Client
	logger  *slog.Logger
	baseURL string
}

func NewExternalAPIClient(logger *slog.Logger) *ExternalAPIClient {
	return &ExternalAPIClient{client: &http.Client{Timeout: 10 * time.Second}, logger: logger, baseURL: baseURL}
}

func (e *ExternalAPIClient) FetchRates() ([]entity.Rate, error) {
	resp, err := e.client.Get(e.baseURL)
	if err != nil {
		e.logger.Error("Ошибка при выполнении запроса к внешнему API", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := errors.New("не удалось получить данные: статус " + resp.Status)
		e.logger.Error("Некорректный статус от API", "status", resp.Status)
		return nil, err
	}

	var rates []entity.Rate
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		e.logger.Error("Ошибка декодирования JSON", "error", err)
		return nil, err
	}

	e.logger.Info("Успешно получены данные о курсах криптовалют", "кол-во валют", len(rates))
	return rates, nil
}
