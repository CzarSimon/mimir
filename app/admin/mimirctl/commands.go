package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// LoginCommand Sets and verifies login credentials
func LoginCommand() cli.Command {
	return cli.Command{
		Name:   "login",
		Usage:  "Sets and verifies login credentials",
		Action: placeholderAction,
	}
}

// ConfigureCommand Configures the cli app for use
func ConfigureCommand() cli.Command {
	return cli.Command{
		Name:   "configure",
		Usage:  fmt.Sprintf("Configures %s for use", APP_NAME),
		Action: placeholderAction,
	}
}

// LogoutCommand Closes the admin api and removes the clients login credentials
func LogoutCommand() cli.Command {
	return cli.Command{
		Name:   "logout",
		Usage:  "Closes the admin api and removes the clients login credentials",
		Action: placeholderAction,
	}
}

func placeholderAction(c *cli.Context) error {
	fmt.Println(c.Args)
	return nil
}
