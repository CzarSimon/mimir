package action

import (
	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
	"github.com/urfave/cli"
)

// Ping Tests that the cli can authenticate with the admin-api
func Ping(c *cli.Context) error {
	config, err := api.GetConfig()
	checkErr(err)
	testApiConfig(config)
	return nil
}

// Login Sets and verifies login credentials
func Login(c *cli.Context) error {
	config, err := api.GetConfig()
	checkErr(err)
	config.Auth.AccessKey = getHiddenInput("Access Key")
	testApiConfig(config)
	config.Save()
	return nil
}

// Logout Closes the admin api and removes the clients login credentials
func Logout(c *cli.Context) error {
	config, err := api.GetConfig()
	checkErr(err)
	config.Auth.AccessKey = ""
	config.Save()
	return nil
}
