package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
)

// stockHandler handles request related to the stock resource.
func (env *Env) stockHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		httputil.SendOK(w)
		return nil
	case http.MethodPost:
		return env.storeNewStock(w, r)
	default:
		return httputil.MethodNotAllowed
	}
}

// getStocks querys for and sends tracked stocks to the requestor.
func (env *Env) getStocks(w http.ResponseWriter, r *http.Request) error {
	stocks, err := queryForStocks(r, env.AppDB)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, stocks)
}

// queryForStocks querys the database for tracked stocks
// based on the request query
func queryForStocks(r *http.Request, db *sql.DB) ([]stock.Stock, error) {
	rows, err := db.Query(getStockQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return constructStocksList(rows)
}

// constructStocksList creates a list of stocks from a result set
func constructStocksList(rows *sql.Rows) ([]stock.Stock, error) {
	stocks := make([]stock.Stock, 0)
	var s stock.Stock
	for rows.Next() {
		err := rows.Scan(&s.Ticker, &s.Ticker, &s.Description, &s.ImageURL, &s.Website)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}
	return stocks, nil
}

// getStockQuery creates query for tracked stocks.
func getStockQuery() string {
	return `
		SELECT TICKER, NAME, DESCRIPTION, IMAGE_URL, WEBSITE
		FROM STOCK`
}

// storeNewStock stores a new stock.
func (env *Env) storeNewStock(w http.ResponseWriter, r *http.Request) error {
	var newStock stock.Stock
	err := util.DecodeJSON(r.Body, &newStock)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	log.Println(newStock)
	httputil.SendOK(w)
	return nil
}
