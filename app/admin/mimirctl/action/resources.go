package action

import (
	"fmt"

	"github.com/urfave/cli"
)

// Resource types
const (
	STOCK             = "stock"
	STOCKS            = "stocks"
	UNTRACKED_TICKERS = "untracked-tickers"
	SPAM              = "spam"
)

// ResourceType names of the types of resources managed by the admin tools
type ResourceType string

// ResourceMap maps a resource to an action function
type ResourceMap map[string]cli.ActionFunc

// GetFunc gets an action function based on a resource type. Exits if unsuccessful
func (rm ResourceMap) GetFunc(resource string) cli.ActionFunc {
	function, ok := rm[resource]
	if !ok {
		checkErr(fmt.Errorf("No action exists for resource \"%s\"", resource))
	}
	return function
}

// getResource gets a resource type from supplied command line arguments
func getResource(c *cli.Context) string {
	resourceArgIndex := 2
	return c.Args().Get(resourceArgIndex)
}
