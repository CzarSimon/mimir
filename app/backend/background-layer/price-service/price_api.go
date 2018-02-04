package main

import (
	"errors"
	"log"
	"os"

	"github.com/CzarSimon/mimir/lib/stock"
)

// GetPrices Retrives the active list of tickers and queris prices for them
func (env *Env) GetPrices() ([]stock.Price, error) {
	emptyPrices := make([]stock.Price, 0)
	if CheckIfNotBusinessDay(env.Config.Timezone) {
		return emptyPrices, errors.New("Not a business day")
	}
	tickers := env.getTickers()
	log.Println(tickers)
	return env.API.QueryPrices(tickers)
}

// PriceAPI Interface for price API
type PriceAPI interface {
	QueryPrices(tickers []string) ([]stock.Price, error)
}

const (
	PriceAPIKey  = "PRICE_API"
	IexAPIName   = "IEX"
	YahooAPIName = "Yahoo"
)

// selectPriceAPI Selects a price api based on environmet variable PRICE_API
func selectPriceAPI(timezone string) PriceAPI {
	switch os.Getenv(PriceAPIKey) {
	case IexAPIName:
		return NewIexAPI(timezone)
	case YahooAPIName:
		return NewYahooAPI(timezone)
	default:
		return NewIexAPI(timezone)
	}
}
