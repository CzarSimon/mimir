package action

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
	"github.com/urfave/cli"
)

// GetMap maps a resource type to a get function.
var GetMap = ResourceMap{
	STOCK:             getStocks,
	STOCKS:            getStocks,
	UNTRACKED_TICKERS: getUntrackedTickers,
}

// Gets a specified resource type from the admin-api
func Get(c *cli.Context) error {
	function := GetMap.GetFunc(getResource(c))
	return function(c)
}

// getStocks gets a list of stocks from the api.
func getStocks(c *cli.Context) error {
	stocks := api.GetStocks()
	printHeader("Stocks")
	for _, s := range stocks {
		fmt.Printf("%s\n\n", s.String())
	}
	return nil
}

// getUntrackedTickers gets a ranked list of untracked tickers from the api.
func getUntrackedTickers(c *cli.Context) error {
	tickers := api.GetUntrackedTickers()
	printHeader("Untracked Tickers")
	for i, ticker := range tickers {
		fmt.Printf("%d. - %s\n", i+1, ticker)
	}
	return nil
}
