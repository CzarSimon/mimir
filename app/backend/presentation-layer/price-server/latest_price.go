package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/lib/pq"
)

// GetLatestPrices Gets the latest prices for a supplied set of tickers
func (env *Env) GetLatestPrices(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.MethodNotAllowed
	}
	tickers, err := parseTickersFromQuery(r)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	prices, err := getLatestPricesFromDB(tickers, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendJSON(w, prices)
	return nil
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
