package api

import (
	"fmt"
	"os"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
)

// GetStocks gets tracked stocks from the admin api.
func GetStocks() []stock.Stock {
	stocks := make([]stock.Stock, 0)
	getJSON("stock", &stocks)
	return stocks
}

// GetUntrackedTickers gets top untracked tickers from the admin api.
func GetUntrackedTickers() []string {
	tickers := make([]string, 0)
	getJSON("untracked-tickers", &tickers)
	return tickers
}

func GetSpamCandidates() []spam.Candidate {
	candidates := make([]spam.Candidate, 0)
	getJSON("spam", &candidates)
	return candidates
}

func getJSON(resource string, v interface{}) {
	resp, err := makeGetRequest(createRoute(resource))
	checkErr(err)
	err = util.DecodeJSON(resp.Body, v)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
