package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
)

// stockHandler handles request related to the stock resource.
func (env *Env) stockHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return env.getStocks(w, r)
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
	var ns stock.NullStock
	for rows.Next() {
		err := rows.Scan(&ns.Ticker, &ns.Name, &ns.Description, &ns.ImageURL, &ns.Website)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, ns.Stock())
	}
	return stocks, nil
}

// getStockQuery creates query for tracked stocks.
func getStockQuery() string {
	return `
		SELECT TICKER, NAME, DESCRIPTION, IMAGE_URL, WEBSITE
		FROM STOCK`
}

// storeNewStock stores a new stock in the connected databases.
func (env *Env) storeNewStock(w http.ResponseWriter, r *http.Request) error {
	var newStock stock.Stock
	err := util.DecodeJSON(r.Body, &newStock)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	err = storeStockInAppDB(newStock, env.AppDB)
	if err != nil {
		return err
	}
	err = storeStockInTweetDB(newStock, env.TweetDB)
	if err != nil {
		return err
	}
	httputil.SendOK(w)
	return nil
}

// storeStockInAppDB attempts to store a new stock in the app database.
func storeStockInAppDB(newStock stock.Stock, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO STOCK(TICKER, NAME, DESCRIPTION, IMAGE_URL, WEBSITE) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		newStock.Ticker, newStock.Name, newStock.Description, newStock.ImageURL, newStock.Website)
	return err
}

// storeStockInTweetDB attempts to store a new stock in the tweet.
func storeStockInTweetDB(newStock stock.Stock, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO STOCKS(TICKER, NAME, STOREDAT) VALUES($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newStock.Ticker, newStock.Name, getFirstStoredDate())
	return err
}

// getFirstStoredDate returns the datestamp when ticker tracking
// should begin from the point of view of the tweet db.
func getFirstStoredDate() time.Time {
	return time.Now().UTC().AddDate(0, 0, 1)
}
