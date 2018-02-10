package main

import (
	"log"
	"os"
	"time"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
)

//
const (
	TweetDBName  = "TWEET_DB"
	AppDBName    = "APP_DB"
	StabeDateKey = "STABLE_DATE"
)

// Config holds configuration values.
type Config struct {
	TweetDB    endpoint.SQLConfig
	AppDB      endpoint.SQLConfig
	StableDate time.Time
}

// GetConfig sets up a new config struct based on environemnt variables.
func GetConfig() Config {
	return Config{
		TweetDB:    endpoint.NewPGConfig(TweetDBName),
		AppDB:      endpoint.NewPGConfig(AppDBName),
		StableDate: getStableDate(),
	}
}

// getStableDate get the first stable date to use for the stats calculation.
func getStableDate() time.Time {
	dateStr := os.Getenv(StabeDateKey)
	if dateStr == "" {
		log.Fatal("No supplied stable date")
	}
	Layout := "2006-01-02"
	date, err := time.Parse(Layout, dateStr)
	util.CheckErrFatal(err)
	return date
}
