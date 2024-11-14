package rest

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

type RateHandler struct {
	service RateServiceInterface
	logger  *slog.Logger
}

func NewRateHandler(service RateServiceInterface, logger *slog.Logger) *RateHandler {
	return &RateHandler{service: service, logger: logger}
}

func (h *RateHandler) GetRates(c *gin.Context) {
	currenciesParam := c.Query("currencies")
	cryptocurrencies := strings.Split(currenciesParam, ",")

	rates, err := h.service.GetRates(c, cryptocurrencies)
	if err != nil {
		h.logger.Error("Ошибка при получении курсов", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Курсы успешно получены")
	c.JSON(200, rates)
}
