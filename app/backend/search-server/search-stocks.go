package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/query"
)

//SearchStocks Searches for stocks using a given query
func (env *Env) SearchStocks(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.MethodNotAllowed
	}
	q, err := query.ParseValue(r, "query")
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	stocks, err := queryForStock(q, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendJSON(w, stocks)
	return nil
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
