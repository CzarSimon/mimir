package main

import "github.com/CzarSimon/util"

// Config Holds service configurations
type Config struct {
	server util.ServerConfig
	db     util.PGConfig
}

func getConfig() Config {
	return Config{
		server: getServerConfig(),
		db:     util.GetPGConfig("localhost", "pwd", "simon", "mimirprod"),
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: "6000",
	}
}
