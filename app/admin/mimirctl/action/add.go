package action

import (
	"fmt"

	"github.com/urfave/cli"
)

// AddMap maps resource types to a add function.
var AddMap = ResourceMap{
	STOCK: addStock,
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
