package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/query"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/lib/pq"
)

const TICKER_KEY = "ticker"

// HandleStockRequest Retrives stock information of a requested stock
func (env *Env) HandleStockRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.MethodNotAllowed
	}
	ticker, err := query.ParseValue(r, TICKER_KEY)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	stck, err := retriveStockInfo(ticker, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, stck)
}

// RetriveStockInfo Retieves stock information from database
func retriveStockInfo(ticker string, db *sql.DB) (stock.Stock, error) {
	query := getStockInfoQuery()
	var stck stock.Stock
	err := db.QueryRow(query, ticker).Scan(
		&stck.Ticker, &stck.Name, &stck.Description, &stck.ImageURL, &stck.Website)
	return stck, err
}

// getStockInfoQuery Returns a query for retriving stock info for a specific ticker
func getStockInfoQuery() string {
	return `SELECT
            TICKER, NAME, DESCRIPTION, IMAGE_URL, WEBSITE
          FROM STOCK
          WHERE TICKER = $1`
}

// HandleStocksRequest Retrives stock information for list of requested stocks
func (env *Env) HandleStocksRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.BadRequest
	}
	tickers, err := query.ParseValues(r, TICKER_KEY)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	stocks, err := retriveStocks(tickers, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, stocks)
}

// retriveStocks Gets a map of stock info for a suplied list of tickers
func retriveStocks(tickers stock.Tickers, db *sql.DB) (map[string]stock.Stock, error) {
	query := "SELECT TICKER, NAME FROM STOCK WHERE TICKER = ANY($1)"
	rows, err := db.Query(query, pq.Array(tickers))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stocks := make(map[string]stock.Stock)
	var stck stock.Stock
	for rows.Next() {
		err = rows.Scan(&stck.Ticker, &stck.Name)
		if err != nil {
			return nil, err
		}
		stocks[stck.Ticker] = stck
	}
	return stocks, nil
}
