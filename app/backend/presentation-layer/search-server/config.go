package main

import (
	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	SERVER_NAME = "SEARCH_SERVER"
)

// Config is the main configuration type.
type Config struct {
	server endpoint.ServerAddr
	db     endpoint.SQLConfig
}

// getConfig sets up service configuration.
func getConfig() Config {
	return Config{
		server: endpoint.NewServerAddr(SERVER_NAME),
		db:     endpoint.NewPGConfig("DB"),
	}
}
