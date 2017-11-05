package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

// Stock Holds information about a stock
type Stock struct {
	Ticker      string `json:"ticker"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty"`
	Website     string `json:"website,omitempty"`
}

// HandleStockRequest Retrives stock information of a requested stock
func (env *Env) HandleStockRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrStatus(res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	ticker, err := parseTickerFromQuery(req)
	if err != nil {
		util.SendErrStatus(res, err, http.StatusBadRequest)
		return
	}
	stock, err := retriveStockInfo(ticker, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonRes, err := json.Marshal(&stock)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonRes)
}

// parseTickerFromQuery Parses request for a ticker
func parseTickerFromQuery(req *http.Request) (string, error) {
	return util.ParseValueFromQuery(req, "ticker", "No ticker supplied")
}

// RetriveStockInfo Retieves stock information from database
func retriveStockInfo(ticker string, db *sql.DB) (Stock, error) {
	query := getStockInfoQuery()
	var stock Stock
	err := db.QueryRow(query, ticker).Scan(
		&stock.Ticker, &stock.Name, &stock.Description, &stock.ImageURL, &stock.Website)
	if err != nil {
		return stock, err
	}
	return stock, nil
}

// getStockInfoQuery Returns a query for retriving stock info for a specific ticker
func getStockInfoQuery() string {
	return `SELECT
            TICKER, NAME, DESCRIPTION, IMAGE_URL, WEBSITE
          FROM STOCK
          WHERE TICKER = $1`
}

// HandleStocksRequest Retrives stock information for list of requested stocks
func (env *Env) HandleStocksRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrStatus(res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	tickers, err := parseTickersFromQuery(req)
	if err != nil {
		util.LogErr(err)
		util.SendErrStatus(res, err, http.StatusBadRequest)
		return
	}
	stocks, err := retriveStocks(tickers, env.db)
	if err != nil {
		util.LogErr(err)
		util.SendErrRes(res, errors.New("Could not retrieve stocks"))
		return
	}
	jsonBody, err := json.Marshal(stocks)
	if err != nil {
		util.LogErr(err)
		util.SendErrRes(res, errors.New("Could not retrieve stocks"))
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// retriveStocks Gets a map of stock info for a suplied list of tickers
func retriveStocks(tickers Tickers, db *sql.DB) (map[string]Stock, error) {
	stocks := make(map[string]Stock)
	query := "SELECT TICKER, NAME FROM STOCK WHERE TICKER = ANY($1)"
	rows, err := db.Query(query, pq.Array(tickers))
	defer rows.Close()
	if err != nil {
		return stocks, err
	}
	var stock Stock
	for rows.Next() {
		err = rows.Scan(&stock.Ticker, &stock.Name)
		if err != nil {
			return stocks, err
		}
		stocks[stock.Ticker] = stock
	}
	return stocks, nil
}
