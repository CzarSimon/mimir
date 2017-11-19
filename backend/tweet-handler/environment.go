package main

import (
	"database/sql"

	"github.com/CzarSimon/util"
)

// Env is the struct for environment objects
type Env struct {
	DB      *sql.DB
	Config  Config
	Tickers TickerSet
	Aliases Aliases
}

// SetupEnv Sets up handler environment
func SetupEnv(config Config) *Env {
	db := util.ConnectPG(config.db)
	tickers, err := GetTickers(db)
	util.CheckErrFatal(err)
	aliases, err := GetAliases(db)
	util.CheckErrFatal(err)
	return &Env{
		DB:      db,
		Config:  config,
		Tickers: tickers,
		Aliases: aliases,
	}
}

// GetTickers Gets tickers to track
func GetTickers(db *sql.DB) (TickerSet, error) {
	tickers := make(TickerSet)
	rows, err := db.Query("SELECT TICKER FROM STOCKS")
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
		err = tickers.Add(ticker)
		if err != nil {
			util.LogErr(err)
		}
	}
	return tickers, nil
}

// GetAliases Gets aliases to track
func GetAliases(db *sql.DB) (Aliases, error) {
	aliases := make(Aliases)
	rows, err := db.Query("SELECT ALIAS, TICKER FROM TICKERALIASES")
	defer rows.Close()
	if err != nil {
		return aliases, err
	}
	var alias, ticker string
	for rows.Next() {
		err = rows.Scan(&alias, &ticker)
		if err != nil {
			return aliases, err
		}
		err = aliases.Add(alias, ticker)
		if err != nil {
			util.LogErr(err)
		}
	}
	return aliases, nil
}
