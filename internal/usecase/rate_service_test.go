package usecase

import (
	"context"
	"currency-service/internal/entity"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRateRepository struct {
	mock.Mock
}

func (mock *MockRateRepository) SaveRate(ctx context.Context, rates []entity.Rate) error {
	args := mock.Called(ctx, rates)
	return args.Error(0)
}

func (mock *MockRateRepository) GetRates(ctx context.Context, cryptocurrencies []string) ([]entity.Rate, error) {
	args := mock.Called(ctx, cryptocurrencies)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Rate), args.Error(1)
}

type MockExternalAPIClient struct {
	mock.Mock
}

func (m *MockExternalAPIClient) FetchRates() ([]entity.Rate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Rate), args.Error(1)
}

func mockRates() []entity.Rate {
	return []entity.Rate{
		{
			Name:            "Bitcoin",
			Symbol:          "Bitcoin",
			Price:           50000.0,
			MinPriceDay:     48000.0,
			MaxPriceDay:     52000.0,
			PercentChange1h: -0.5,
			LastUpdated:     time.Now(),
		},
		{
			Name:            "USDC",
			Symbol:          "USDC",
			Price:           4000.0,
			MinPriceDay:     3800.0,
			MaxPriceDay:     4200.0,
			PercentChange1h: 0.7,
			LastUpdated:     time.Now(),
		},
	}
}

func TestUpdatesRates(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mockRepo := new(MockRateRepository)
	mockAPI := new(MockExternalAPIClient)
	service := NewRateService(mockRepo, mockAPI, logger)

	mockRates := mockRates()

	mockAPI.On("FetchRates").Return(mockRates, nil)
	mockRepo.On("SaveRate", mock.Anything, mockRates).Return(nil)

	err := service.UpdateRates()

	assert.NoError(t, err)
	mockAPI.AssertCalled(t, "FetchRates")
	mockRepo.AssertCalled(t, "SaveRate", mock.Anything, mockRates)
}

func TestUpdateRates_ErrorInFetchRates(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mockRepo := new(MockRateRepository)
	mockAPI := new(MockExternalAPIClient)
	service := NewRateService(mockRepo, mockAPI, logger)

	mockAPI.On("FetchRates").Return(nil, errors.New("fetch error"))

	err := service.UpdateRates()

	assert.Error(t, err)
	mockAPI.AssertCalled(t, "FetchRates")
	mockRepo.AssertNotCalled(t, "SaveRate", mock.Anything, mock.Anything)
}

func TestGetRates(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mockRepo := new(MockRateRepository)
	mockAPI := new(MockExternalAPIClient)
	service := NewRateService(mockRepo, mockAPI, logger)

	inputCryptocurrencies := []string{"USDC", "Bitcoin"}

	mockRates := mockRates()

	mockRepo.On("GetRates", mock.Anything, inputCryptocurrencies).Return(mockRates, nil)

	rates, err := service.GetRates(context.Background(), inputCryptocurrencies)

	assert.NoError(t, err)
	assert.Equal(t, mockRates, rates)
	mockRepo.AssertCalled(t, "GetRates", mock.Anything, inputCryptocurrencies)
}

func TestGetRatesError(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mockRepo := new(MockRateRepository)
	mockAPI := new(MockExternalAPIClient)
	service := NewRateService(mockRepo, mockAPI, logger)

	inputCryptocurrencies := []string{"USDC", "Bitcoin"}

	mockRepo.On("GetRates", mock.Anything, inputCryptocurrencies).Return(nil, errors.New("fetch error"))

	_, err := service.GetRates(context.Background(), inputCryptocurrencies)

	assert.Error(t, err)
	mockRepo.AssertCalled(t, "GetRates", mock.Anything, inputCryptocurrencies)
}
