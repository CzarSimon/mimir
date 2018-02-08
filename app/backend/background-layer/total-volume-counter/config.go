package main

import endpoint "github.com/CzarSimon/go-endpoint"

// Config holds configuration values.
type Config struct {
	DB endpoint.SQLConfig
}

// GetConfig sets up a new config struct based on environemnt variables.
func GetConfig() Config {
	return Config{
		DB: endpoint.NewPGConfig("DB"),
	}
}
