package main

import (
	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	DATABASE_NAME = "DB"
	SERVER_NAME   = "APP_SERVER"
)

// Config Struct to hold configurations
type Config struct {
	DB     endpoint.SQLConfig
	Server endpoint.ServerAddr
}

// getConfig Gets main configurations
func getConfig() Config {
	return Config{
		DB:     endpoint.NewPGConfig(DATABASE_NAME),
		Server: endpoint.NewServerAddr(SERVER_NAME),
	}
}
