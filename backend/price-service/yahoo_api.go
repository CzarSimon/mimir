package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/FlashBoys/go-finance"
)

// YahooAPI API handler for yahoo finance
type YahooAPI struct {
	Tickers  []string
	Timezone string
}

// NewYahooAPI Creates new api handler
func NewYahooAPI(timezone string) YahooAPI {
	return YahooAPI{
		Timezone: timezone,
	}
}

// QueryPrices Querys yahoo finance for prices for all supplied tickers
func (api YahooAPI) QueryPrices(tickers []string) ([]stock.Price, error) {
	prices := make([]stock.Price, 0)
	quotes, err := finance.GetQuotes(tickers)
	if err != nil {
		return prices, err
	}
	date := getCurrentExchangeDate(api.Timezone)
	var price stock.Price
	for _, quote := range quotes {
		price = stock.QuoteToPrice(quote, date)
		fmt.Println(price.ToString(), price.IsValid())
		if price.IsValid() {
			prices = append(prices, price)
		}
	}
	return prices, nil
}

// QuoteToPrice Converts a finance.Quote to a Price given a date
func QuoteToPrice(quote finance.Quote, date time.Time) stock.Price {
	lastTradePrice, _ := quote.LastTradePrice.Float64()
	return stock.Price{
		Ticker:   strings.ToUpper(quote.Symbol),
		Price:    lastTradePrice,
		Currency: strings.ToUpper(quote.Currency),
		Date:     date,
	}
}
