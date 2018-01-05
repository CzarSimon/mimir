package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/CzarSimon/util"
)

// IexAPI API handler for the IEX trading api
type IexAPI struct {
	Timezone string
}

// NewIexAPI Creates new api handler
func NewIexAPI(timezone string) IexAPI {
	return IexAPI{
		Timezone: timezone,
	}
}

// QueryPrices Query IEX tarding for prices for all supplied tickes
func (api IexAPI) QueryPrices(tickers []string) ([]stock.Price, error) {
	prices := make([]stock.Price, 0)
	quotes, err := getIEXQuotes(tickers)
	if err != nil {
		return prices, err
	}
	date := getCurrentExchangeDate(api.Timezone)
	var price stock.Price
	for _, quote := range quotes {
		price = IEXQuoteToPrice(quote, "USD", date)
		if price.IsValid() {
			prices = append(prices, price)
		}
	}
	return prices, nil
}

// getIEXQuotes Makes get request to IEX API
func getIEXQuotes(tickers []string) ([]IEXQuote, error) {
	var quotes []IEXQuote
	res, err := http.Get(createIEXURL(tickers))
	if err != nil {
		return quotes, err
	}
	err = util.DecodeJSON(res.Body, &quotes)
	if err != nil {
		return quotes, err
	}
	return quotes, nil
}

// createIEXURL Creates URL to query IEX API for quotes
func createIEXURL(tickers []string) string {
	IEXEndpoint := "https://ws-api.iextrading.com/1.0"
	return fmt.Sprintf(
		"%s/stock/market/quote?symbols=%s", IEXEndpoint, strings.Join(tickers, ","))
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
		Ticker:      strings.ToUpper(quote.Symbol),
		Price:       quote.LatestPrice,
		PriceChange: quote.ChangePercent,
		Currency:    strings.ToUpper(currency),
		Date:        date,
	}
}
