package main

import (
	"database/sql"
	"fmt"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/CzarSimon/util"
)

// GetAndStoreLatetsPrices Retrives and stores the latest prices for the active list of tickers
func GetAndStoreLatestPrices(config Config) {
	fmt.Println("Running: GetAndStoreLatestPrices")
	prices, err := GetPrices(config)
	if err != nil {
		util.LogErr(err)
		return
	}
	db, err := util.ConnectPGErr(config.PriceDB)
	defer db.Close()
	if err != nil {
		util.LogErr(err)
		return
	}
	err = StoreLatestPrices(prices, db)
	if err != nil {
		util.LogErr(err)
		return
	}
	err = StoreHistoricalPrices(prices, db)
	if err != nil {
		util.LogErr(err)
		return
	}
}

// StoreLatestPrices Stores latest prices in database
func StoreLatestPrices(prices []stock.Price, db *sql.DB) error {
	notUpdatedPrices, err := updateLatestPrices(prices, db)
	if err != nil {
		util.LogErr(err)
	}
	if len(notUpdatedPrices) > 0 {
		err = insertLatestPrices(notUpdatedPrices, db)
	}
	return err
}

// updateLatestPrices Updates the latest prices of a supplied set of tickers
// returns a list of the prices which could not be updated
func updateLatestPrices(prices []stock.Price, db *sql.DB) ([]stock.Price, error) {
	notUpdatedPrices := make([]stock.Price, 0)
	stmt, err := db.Prepare("UPDATE LATEST_PRICE SET DATE_INSERTED=$1, PRICE=$2, CURRENCY=$3 WHERE TICKER=$4")
	defer stmt.Close()
	if err != nil {
		return notUpdatedPrices, err
	}
	for _, price := range prices {
		res, err := stmt.Exec(price.Date, price.Price, price.Currency, price.Ticker)
		if err != nil {
			util.LogErr(err)
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
func StoreHistoricalPrices(prices []stock.Price, db *sql.DB) error {
	return insertNewPrices("HISTORICAL_PRICE", prices, db)
}

// insertNewPrices Inserts new prices either as historical or latest price
func insertNewPrices(table string, prices []stock.Price, db *sql.DB) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (TICKER, PRICE, CURRENCY, DATE_INSERTED) VALUES ($1, $2, $3, $4)",
		table)
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, price := range prices {
		_, err = stmt.Exec(price.Ticker, price.Price, price.Currency, price.Date)
		if err != nil {
			util.LogErr(err)
		}
	}
	return err
}

func printPrice(price stock.Price) {
	fmt.Printf("ticker=%s price=%f currency=%s\n", price.Ticker, price.Price, price.Currency)
}
