package external

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"telegram-service/internal/entity"
	"time"
)

type RatesRepository struct {
	client  *http.Client
	logger  *slog.Logger
	baseURL string
}

var baseURL = "localhost:8080"

func NewCryptoRepository(logger *slog.Logger) *RatesRepository {
	return &RatesRepository{client: &http.Client{Timeout: 10 * time.Second}, logger: logger, baseURL: baseURL}
}

func (r *RatesRepository) GetRates(currencies []string) ([]entity.Rate, error) {
	url := r.baseURL + "/rates"
	if len(currencies) > 0 {
		url += "?currencies=" + strings.Join(currencies, ",")
	}

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("некорректный статус ответа: %d", resp.StatusCode)
	}

	var rates []entity.Rate
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании JSON: %w", err)
	}

	return rates, nil
}
