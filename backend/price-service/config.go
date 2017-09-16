package main

import (
	"log"

	"github.com/CzarSimon/util"
)

// Config Holds configuration values
type Config struct {
	PriceDB  util.PGConfig
	TickerDB RDBConfig
	Timing   string
	Timezone string
}

// GetConfig Returns a new Config struct based on evnirionment variables
func GetConfig() Config {
	pdbHost := util.GetEnvVar("PG_HOST", "localhost")
	pdbPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	return Config{
		PriceDB:  util.GetPGConfig(pdbHost, pdbPwd, "simon", "mimirprod"),
		TickerDB: getRDBConfig(),
		Timing:   util.GetEnvVar("RETRIVAL_TIME", "02:05"),
		Timezone: util.GetEnvVar("TIMEZONE", "America/New_York"),
	}
}

// LogTiming Logs the timing configuration
func (config Config) LogTiming() {
	log.Printf("Trigger time: %s Exchange timezone: %s", config.Timing, config.Timezone)
}

// RDBConfig Holds configuration values for connecting to a rethinkdb database
type RDBConfig struct {
	Host, Port, DB string
}

// getRDBConfig Returns a new RDBConfig based on evnirionment variables
func getRDBConfig() RDBConfig {
	return RDBConfig{
		Host: util.GetEnvVar("TICKER_DB_HOST", "localhost"),
		Port: "28015",
		DB:   util.GetEnvVar("TICKER_DB_NAME", "mimir_app_server"),
	}
}
