package main

import (
	"errors"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/FlashBoys/go-finance"
)

// GetPrices Retrives the active list of tickers and queris prices for them
func GetPrices(config Config) ([]stock.Price, error) {
	emptyPrices := make([]stock.Price, 0)
	if CheckIfNotBusinessDay(config.Timezone) {
		return emptyPrices, errors.New("Not a business day")
	}
	tickers, err := GetTickers(config.TickerDB)
	if err != nil {
		return emptyPrices, err
	}
	logTickers(tickers)
	return QueryYahooPrices(tickers, config.Timezone)
}

// QueryPrices Querys yahoo finance for prices for all supplied tickers
func QueryYahooPrices(tickers []string, timezone string) ([]stock.Price, error) {
	prices := make([]stock.Price, 0)
	quotes, err := finance.GetQuotes(tickers)
	if err != nil {
		return prices, err
	}
	date := getCurrentExchangeDate(timezone)
	var price stock.Price
	for _, quote := range quotes {
		price = stock.QuoteToPrice(quote, date)
		if price.IsValid() {
			prices = append(prices, price)
		}
	}
	return prices, nil
}
