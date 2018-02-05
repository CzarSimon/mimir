package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CzarSimon/mimir/app/lib/go/api"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
)

// GetAndStoreLatetsPrices Retrives and stores the latest prices for the active list of tickers
func (env *Env) GetAndStoreLatestPrices() {
	log.Println("Running: GetAndStoreLatestPrices")
	prices, err := env.getPrices()
	if err != nil {
		log.Println(err)
		return
	}
	db, err := env.Config.PriceDB.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	err = storePrices(prices, db)
	if err != nil {
		log.Println(err)
		return
	}
}

// getPrices gets latest prices if run an a business day.
func (env *Env) getPrices() ([]stock.Price, error) {
	if !api.IsBusinessDay(env.Config.Timezone) || !env.exchangeIsOpen() {
		return nil, fmt.Errorf("Exchange not open")
	}
	return env.API.QueryPrices(env.Tickers)
}

// storePrices stores queried prices in the supplied database.
func storePrices(prices []stock.Price, db *sql.DB) error {
	err := storeLatestPrices(prices, db)
	if err != nil {
		return err
	}
	err = storeHistoricalPrices(prices, db)
	return err
}

// StoreLatestPrices stores latest prices in database.
func storeLatestPrices(prices []stock.Price, db *sql.DB) error {
	notUpdatedPrices, err := updateLatestPrices(prices, db)
	if err != nil {
		log.Println(err)
	}
	if len(notUpdatedPrices) > 0 {
		err = insertLatestPrices(notUpdatedPrices, db)
	}
	return err
}

// updateLatestPrices updates the latest prices of a supplied set of tickers
// returns a list of the prices which could not be updated.
func updateLatestPrices(prices []stock.Price, db *sql.DB) ([]stock.Price, error) {
	stmt, err := db.Prepare(getUpdatePricesQuery())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	notUpdatedPrices := make([]stock.Price, 0)
	for _, price := range prices {
		res, err := stmt.Exec(price.Date, price.Price, price.PriceChange, price.Currency, price.Ticker)
		if err != nil {
			log.Println(err)
			notUpdatedPrices = append(notUpdatedPrices, price)
			continue
		}
		noRowsAffected, err := res.RowsAffected()
		if err != nil || noRowsAffected == 0 {
			notUpdatedPrices = append(notUpdatedPrices, price)
		}
	}
	return notUpdatedPrices, err
}

// insertLatestPrices Inserts the latest prices of a supplied set of tickers
func insertLatestPrices(prices []stock.Price, db *sql.DB) error {
	return insertNewPrices("LATEST_PRICE", prices, db)
}

// StoreHistoricalPrices Stores historical prices in database
func storeHistoricalPrices(prices []stock.Price, db *sql.DB) error {
	return insertNewPrices("HISTORICAL_PRICE", prices, db)
}

// insertNewPrices Inserts new prices either as historical or latest price
func insertNewPrices(table string, prices []stock.Price, db *sql.DB) error {
	stmt, err := db.Prepare(getInsertPricesQuery(table))
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, price := range prices {
		_, err = stmt.Exec(price.Ticker, price.Price, price.Currency, price.Date)
		if err != nil {
			log.Println(err)
		}
	}
	return err
}

// getInsertPricesQuery creates an insert query in a supplied table.
func getInsertPricesQuery(table string) string {
	return fmt.Sprintf(
		"INSERT INTO %s (TICKER, PRICE, CURRENCY, DATE_INSERTED) VALUES ($1, $2, $3, $4)", table)
}

// getUpdatePricesQuery creates an update query for the latset price query.
func getUpdatePricesQuery() string {
	return `UPDATE LATEST_PRICE
						SET DATE_INSERTED=$1, PRICE=$2, PRICE_CHANGE=$3, CURRENCY=$4
						WHERE TICKER=$5`
}

// exchangeIsOpen checks if the current time is withing the exchange opening hours.
func (env *Env) exchangeIsOpen() bool {
	conf := env.Config
	hour := api.GetCurrentExchangeDate(env.Config.Timezone).Hour()
	return hour >= conf.ExchangeOpenHour && hour < conf.ExchangeCloseHour
}
