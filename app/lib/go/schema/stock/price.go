package stock

import (
	"fmt"
	"time"
)

// Price Holds price information about a ticker
type Price struct {
	Ticker      string    `json:"ticker,omitempty"`
	Price       float64   `json:"price"`
	PriceChange float64   `json:"priceChange,omitempty"`
	Currency    string    `json:"currency,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

// IsValid Checks if the content of a price is valid
func (price Price) IsValid() bool {
	return price.Price != 0.0 && price.Currency != "" && price.Ticker != ""
}

// ToString Creates a string of the contrents of a price
func (price Price) ToString() string {
	return fmt.Sprintf("Ticker=%s Price=%f Currency=%s PriceChange=%s",
		price.Ticker, price.Price, price.Currency, price.PriceChange)
}
