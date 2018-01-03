package action

import (
	"fmt"

	"github.com/urfave/cli"
)

// Placeholder Action to user for command development
func Placeholder(c *cli.Context) error {
	fmt.Printf("%+v\n", c.Args)
	return nil
}
