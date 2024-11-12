package rest

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type RateHandler struct {
	service RateServiceInterface
	logger  *slog.Logger
}

func NewRateHandler(service RateServiceInterface, logger *slog.Logger) *RateHandler {
	return &RateHandler{service: service, logger: logger}
}

// Отдает все курсы валют
func (h *RateHandler) GetAllRates(c *gin.Context) {
}

// Отдает курс конкретной валюты
func (h *RateHandler) GetRateByCryptocurrency(c *gin.Context) {
}
