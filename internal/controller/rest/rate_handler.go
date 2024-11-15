package rest

import (
	_ "currency-service/internal/entity"
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

// @Summary Получить курсы валют
// @Description Возвращает курсы валют на основе переданных параметров
// @Tags Курсы валют
// @Accept json
// @Produce json
// @Param currencies query string true "Список валют через запятую" example="USD,EUR,BTC"
// @Success 200 {array} entity.Rate "Возвращает список курсов валют" // Указываем полный путь до типа Rate
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /rates [get]
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
