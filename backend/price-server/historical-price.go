package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/CzarSimon/util"
)

// StockPrice Contains the closing trade price for a given date
type StockPrice struct {
	Price float64   `json:"price"`
	Date  time.Time `json:"date"`
}

// PriceRequest Conatins ticker and date information for a price query
type PriceRequest struct {
	Ticker   string    `json:"ticker"`
	FromDate time.Time `json:"fromDate"`
}

// GetHistoricalPrices Request handler for retrival of a tickers historical
// price after a given start date
func (env *Env) GetHistoricalPrices(res http.ResponseWriter, req *http.Request) {
	priceRequest, err := parsePriceRequest(req)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	prices, err := GetPricesFromDB(priceRequest, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonBody, err := json.Marshal(prices)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// GetPricesFromDB Retrives historical prices from database
func GetPricesFromDB(priceRequest PriceRequest, conn *sql.DB) ([]StockPrice, error) {
	prices := make([]StockPrice, 0)
	rows, err := conn.Query(
		getPriceSelectQuery(), priceRequest.Ticker, priceRequest.FromDate)
	defer rows.Close()
	if err != nil {
		return prices, err
	}
	var price StockPrice
	for rows.Next() {
		err = rows.Scan(&price.Date, &price.Price)
		if err != nil {
			return prices, err
		}
		prices = append(prices, price)
	}
	return prices, nil
}

// getPriceSelectQuery Returns the query to select historical prices of a given stock
func getPriceSelectQuery() string {
	return `SELECT PRICE_DATE, PRICE FROM HISTORICAL_PRICE
          WHERE TICKER=$1 AND PRICE_DATE>=$2
          ORDER BY PRICE_DATE DESC`
}

// parsePriceRequest Parses a price request from a http GET request
func parsePriceRequest(req *http.Request) (PriceRequest, error) {
	priceRequest := PriceRequest{}
	if req.Method != http.MethodGet {
		return priceRequest, errors.New("Method not allowed. Only GET")
	}
	priceRequest.Ticker = parseTicker(req)
	if priceRequest.Ticker == "" {
		return priceRequest, errors.New("No ticker in query")
	}
	var err error
	priceRequest.FromDate, err = parseDate(req, "fromDate")
	if err != nil {
		return priceRequest, errors.New("Could not parse fromDate")
	}
	return priceRequest, nil
}

// parseTicker Parses and uppercases ticker from URL query
func parseTicker(req *http.Request) string {
	ticker := req.URL.Query().Get("ticker")
	return strings.ToUpper(ticker)
}

// parseFromDate Parses a date from URL query
func parseDate(req *http.Request, key string) (time.Time, error) {
	Layout := "2006-01-02"
	date, err := time.Parse(Layout, req.URL.Query().Get(key))
	if err != nil {
		return date, err
	}
	return date, nil
}
