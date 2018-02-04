package main

import (
	"database/sql"
	"log"
	"strings"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
)

// Env holds environmet information
type Env struct {
	Tickers []string
	PriceDB *sql.DB
	config  Config
}

// SetupEnv creates an environmet based on supplied configuration.
func SetupEnv(config Config) *Env {
	priceDB, err := config.PriceDB.Connect()
	util.CheckErrFatal(err)
	return &Env{
		Tickers: getTickers(config.TickerDB),
		PriceDB: priceDB,
		config:  config,
	}
}

// getTickers connects to ticker database and retrives the list of tracked tickers.
func getTickers(dbConfig endpoint.SQLConfig) []string {
	db, err := dbConfig.Connect()
	util.CheckErrFatal(err)
	defer db.Close()
	rows, err := db.Query("SELECT TICKER FROM STOCK")
	util.CheckErrFatal(err)
	defer rows.Close()
	return constructTickerList(rows)
}

// constructTickerList Creates a list of tickers to check prices for.
func constructTickerList(rows *sql.Rows) []string {
	tickers := make([]string, 0)
	var ticker string
	for rows.Next() {
		err := rows.Scan(&ticker)
		util.CheckErrFatal(err)
		tickers = append(tickers, ticker)
	}
	return tickers
}

// logTickers Logs the tracked tickers.
func logTickers(env *Env) {
	log.Printf("Tickers: [%s]", strings.Join(env.Tickers, " "))
}
