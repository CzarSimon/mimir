package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CzarSimon/util"
)

// Stock Holds information about a stock
type Stock struct {
	Ticker      string `json:"ticker"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
	stock, err := RetriveStockInfo(ticker, env.db)
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
func RetriveStockInfo(ticker string, db *sql.DB) (Stock, error) {
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
