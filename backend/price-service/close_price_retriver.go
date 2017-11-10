package main

import (
	"log"
	"time"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/CzarSimon/util"
)

// GetAndStoreClosePrices Retrives end of day prices for a fetched set of tickers
// and stores the result
func (env *Env) GetAndStoreClosePrices() {
	log.Println("Running: GetAndStoreClosePrices")
	prices, err := env.GetPrices()
	if err != nil {
		util.LogErr(err)
		return
	}
	err = StoreClosePrices(prices, env.Config.PriceDB)
	if err != nil {
		util.LogErr(err)
		return
	}
	log.Printf("Retrived and stored prices for %d tickers\n", len(prices))
}

// CheckIfNotBusinessDay Checks if the day of the exchange date is a business day or not
func CheckIfNotBusinessDay(timezone string) bool {
	weekday := getCurrentExchangeDate(timezone).Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// getCurrentExchangeDate Gets the current date of the Exacnages in New York
func getCurrentExchangeDate(timezone string) time.Time {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Now().UTC().Add(-4 * time.Hour)
	}
	return time.Now().In(location)
}

// StoreClosePrices Connects to database and stores prices for queried tickers
func StoreClosePrices(prices []stock.Price, dbConfig util.PGConfig) error {
	db, err := util.ConnectPGErr(dbConfig)
	defer db.Close()
	if err != nil {
		return err
	}
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
