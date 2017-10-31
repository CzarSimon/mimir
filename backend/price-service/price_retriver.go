package main

import (
	"log"
	"time"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/CzarSimon/util"
	"github.com/FlashBoys/go-finance"
)

// GetAndStorePrices Retrives end of day prices for a fetched set of tickers
// and stores the result
func GetAndStorePrices(config Config) {
	if CheckIfNotBusinessDay(config.Timezone) {
		log.Println("Not a business day")
		return
	}
	tickers, err := GetTickers(config.TickerDB)
	if err != nil {
		util.LogErr(err)
		return
	}
	logTickers(tickers)
	prices, err := QueryPrices(tickers, config.Timezone)
	if err != nil {
		util.LogErr(err)
		return
	}
	err = StorePrices(prices, config.PriceDB)
	if err != nil {
		util.LogErr(err)
		return
	}
	log.Printf("Retrived and stored prices for %d tickers\n", len(tickers))
}

// CheckIfNotBusinessDay Checks if the day of the exchange date is a business day or not
func CheckIfNotBusinessDay(timezone string) bool {
	weekday := getCurrentExchangeDate(timezone).Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// QuoteToPrice Converts a finance.Quote to a Price given a date
func QuoteToPrice(quote finance.Quote, date time.Time) stock.Price {
	lastTradePrice, _ := quote.LastTradePrice.Float64()
	return stock.Price{
		Ticker: quote.Symbol,
		Price:  lastTradePrice,
		Date:   date,
	}
}

// QueryPrices Querys yahoo finance for prices for all supplied tickers
func QueryPrices(tickers []string, timezone string) ([]stock.Price, error) {
	prices := make([]stock.Price, 0)
	quotes, err := finance.GetQuotes(tickers)
	if err != nil {
		return prices, err
	}
	date := getCurrentExchangeDate(timezone)
	for _, quote := range quotes {
		prices = append(prices, QuoteToPrice(quote, date))
	}
	return prices, nil
}

// getCurrentExchangeDate Gets the current date of the Exacnages in New York
func getCurrentExchangeDate(timezone string) time.Time {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Now().UTC().Add(-4 * time.Hour)
	}
	return time.Now().In(location)
}

// StorePrices Connects to database and stores prices for queried tickers
func StorePrices(prices []stock.Price, dbConfig util.PGConfig) error {
	db := util.ConnectPG(dbConfig)
	defer db.Close()
	query := "INSERT INTO CLOSE_PRICE (TICKER, PRICE_DATE, PRICE) VALUES ($1, $2, $3)"
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	for _, price := range prices {
		_, err = stmt.Exec(price.Ticker, price.Date, price.Price)
		if err != nil {
			util.LogErr(err)
		}
	}
	return err
}
