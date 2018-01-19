package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CzarSimon/httputil/query"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

const TICKER_KEY = "ticker"

// HandleStockRequest Retrives stock information of a requested stock
func (env *Env) HandleStockRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrStatus(res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	ticker, err := query.ParseValue(req, TICKER_KEY)
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
func (env *Env) HandleStocksRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrStatus(res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	tickers, err := query.ParseValues(req, TICKER_KEY)
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
func retriveStocks(tickers stock.Tickers, db *sql.DB) (map[string]stock.Stock, error) {
	stocks := make(map[string]stock.Stock)
	query := "SELECT TICKER, NAME FROM STOCK WHERE TICKER = ANY($1)"
	rows, err := db.Query(query, pq.Array(tickers))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stck stock.Stock
	for rows.Next() {
		err = rows.Scan(&stck.Ticker, &stck.Name)
		if err != nil {
			return stocks, err
		}
		stocks[stck.Ticker] = stck
	}
	return stocks, nil
}
