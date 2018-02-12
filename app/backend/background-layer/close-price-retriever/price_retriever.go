package main

import (
	"database/sql"
	"log"

	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
)

// GetAndStoreClosePrices Retrives end of day prices for a fetched set of tickers
// and stores the result
func (env *Env) GetAndStoreClosePrices() error {
	prices, err := env.API.QueryPrices(env.Tickers)
	if err != nil {
		return err
	}
	return StoreClosePrices(prices, env.PriceDB)
}

// StoreClosePrices stores prices for queried tickers.
func StoreClosePrices(prices []stock.Price, db *sql.DB) error {
	query := "INSERT INTO CLOSE_PRICE (TICKER, PRICE_DATE, PRICE) VALUES ($1, $2, $3)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, price := range prices {
		_, err = stmt.Exec(price.Ticker, price.Date, price.Price)
		if err != nil {
			log.Println(err)
		}
	}
	return err
}
