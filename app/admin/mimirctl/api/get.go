package api

import (
	"fmt"
	"os"

	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
)

// GetStocks gets tracked stocks from the admin api.
func GetStocks() []stock.Stock {
	resp, err := makeGetRequest(createRoute("stock"))
	checkErr(err)
	stocks := make([]stock.Stock, 0)
	err = util.DecodeJSON(resp.Body, &stocks)
	checkErr(err)
	return stocks
}

// GetUntrackedTickers gets top untracked tickers from the admin api.
func GetUntrackedTickers() []string {
	resp, err := makeGetRequest(createRoute("untracked-tickers"))
	checkErr(err)
	tickers := make([]string, 0)
	err = util.DecodeJSON(resp.Body, &tickers)
	checkErr(err)
	return tickers
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
