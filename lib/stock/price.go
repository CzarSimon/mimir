package stock

import (
	"fmt"
	"strings"
	"time"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/FlashBoys/go-finance"
)

// Price Holds price information about a ticker
type Price struct {
	Ticker   string    `json:"ticker,omitempty"`
	Price    float64   `json:"price"`
	Currency string    `json:"currency,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}

// QuoteToPrice Converts a finance.Quote to a Price given a date
func QuoteToPrice(quote finance.Quote, date time.Time) Price {
	lastTradePrice, _ := quote.LastTradePrice.Float64()
	return Price{
		Ticker:   strings.ToUpper(quote.Symbol),
		Price:    lastTradePrice,
		Currency: strings.ToUpper(quote.Currency),
		Date:     date,
	}
}

// IsValid Checks if the content of a price is valid
func (price Price) IsValid() bool {
	return price.Price != 0.0 && price.Currency != "" && price.Ticker != ""
}

// ToString Creates a string of the contrents of a price
func (price Price) ToString() string {
	return fmt.Sprintf("Ticker=%s Price=%f Currency=%s", price.Ticker, price.Price, price.Currency)
}

// IEXQuote Stock quote from IEX trading API
type IEXQuote struct {
	Symbol        string  `json:"symbol"`
	LatestPrice   float64 `json:"latestPrice"`
	ChangePercent float64 `json:"changePercent"`
}

// IEXQuoteToPrice Converts an IEXQuote to Price
func IEXQuoteToPrice(quote IEXQuote, currency string, date time.Time) stock.Price {
	return stock.Price{
		Ticker:   strings.ToUpper(quote.Symbol),
		Price:    quote.LatestPrice,
		Currency: strings.ToUpper(currency),
		Date:     date,
	}
}
