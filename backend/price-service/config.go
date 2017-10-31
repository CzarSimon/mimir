package main

import (
	"log"

	"github.com/CzarSimon/util"
)

// Config Holds configuration values
type Config struct {
	PriceDB  util.PGConfig
	TickerDB util.PGConfig
	Timing   string
	Timezone string
}

// GetConfig Returns a new Config struct based on evnirionment variables
func GetConfig() Config {
	pdbHost := util.GetEnvVar("PG_HOST", "localhost")
	tdbHost := util.GetEnvVar("TICKER_DB_HOST", "localhost")
	dbPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	tickerDBConfig := util.GetPGConfig(tdbHost, dbPwd, "simon", "mimirprod")
	return Config{
		PriceDB:  util.GetPGConfig(pdbHost, dbPwd, "simon", "mimirprod"),
		TickerDB: tickerDBConfig,
		Timing:   util.GetEnvVar("RETRIVAL_TIME", "02:05"),
		Timezone: util.GetEnvVar("TIMEZONE", "America/New_York"),
	}
}

// LogTiming Logs the timing configuration
func (config Config) LogTiming() {
	log.Printf("Trigger time: %s Exchange timezone: %s", config.Timing, config.Timezone)
}
