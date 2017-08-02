package main

import "github.com/CzarSimon/util"

// DefaultPeriod Is the default time period for retrieving news
const DefaultPeriod string = "TODAY"

//Config contains internal config variables for the server
type Config struct {
	server util.ServerConfig
	db     util.PGConfig
}

func getConfig() Config {
	return Config{
		server: getServerConfig(),
		db:     util.GetPGConfig("localhost", "pwd", "simon", "mimirprod"),
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: "5050",
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
