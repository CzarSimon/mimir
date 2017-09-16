package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CzarSimon/util"
	"github.com/FlashBoys/go-finance"
	r "gopkg.in/gorethink/gorethink.v3"
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

// logTickers Logs the tickes for which prices are retrived
func logTickers(tickers []string) {
	log.Println(strings.Join(tickers, " "))
}

// GetTickers Retrives a set of tickers
func GetTickers(rdbConfig RDBConfig) ([]string, error) {
	tickers := make([]string, 0)
	session, err := connectRDB(rdbConfig)
	defer session.Close()
	if err != nil {
		return tickers, err
	}
	rows, err := r.Table("stocks").Pluck("ticker").Run(session)
	defer rows.Close()
	if err != nil {
		return tickers, err
	}
	var ticker struct {
		Ticker string
	}
	for rows.Next(&ticker) {
		tickers = append(tickers, ticker.Ticker)
	}
	return tickers, nil
}

// connectRDB Connects to a rethinkdb server based on the supplied config parameters
func connectRDB(config RDBConfig) (*r.Session, error) {
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	session, err := r.Connect(r.ConnectOpts{
		Address:  address,
		Database: config.DB,
	})
	if err != nil {
		return session, err
	}
	return session, nil
}

// Price Holds pricing info for a ticker at a given date
type Price struct {
	Ticker string
	Price  float64
	Date   time.Time
}

// QuoteToPrice Converts a finance.Quote to a Price given a date
func QuoteToPrice(quote finance.Quote, date time.Time) Price {
	price, _ := quote.LastTradePrice.Float64()
	return Price{
		Ticker: quote.Symbol,
		Price:  price,
		Date:   date,
	}
}

// QueryPrices Querys yahoo finance for prices for all supplied tickers
func QueryPrices(tickers []string, timezone string) ([]Price, error) {
	prices := make([]Price, 0)
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
func StorePrices(prices []Price, dbConfig util.PGConfig) error {
	db := util.ConnectPG(dbConfig)
	defer db.Close()
	query := "INSERT INTO HISTORICAL_PRICE (TICKER, PRICE_DATE, PRICE) VALUES ($1, $2, $3)"
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
