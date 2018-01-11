package action

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/urfave/cli"
)

// AddMap maps resource types to a add function.
var AddMap = ResourceMap{
	STOCK: addStock,
	SPAM:  addSpam,
}

// Add constructs a new record of a specified resource
// and instructs the admin-api to store it.
func Add(c *cli.Context) error {
	function := AddMap.GetFunc(getResource(c))
	return function(c)
}

// addStock add command for the stock resource.
func addStock(c *cli.Context) error {
	fmt.Println("Adding stock")
	return nil
}

// getNewStock gets a new stock by interacting with user and apis.
func getNewStock(ticker string) stock.Stock {
	newStock := api.GetNewStock(ticker)
	fmt.Printf("Ticker: %s\n", ticker)
	newStock.Name = getInput("Name")
	newStock.Description = getInput("Description")
	newStock.ImageURL = getInput("Image Url")
	newStock.Website = getInput("Website")
	return newStock
}

// addSpam interactive command for labeling a custom spam candidate.
func addSpam(c *cli.Context) error {
	text := getInput("Text")
	labelSpam(spam.NewCandidate(text))
	return nil
}
