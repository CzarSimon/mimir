package main

import (
	"database/sql"
	"log"
	"strings"

	endpoint "github.com/CzarSimon/go-endpoint"
)

// GetTickers Retrives a set of tickers
func GetTickers(dbConfig endpoint.SQLConfig) ([]string, error) {
	db, err := dbConfig.Connect()
	defer db.Close()
	if err != nil {
		return make([]string, 0), err
	}
	return queryForTickers(db)
}

// queryForTickers Querys for all tickers in a supplied database
func queryForTickers(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT TICKER FROM STOCK")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return constructTickerList(rows)
}

// constructTickerList Creates a list of tickers to check prices for.
func constructTickerList(rows *sql.Rows) ([]string, error) {
	tickers := make([]string, 0)
	var ticker string
	var err error
	for rows.Next() {
		err := rows.Scan(&ticker)
		if err != nil {
			log.Println(err)
			continue
		}
		tickers = append(tickers, ticker)
	}
	return tickers, err
}

// logTickers Logs the tickes for which prices are retrived
func logTickers(env *Env) {
	log.Printf("Tickers: [%s]\n", strings.Join(env.Tickers, " "))
}
