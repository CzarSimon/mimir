package main

import (
	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	SERVER_NAME = "PRICE_SERVER"
)

// Config Holds configuration values
type Config struct {
	DB     endpoint.SQLConfig
	Server endpoint.ServerAddr
}

// GetConfig Returns a new Config struct based on evnirionment variables
func GetConfig() Config {
	return Config{
		DB:     endpoint.NewPGConfig("DB"),
		Server: endpoint.NewServerAddr(SERVER_NAME),
	}
}
