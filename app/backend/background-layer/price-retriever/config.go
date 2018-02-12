package main

import (
	"strconv"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/go-file-heartbeat/heartbeat"
	"github.com/CzarSimon/util"
)

const (
	DefaultExchangeOpenHour  = "9"
	DefaultExchangeCloseHour = "16"
)

// Config holds configuration values.
type Config struct {
	PriceDB           endpoint.SQLConfig
	TickerDB          endpoint.SQLConfig
	Timezone          string
	ExchangeOpenHour  int
	ExchangeCloseHour int
	Heartbeat         heartbeat.Config
}

// GetConfig Returns a new Config struct based on evnirionment variables
func GetConfig() Config {
	heartbeatConfig, err := heartbeat.NewConfigFromEnv()
	util.CheckErrFatal(err)
	open, close := getExchangeHours()
	return Config{
		PriceDB:           endpoint.NewPGConfig("PRICE_DB"),
		TickerDB:          endpoint.NewPGConfig("TICKER_DB"),
		Timezone:          util.GetEnvVar("TIMEZONE", "America/New_York"),
		ExchangeOpenHour:  open,
		ExchangeCloseHour: close,
		Heartbeat:         heartbeatConfig,
	}
}

// getExchangeHours gets set exchange opening hours from environment.
func getExchangeHours() (int, int) {
	open := util.GetEnvVar("EXCHANGE_OPEN_HOUR", DefaultExchangeOpenHour)
	close := util.GetEnvVar("EXCHANGE_CLOSE_HOUR", DefaultExchangeCloseHour)
	openHour, err := strconv.Atoi(open)
	util.CheckErrFatal(err)
	closeHour, err := strconv.Atoi(close)
	util.CheckErrFatal(err)
	return openHour, closeHour
}
