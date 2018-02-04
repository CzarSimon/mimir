package main

import (
	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
)

// Config Holds configuration values
type Config struct {
	PriceDB  endpoint.SQLConfig
	TickerDB endpoint.SQLConfig
	Timezone string
}

// GetConfig Returns a new Config struct based on evnirionment variables
func GetConfig() Config {
	return Config{
		PriceDB:  endpoint.NewPGConfig("PRICE_DB"),
		TickerDB: endpoint.NewPGConfig("TICKER_DB"),
		Timezone: util.GetEnvVar("TIMEZONE", "America/New_York"),
	}
}
