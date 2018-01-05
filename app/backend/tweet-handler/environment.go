package main

import (
	"database/sql"

	"github.com/CzarSimon/mimir/lib/news"
	"github.com/CzarSimon/util"
)

// Env is the struct for environment objects
type Env struct {
	DB      *sql.DB
	Config  Config
	Tickers Tickers
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
func GetTickers(db *sql.DB) (Tickers, error) {
	rows, err := db.Query("SELECT TICKER, NAME FROM STOCKS")
	if err != nil {
		return Tickers{}, err
	}
	defer rows.Close()
	tickers, err := saveTickersFromQueryResult(rows)
	return tickers, err
}

// saveTickersFromQueryResult Saves query result in a TickerSet
func saveTickersFromQueryResult(rows *sql.Rows) (Tickers, error) {
	tickers := make(Tickers)
	var err error
	for rows.Next() {
		err = addRecivedTicker(rows, &tickers)
		if err != nil {
			return tickers, err
		}
	}
	return tickers, nil
}

// addRecivedTicker Scans ticker and name from a query result and stores in supplied Tickers
func addRecivedTicker(rows *sql.Rows, tickers *Tickers) error {
	var subject news.Subject
	err := rows.Scan(&subject.Ticker, &subject.Name)
	if err != nil {
		return err
	}
	err = tickers.Add(subject)
	if err != nil {
		util.LogErr(err)
	}
	return nil
}

// GetAliases Gets aliases to track
func GetAliases(db *sql.DB) (Aliases, error) {
	rows, err := db.Query("SELECT ALIAS, TICKER FROM TICKERALIASES")
	if err != nil {
		return Aliases{}, err
	}
	defer rows.Close()
	aliases, err := saveAliasesFromRows(rows)
	return aliases, err
}

// saveAliasesFromRows Saves aliases from query result in Aliases
func saveAliasesFromRows(rows *sql.Rows) (Aliases, error) {
	aliases := make(Aliases)
	for rows.Next() {
		err := addRecivedAlias(rows, &aliases)
		if err != nil {
			return aliases, err
		}
	}
	return aliases, nil
}

// addRecivedAlias Scans alias-ticker result and stores in supplied Aliases
func addRecivedAlias(rows *sql.Rows, aliases *Aliases) error {
	var alias, ticker string
	err := rows.Scan(&alias, &ticker)
	if err != nil {
		return err
	}
	err = aliases.Add(alias, ticker)
	if err != nil {
		util.LogErr(err)
	}
	return nil
}
