package main

import endpoint "github.com/CzarSimon/go-endpoint"

// DefaultPeriod Is the default time period for retrieving news
const (
	DefaultPeriod = "TODAY"
	SERVER_NAME   = "NEWS_SERVER"
)

//Config contains internal config variables for the server
type Config struct {
	server endpoint.ServerAddr
	db     endpoint.SQLConfig
}

func getConfig() Config {
	return Config{
		server: endpoint.NewServerAddr(SERVER_NAME),
		db:     endpoint.NewPGConfig("PG"),
	}
}

// periodMonthMap Maps a Time period to the months to subtract
type periodMonthMap map[string]int

// newPeriodMonthMap Generates a new periodMonthMap
func newPeriodMonthMap() periodMonthMap {
	pmm := make(periodMonthMap)
	pmm[DefaultPeriod] = 0
	pmm["0M"] = 0
	pmm["1M"] = -1
	pmm["3M"] = -3
	return pmm
}
