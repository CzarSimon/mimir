package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CzarSimon/util"
)

//SearchStocks Searches for stocks using a given query
func (env *Env) SearchStocks(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrRes(res, errors.New("Method not allowed"))
		return
	}
	query, err := parseQuery(req)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	stocks, err := queryForStock(query, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonBody, err := json.Marshal(stocks)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// queryForStock Query database for stocks matching the supplied query
func queryForStock(query string, db *sql.DB) ([]Stock, error) {
	stocks := make([]Stock, 0)
	sqlQuery := `SELECT TICKER, NAME FROM STOCKS
							 WHERE TICKER LIKE '%' || $1 || '%'
							 OR LOWER(NAME) LIKE '%' || LOWER($1) || '%'
							 ORDER BY TOTAL_COUNT DESC LIMIT 10`
	rows, err := db.Query(sqlQuery, query)
	defer rows.Close()
	if err != nil {
		return stocks, err
	}
	var stock Stock
	for rows.Next() {
		err := rows.Scan(&stock.Ticker, &stock.Name)
		if err != nil {
			return stocks, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

// parseQuery Parses the requested search query
func parseQuery(req *http.Request) (string, error) {
	query := req.URL.Query().Get("query")
	if query == "" {
		return query, errors.New("No query supplied")
	}
	return query, nil
}
