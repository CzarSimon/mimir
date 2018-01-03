package action

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
	"github.com/CzarSimon/util"
	"github.com/urfave/cli"
)

// Placeholder Action to user for command development
func Placeholder(c *cli.Context) error {
	fmt.Printf("%+v\n", c.Args)
	return nil
}

// testApiConfig Checks if a supplied config is correct
func testApiConfig(config api.Config) {
	if err := api.Ping(config); err != nil {
		util.CheckErrFatal(err)
	}
	fmt.Println("Success! mimirctl authenticated")
}

// getInput Gets user input from stdin
func getInput(key string) string {
	var value string
	fmt.Print(key + ": ")
	fmt.Scanf("%s", &value)
	return value
}

// getInputWithDefault Gets user input, if empty returns a supplied default value
func getInputWithDefault(key, defaultVal string) string {
	value := getInput(fmt.Sprintf("%s (%s)", key, defaultVal))
	if value == "" {
		return defaultVal
	}
	return value
}
