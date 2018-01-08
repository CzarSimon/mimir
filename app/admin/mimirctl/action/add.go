package action

import (
	"fmt"

	"github.com/urfave/cli"
)

// Add constructs a new record of a specified resource
// and instructs the admin-api to store it.
func Add(c *cli.Context) error {
	fmt.Println("Running add command")
	return nil
}
