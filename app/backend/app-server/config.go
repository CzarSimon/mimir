package main

import (
	"github.com/CzarSimon/util"
)

// Config Struct to hold configurations
type Config struct {
	DB     util.PGConfig
	Server util.ServerConfig
}

// getConfig Gets main configurations
func getConfig() Config {
	return Config{
		DB: util.GetPGConfig("localhost", "pwd", "simon", "mimirprod"),
		Server: util.ServerConfig{
			Port: "3000",
		},
	}
}
