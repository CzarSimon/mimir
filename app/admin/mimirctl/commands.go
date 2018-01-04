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
		Action: action.Configure,
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

// PwdGenCommand Generates a password for use in the admin-api
func PwdGenCommand() cli.Command {
	return cli.Command{
		Name:   "pwd-gen",
		Usage:  "Generates a password for use in the admin-api",
		Action: action.GeneratePassword,
	}
}

// PingCommand Tests that the cli can authenticate with the admin-api
func PingCommand() cli.Command {
	return cli.Command{
		Name:   "ping",
		Usage:  "Tests that the cli can authenticate with the admin-api",
		Action: action.Placeholder,
	}
}
