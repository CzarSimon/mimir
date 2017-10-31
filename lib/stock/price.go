package stock

import (
	"time"
)

// Price Holds price information about a ticker
type Price struct {
	Ticker   string    `json:"ticker"`
	Price    float64   `json:"price"`
	Currency string    `json:"currency,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}
