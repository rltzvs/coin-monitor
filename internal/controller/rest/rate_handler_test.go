package rest

import (
	"context"
	"currency-service/internal/entity"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRateService struct {
	mock.Mock
}

func (m *MockRateService) GetRates(ctx context.Context, cryptocurrencies []string) ([]entity.Rate, error) {
	args := m.Called(ctx, cryptocurrencies)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Rate), args.Error(1)
}

func mockRates() []entity.Rate {
	return []entity.Rate{
		{
			Name:            "USDC",
			Symbol:          "usdc",
			Price:           1.001,
			MinPriceDay:     0.996993,
			MaxPriceDay:     1.004,
			PercentChange1h: -0.06257648016699595,
			LastUpdated:     time.Date(2024, 11, 18, 10, 26, 44, 659000000, time.UTC),
		},
		{
			Name:            "Bitcoin",
			Symbol:          "btc",
			Price:           91945.0,
			MinPriceDay:     88774.0,
			MaxPriceDay:     92142.0,
			PercentChange1h: -0.21437404906796934,
			LastUpdated:     time.Date(2024, 11, 18, 10, 26, 40, 342000000, time.UTC),
		},
	}
}

func TestRateHandler_GetRates(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		mockResponse   []entity.Rate
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Успешное получение курсов валют",
			query:          "currencies=Bitcoin,USDC",
			mockResponse:   mockRates(),
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"name":"USDC","symbol":"usdc","current_price":1.001,"low_24h":0.996993,"high_24h":1.004,"price_change_percentage_1h_in_currency":-0.06257648016699595,"last_updated":"2024-11-18T10:26:44.659Z"}`,
		},
		{
			name:           "Ошибка получения курсов валют",
			query:          "currencies=Bitcoin",
			mockResponse:   nil,
			mockError:      errors.New("internal error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"internal error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockRateService)
			if tt.mockError == nil {
				mockService.On("GetRates", mock.Anything, mock.Anything).Return(tt.mockResponse, nil)
			} else {
				mockService.On("GetRates", mock.Anything, mock.Anything).Return(nil, tt.mockError)
			}

			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
			handler := NewRateHandler(mockService, logger)

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.GET("/rates", handler.GetRates)

			req := httptest.NewRequest(http.MethodGet, "/rates?"+tt.query, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)

			mockService.AssertExpectations(t)
		})
	}
}
