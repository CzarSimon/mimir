package main

import (
	"database/sql"
	"log"
	"strings"

	"github.com/CzarSimon/util"
)

// GetTickers Retrives a set of tickers
func GetTickers(config util.PGConfig) ([]string, error) {
	db, err := util.ConnectPGErr(config)
	defer db.Close()
	if err != nil {
		return make([]string, 0), err
	}
	return queryForTickers(db)
}

// queryForTickers Querys for all tickers in a supplied database
func queryForTickers(db *sql.DB) ([]string, error) {
	tickers := make([]string, 0)
	rows, err := db.Query("SELECT TICKER FROM STOCK")
	defer rows.Close()
	if err != nil {
		return tickers, err
	}
	var ticker string
	for rows.Next() {
		err = rows.Scan(&ticker)
		if err != nil {
			return tickers, err
		}
		tickers = append(tickers, ticker)
	}
	return tickers, nil
}

// logTickers Logs the tickes for which prices are retrived
func logTickers(tickers []string) {
	log.Println(strings.Join(tickers, " "))
}
