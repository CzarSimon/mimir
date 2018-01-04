package action

import (
	"fmt"
	"log"
	
	"golang.org/x/crypto/ssh/terminal"
	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
	"github.com/urfave/cli"
)

// Placeholder Action to user for command development
func Placeholder(c *cli.Context) error {
	fmt.Printf("Command: mimirctl %s", c.Args(1))
	return nil
}

// testApiConfig Checks if a supplied config is correct
func testApiConfig(config api.Config) {
	checkErr(api.Ping(config))
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

// getHiddenInput Gets hidden user input from stdin
func getHiddenInput(key string) string {
	fmt.Printf("%s: \n", key)
	value, err := terminal.ReadPassword(0)
	checkErr(err)
	return string(value)
}

// checkErr Prints error and halts execution if error is not nil
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal()
	}
}
