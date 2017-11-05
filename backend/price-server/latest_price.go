package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

// GetLatestPrices Gets the latest prices for a supplied set of tickers
func (env *Env) GetLatestPrices(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrStatus(
			res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	tickers, err := parseTickersFromQuery(req)
	if err != nil {
		util.SendErrStatus(res, err, http.StatusBadRequest)
		return
	}
	prices, err := getLatestPricesFromDB(tickers, env.db)
	if err != nil {
		util.LogErr(err)
		util.SendErrRes(res, errors.New("Could not get prices"))
		return
	}
	jsonBody, err := json.Marshal(prices)
	if err != nil {
		util.LogErr(err)
		util.SendErrRes(res, errors.New("Could not get prices"))
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// getLatestPricesFromDB Gets latest prices for a supplied set of tickers from the database
func getLatestPricesFromDB(tickers []string, db *sql.DB) ([]stock.Price, error) {
	prices := make([]stock.Price, 0)
	query := "SELECT TICKER, PRICE, CURRENCY, PRICE_CHANGE, DATE_INSERTED FROM LATEST_PRICE WHERE TICKER=ANY($1)"
	rows, err := db.Query(query, pq.Array(tickers))
	defer rows.Close()
	if err != nil {
		return prices, err
	}
	var price stock.Price
	for rows.Next() {
		err = rows.Scan(
			&price.Ticker, &price.Price, &price.Currency, &price.PriceChange, &price.Date)
		if err != nil {
			return prices, err
		}
		prices = append(prices, price)
	}
	return prices, nil
}

// parseTickersFromQuery Parses ticker from query
func parseTickersFromQuery(req *http.Request) ([]string, error) {
	tickers := req.URL.Query()["ticker"]
	if len(tickers) < 1 {
		return tickers, errors.New("No tickers supplied")
	}
	return tickers, nil
}
