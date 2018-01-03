package main

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/action"
	"github.com/urfave/cli"
)

// LoginCommand Sets and verifies login credentials
func LoginCommand() cli.Command {
	return cli.Command{
		Name:   "login",
		Usage:  "Sets and verifies login credentials",
		Action: action.Placeholder,
	}
}

// ConfigureCommand Configures the cli app for use
func ConfigureCommand() cli.Command {
	return cli.Command{
		Name:   "configure",
		Usage:  fmt.Sprintf("Configures %s for use", APP_NAME),
		Action: action.Placeholder,
	}
}

// LogoutCommand Closes the admin api and removes the clients login credentials
func LogoutCommand() cli.Command {
	return cli.Command{
		Name:   "logout",
		Usage:  "Closes the admin api and removes the clients login credentials",
		Action: action.Placeholder,
	}
}
