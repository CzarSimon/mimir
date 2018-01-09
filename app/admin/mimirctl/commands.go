package main

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/action"
	"github.com/urfave/cli"
)

// LoginCommand sets and verifies login credentials.
func LoginCommand() cli.Command {
	return cli.Command{
		Name:   "login",
		Usage:  "Sets and verifies login credentials",
		Action: action.Login,
	}
}

// ConfigureCommand configures the cli app for use.
func ConfigureCommand() cli.Command {
	return cli.Command{
		Name:   "configure",
		Usage:  fmt.Sprintf("Configures %s for use", APP_NAME),
		Action: action.Configure,
	}
}

// LogoutCommand closes the admin api and removes the clients login credentials.
func LogoutCommand() cli.Command {
	return cli.Command{
		Name:   "logout",
		Usage:  "Closes the admin api and removes the clients login credentials",
		Action: action.Logout,
	}
}

// PwdGenCommand generates a password for use in the admin-api.
func PwdGenCommand() cli.Command {
	return cli.Command{
		Name:   "pwd-gen",
		Usage:  "Generates a password for use in the admin-api",
		Action: action.GeneratePassword,
	}
}

// PingCommand tests that the cli can authenticate with the admin-api.
func PingCommand() cli.Command {
	return cli.Command{
		Name:   "ping",
		Usage:  "Tests that the cli can authenticate with the admin-api",
		Action: action.Ping,
	}
}

// GetCommand querys the admin api for a list of the specified resource.
func GetCommand() cli.Command {
	return cli.Command{
		Name:   "get",
		Usage:  "Querys the admin api for a list of the specified resource",
		Action: action.Get,
	}
}

// AddCommand adds a specified resource.
func AddCommand() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Adds a specified resource",
		Action: action.Add,
	}
}

// LabelCommand querys the admin api for labeling candidtes
// and sends back labeled data.
func LabelCommand() cli.Command {
	return cli.Command{
		Name:   "label",
		Usage:  "Querys the admin api for labeling candidtes and sends back labeled data",
		Action: action.Label,
	}
}
