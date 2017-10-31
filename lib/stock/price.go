package stock

import (
	"time"

	"github.com/FlashBoys/go-finance"
)

// Price Holds price information about a ticker
type Price struct {
	Ticker   string    `json:"ticker"`
	Price    float64   `json:"price"`
	Currency string    `json:"currency,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}

// QuoteToPrice Converts a finance.Quote to a Price given a date
func QuoteToPrice(quote finance.Quote, date time.Time) Price {
	lastTradePrice, _ := quote.LastTradePrice.Float64()
	return Price{
		Ticker:   quote.Symbol,
		Price:    lastTradePrice,
		Currency: quote.Currency,
		Date:     date,
	}
}
