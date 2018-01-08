package action

import (
	"fmt"

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
	fmt.Println("getting stocks")
	return nil
}

// getUntrackedTickers gets a ranked list of untracked tickers from the api.
func getUntrackedTickers(c *cli.Context) error {
	fmt.Println("getting untracked tickers")
	return nil
}
