package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/query"
)

const defaultUntrackedTickerLimit = 10

// untrackedTickerHandler handles request for the resource untracked tickers.
func (env *Env) untrackedTickerHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return env.getUntrackedTickers(w, r)
	default:
		return httputil.MethodNotAllowed
	}
}

// getUntrackedTickers gets a list of untracked tickers.
func (env *Env) getUntrackedTickers(w http.ResponseWriter, r *http.Request) error {
	tickers, err := queryUntrackedTickers(getUntrackedTickersLimit(r), env.TweetDB)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, tickers)
}

// queryUntrackedTickers querys a database for untracked tickers
func queryUntrackedTickers(limit int, db *sql.DB) ([]string, error) {
	rows, err := db.Query(topUntrackedQuery(), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return constructUntrackedTickers(rows)
}

// constructUntrackedTickersList constructs a list of untracked tickers.
func constructUntrackedTickers(rows *sql.Rows) ([]string, error) {
	tickers := make([]string, 0)
	var ticker string
	for rows.Next() {
		err := rows.Scan(&ticker)
		if err != nil {
			return nil, err
		}
		tickers = append(tickers, ticker)
	}
	return tickers, nil
}

// topUntrackedQuery creates the query for getting top untracked tickers.
func topUntrackedQuery() string {
	return `
		SELECT TICKER
		FROM UNTRACKED_TICKERS
		WHERE TICKER NOT IN (SELECT TICKER FROM STOCKS WHERE IS_TRACKED=TRUE)
		GROUP BY TICKER ORDER BY COUNT(*) DESC LIMIT $1`
}

// getUntrackedTickersLimit gets the maximum number of
// untracked tickers that should be fetched.
func getUntrackedTickersLimit(r *http.Request) int {
	value, err := query.ParseValue(r, "limit")
	if err != nil {
		return defaultUntrackedTickerLimit
	}
	limit, err := strconv.Atoi(value)
	if err != nil {
		return defaultUntrackedTickerLimit
	}
	return limit
}
