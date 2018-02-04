package main

import endpoint "github.com/CzarSimon/go-endpoint"

const (
	SERVER_NAME = "NEWS_CLUSTERER"
)

// Config Holds service configurations
type Config struct {
	server endpoint.ServerAddr
	db     endpoint.SQLConfig
}

func getConfig() Config {
	return Config{
		server: endpoint.NewServerAddr(SERVER_NAME),
		db:     endpoint.NewPGConfig("DB"),
	}
}
