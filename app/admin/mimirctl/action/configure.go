package action

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
	"github.com/urfave/cli"
)

// Configure Gets and stores api configuration information from the user
func Configure(c *cli.Context) error {
	config := getApiConfig()
	fmt.Println(config)
	testApiConfig(config)
	config.Save()
	return nil
}

// getApiConfig Prompts the user to input api configuration
func getApiConfig() api.Config {
	var config api.Config
	fmt.Println("Enter api configuration")
	config.API.Host = getInput("Host")
	config.API.Port = getInput("Port")
	config.API.Protocol = getInputWithDefault("Protocol", "https")
	config.Auth.AccessKey = getInput("Access Key")
	return config
}
