package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/CzarSimon/util"
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
	return QueryIEXPrices(tickers, config.Timezone)
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
		fmt.Println(price.ToString(), price.IsValid())
		if price.IsValid() {
			prices = append(prices, price)
		}
	}
	return prices, nil
}

// QueryIEXPrices Query IEX tarding for prices for all supplied tickes
func QueryIEXPrices(tickers []string, timezone string) ([]stock.Price, error) {
	prices := make([]stock.Price, 0)
	quotes, err := getIEXQuotes(tickers)
	if err != nil {
		return prices, err
	}
	date := getCurrentExchangeDate(timezone)
	var price stock.Price
	for _, quote := range quotes {
		price = IEXQuoteToPrice(quote, "USD", date)
		fmt.Println(price.ToString(), price.IsValid())
		if price.IsValid() {
			prices = append(prices, price)
		}
	}
	return prices, nil
}

func getIEXQuotes(tickers []string) ([]IEXQuote, error) {
	fmt.Println(createIEXURL(tickers))
	var quotes []IEXQuote
	res, err := http.Get(createIEXURL(tickers))
	if err != nil {
		return make([]IEXQuote, 0), err
	}
	err = util.DecodeJSON(res.Body, &quotes)
	if err != nil {
		return quotes, err
	}
	return quotes, nil
}

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
		Ticker:   strings.ToUpper(quote.Symbol),
		Price:    quote.LatestPrice,
		Currency: strings.ToUpper(currency),
		Date:     date,
	}
}
