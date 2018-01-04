package main

import (
	"os"

	"github.com/urfave/cli"
)

const (
	APP_NAME    = "mimirctl"
	APP_USAGE   = "Administration tool for mimir"
	APP_VERSION = "0.1"
)

// getApp Sets up a cli app and returns it
func getApp() *cli.App {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = APP_USAGE
	app.Version = APP_VERSION
	app.Commands = getAppCommands()
	return app
}

// getAppCommands Returns a list of cli subcommands
func getAppCommands() []cli.Command {
	return []cli.Command{
		LoginCommand(),
		ConfigureCommand(),
		LogoutCommand(),
		PwdGenCommand(),
		PingCommand(),
	}
}

func main() {
	app := getApp()
	app.Run(os.Args)
}
