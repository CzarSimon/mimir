package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
)

// PriceRequest Conatins ticker and date information for a price query
type PriceRequest struct {
	Ticker   string    `json:"ticker"`
	FromDate time.Time `json:"fromDate"`
}

// GetHistoricalPrices Request handler for retrival of a tickers historical
// price after a given start date
func (env *Env) GetHistoricalPrices(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.MethodNotAllowed
	}
	priceRequest, err := parsePriceRequest(r)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	prices, err := GetPricesFromDB(priceRequest, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendJSON(w, prices)
	return nil
}

// GetPricesFromDB Retrives historical prices from database
func GetPricesFromDB(priceRequest PriceRequest, conn *sql.DB) ([]stock.Price, error) {
	prices := make([]stock.Price, 0)
	rows, err := conn.Query(
		getPriceSelectQuery(), priceRequest.Ticker, priceRequest.FromDate)
	defer rows.Close()
	if err != nil {
		return prices, err
	}
	var price stock.Price
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
	return `SELECT PRICE_DATE, PRICE FROM CLOSE_PRICE
          WHERE TICKER=$1 AND PRICE_DATE>=$2
          ORDER BY PRICE_DATE DESC`
}

// parsePriceRequest Parses a price request from a http GET request
func parsePriceRequest(req *http.Request) (PriceRequest, error) {
	var priceRequest PriceRequest
	var err error
	priceRequest.Ticker, err = parseTicker(req)
	if err != nil {
		return priceRequest, err
	}
	priceRequest.FromDate, err = parseDate(req)
	if err != nil {
		return priceRequest, errors.New("Could not parse fromDate")
	}
	return priceRequest, nil
}

// parseTicker Parses and uppercases ticker from URL query
func parseTicker(req *http.Request) (string, error) {
	ticker, err := util.ParseValueFromQuery(req, "ticker", "Could not parse ticker")
	if err != nil {
		return ticker, err
	}
	return strings.ToUpper(ticker), nil
}

// parseFromDate Parses a date from URL query
func parseDate(req *http.Request) (time.Time, error) {
	dateStr, err := util.ParseValueFromQuery(req, "fromDate", "Could not parse from date")
	if err != nil {
		return time.Now(), err
	}
	Layout := "2006-01-02"
	date, err := time.Parse(Layout, dateStr)
	if err != nil {
		return date, err
	}
	return date, nil
}
