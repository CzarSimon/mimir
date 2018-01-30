package main

import (
	"log"

	"github.com/CzarSimon/go-file-heartbeat/heartbeat"
	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
)

// Config Holds configuration values
type Config struct {
	PriceDB   endpoint.SQLConfig
	TickerDB  endpoint.SQLConfig
	Timing    TimingConfig
	Timezone  string
	Heartbeat heartbeat.Config
}

// TimingConfig Holds timing config for scheduling
type TimingConfig struct {
	ClosePriceTime  string
	LatestPriceTime string
}

// GetConfig Returns a new Config struct based on evnirionment variables
func GetConfig() Config {
	heartbeatConfig, err := heartbeat.NewConfigFromEnv()
	util.CheckErrFatal(err)
	return Config{
		PriceDB:   endpoint.NewPGConfig("PRICE_DB"),
		TickerDB:  endpoint.NewPGConfig("TICKER_DB"),,
		Timing:    getTimingConfig(),
		Timezone:  util.GetEnvVar("TIMEZONE", "America/New_York"),
		Heartbeat: heartbeatConfig,
	}
}

// getTimingConfig Returns the timing config for scheduled tasks
func getTimingConfig() TimingConfig {
	return TimingConfig{
		ClosePriceTime: util.GetEnvVar("RETRIVAL_TIME", "02:05"),
	}
}

// LogTiming Logs the timing configuration
func (config Config) LogTiming() {
	log.Printf("Trigger time: %s Exchange timezone: %s", config.Timing.ClosePriceTime, config.Timezone)
}
