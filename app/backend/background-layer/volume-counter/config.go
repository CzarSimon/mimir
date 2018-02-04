package main

import (
	"log"
	"os"
	"time"

	"github.com/CzarSimon/util"
)

//Config holds configurations and environment variables
type Config struct {
	DB         util.PGConfig
	Timing     TimingConfig
	Server     util.ServerConfig
	Routes     RoutesConfig
	StableDate time.Time
}

func getConfig() Config {
	pgHost := util.GetEnvVar("PG_HOST", "localhost")
	pgPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	return Config{
		DB:         util.GetPGConfig(pgHost, pgPwd, "simon", "mimirprod"),
		Timing:     getTimingConfig(),
		Server:     getServerConfig(),
		Routes:     getRoutesConfig(),
		StableDate: getStableDate(),
	}
}

//TimingConfig holds the values for executions times of scheduled jobs
type TimingConfig struct {
	Sleep       int
	VolumeCount string
	TotalCount  string
	StatsCount  string
}

func getTimingConfig() TimingConfig {
	return TimingConfig{
		Sleep:       10,
		VolumeCount: ":30",
		TotalCount:  "07:00",
		StatsCount:  "06:00",
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     util.GetEnvVar("server_ip", "localhost"),
		Port:     util.GetEnvVar("server_port", "3000"),
	}
}

// RoutesConfig Holds configuration values for routes to which results should be sent
type RoutesConfig struct {
	VolumeResult string
	StatsResult  string
}

func getRoutesConfig() RoutesConfig {
	return RoutesConfig{
		VolumeResult: "api/app/twitter-data/volumes",
		StatsResult:  "api/app/twitter-data/mean-and-stdev",
	}
}

func getStableDate() time.Time {
	dateStr := os.Getenv("STABLE_DATE")
	if dateStr == "" {
		log.Fatal("No supplied stable date")
	}
	Layout := "2006-01-02"
	date, err := time.Parse(Layout, dateStr)
	if err != nil {
		util.CheckErrFatal(err)
	}
	return date
}
