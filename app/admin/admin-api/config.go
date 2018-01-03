package main

import (
	endpoint "github.com/CzarSimon/go-endpoint"
)

const SERVER_NAME = "ADMIN_API"

// Config is the main configuration type
type Config struct {
	server endpoint.ServerAddr
}

// NewConfig Sets up a new configuration struct
func NewConfig() Config {
	return Config{
		server: endpoint.NewServerAddr(SERVER_NAME),
	}
}
