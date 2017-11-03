package main

import (
	"errors"
	"log"

	"github.com/CzarSimon/mimir/lib/stock"
)

// GetPrices Retrives the active list of tickers and queris prices for them
func GetPrices(config Config, api PriceAPI) ([]stock.Price, error) {
	emptyPrices := make([]stock.Price, 0)
	if CheckIfNotBusinessDay(config.Timezone) {
		return emptyPrices, errors.New("Not a business day")
	}
	tickers, err := GetTickers(config.TickerDB)
	if err != nil {
		return emptyPrices, err
	}
	log.Println(tickers)
	return api.QueryPrices(tickers)
}

// PriceAPI Interface for price API
type PriceAPI interface {
	QueryPrices(tickers []string) ([]stock.Price, error)
}
