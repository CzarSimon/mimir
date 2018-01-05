package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

// SugestionRequest holds the tickers to exclude in a search recomendation
type SugestionRequest struct {
	Tickers []string `json:"tickers"`
}

// Stock holds stock informaton
type Stock struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

// GetSearchSugestions retrives search sugestions of a user excluding a given set of tickers
func (env *Env) GetSearchSugestions(res http.ResponseWriter, req *http.Request) {
	tickers := getTickers(req)
	sugestions, err := getSugestions(tickers, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	js, err := json.Marshal(sugestions)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, js)
}

//ParseTickers turns a search request into a query
func (query SugestionRequest) ParseTickers() Tickers {
	tickers := make(Tickers, 0)
	for _, ticker := range query.Tickers {
		tickers = append(tickers, strings.ToUpper(ticker))
	}
	return tickers
}

//Tickers is a slice of tickers
type Tickers []string

//ToUpper turns every ticker in a slice to uppercase
func (tickers Tickers) ToUpper() Tickers {
	upperTickers := make(Tickers, 0)
	for _, ticker := range tickers {
		upperTickers = append(upperTickers, strings.ToUpper(ticker))
	}
	return upperTickers
}

// getTickers Parses the tickers to exclude from the sugestion
func getTickers(req *http.Request) Tickers {
	tickers, err := parseTickersFromBody(req)
	if err == nil {
		return tickers
	}
	tickers = req.URL.Query()["ticker"]
	return tickers.ToUpper()
}

// parseTickersFromBody Attempts to parse the request body for tickers
func parseTickersFromBody(req *http.Request) (Tickers, error) {
	var tickers Tickers
	var query SugestionRequest
	err := util.DecodeJSON(req.Body, &query)
	if err != nil {
		return tickers, err
	}
	tickers = query.ParseTickers()
	return tickers, nil
}

// getSugestions Query the database for search sugestions
func getSugestions(tickers Tickers, db *sql.DB) ([]Stock, error) {
	stocks := make([]Stock, 0)
	rows, err := queryForSugestions(tickers, db)
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
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

// queryForSugestions Construct and execute a query for search sugestions
func queryForSugestions(tickers Tickers, db *sql.DB) (*sql.Rows, error) {
	if len(tickers) > 0 {
		query := "SELECT ticker, name FROM stocks WHERE is_tracked=TRUE AND NOT ticker = ANY($1) ORDER BY total_count DESC LIMIT 5"
		return db.Query(query, pq.Array(tickers))
	}
	return db.Query("SELECT ticker, name FROM stocks WHERE is_tracked=TRUE ORDER BY total_count DESC LIMIT 5")
}
