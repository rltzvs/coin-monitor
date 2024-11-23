package entity

import "time"

type Crypto struct {
	ID     int
	Symbol string
}

type Rate struct {
	Name            string    `json:"name"`
	Symbol          string    `json:"symbol"`
	Price           float64   `json:"current_price"`
	MinPriceDay     float64   `json:"low_24h"`
	MaxPriceDay     float64   `json:"high_24h"`
	PercentChange1h float64   `json:"price_change_percentage_1h_in_currency"`
	LastUpdated     time.Time `json:"last_updated"`
}
