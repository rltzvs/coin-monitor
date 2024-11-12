package entity

import "time"

type Rate struct {
	ID                uint
	Name              string
	Price             float64
	MinPriceDay       float64
	MaxPriceDay       float64
	Percent_change_1h float64
	LastUpdated       time.Time
	Created_at        time.Time
}
