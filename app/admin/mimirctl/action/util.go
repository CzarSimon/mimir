package action

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// Placeholder Action to user for command development
func Placeholder(c *cli.Context) error {
	fmt.Printf("Command: %+v\n", c.Args())
	return nil
}

// testApiConfig Checks if a supplied config is correct
func testApiConfig(config api.Config) {
	checkErr(api.Ping(config))
	fmt.Println("Success! mimirctl authenticated")
}

// getInput Gets user input from stdin
func getInput(key string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(key + ": ")
	text, err := reader.ReadString('\n')
	checkErr(err)
	return strings.Replace(text, "\n", "", -1)
}

// getInputWithDefault Gets user input, if empty returns a supplied default value
func getInputWithDefault(key, defaultVal string) string {
	value := getInput(fmt.Sprintf("%s (%s)", key, defaultVal))
	if value == "" {
		return defaultVal
	}
	return value
}

// getHiddenInput gets hidden user input from stdin.
func getHiddenInput(key string) string {
	fmt.Printf("%s: ", key)
	value, err := terminal.ReadPassword(0)
	fmt.Println()
	checkErr(err)
	return string(value)
}

// checkErr prints error and halts execution if error is not nil.
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

// printHeader prints a highighted message to stdout.
func printHeader(msg string) {
	fmt.Printf("------- %s -------\n", msg)
}
