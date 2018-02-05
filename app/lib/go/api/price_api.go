package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
)

// PriceAPI interface for price API.
type PriceAPI interface {
	QueryPrices(tickers []string) ([]stock.Price, error)
}

// IsBusinessDay checks if the day of the exchange date is a business day or not.
func IsBusinessDay(timezone string) bool {
	weekday := GetCurrentExchangeDate(timezone).Weekday()
	return weekday != time.Saturday && weekday != time.Sunday
}

// GetCurrentExchangeDate gets the current date of the Exchanges in New York.
func GetCurrentExchangeDate(timezone string) time.Time {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Now().UTC().Add(-4 * time.Hour)
	}
	return time.Now().In(location)
}

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
	date := GetCurrentExchangeDate(api.Timezone)
	var price stock.Price
	for _, quote := range quotes {
		price = iexQuoteToPrice(quote, "USD", date)
		if price.IsValid() {
			prices = append(prices, price)
		}
	}
	return prices, nil
}

// getIEXQuotes Makes get request to IEX API
func getIEXQuotes(tickers []string) ([]iexQuote, error) {
	var quotes []iexQuote
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
type iexQuote struct {
	Symbol        string  `json:"symbol"`
	LatestPrice   float64 `json:"latestPrice"`
	ChangePercent float64 `json:"changePercent"`
}

// IEXQuoteToPrice Converts an IEXQuote to Price
func iexQuoteToPrice(quote iexQuote, currency string, date time.Time) stock.Price {
	return stock.Price{
		Ticker:      strings.ToUpper(quote.Symbol),
		Price:       quote.LatestPrice,
		PriceChange: quote.ChangePercent,
		Currency:    strings.ToUpper(currency),
		Date:        date,
	}
}
