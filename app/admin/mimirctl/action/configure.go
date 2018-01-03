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
	config.API.Host = getInput("\nHost")
	config.API.Port = getInput("\nPort")
	config.API.Protocol = getInputWithDefault("\nProtocol", "https")
	config.Auth.AccessKey = getInput("\nAccess Key")
	return config
}
