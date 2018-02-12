package main

import (
	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/go-file-heartbeat/heartbeat"
	"github.com/CzarSimon/util"
)

// Environment variable names.
const (
	TweetDBName = "TWEET_DB"
	AppDBName   = "APP_DB"
)

// Config holds configuration values.
type Config struct {
	TweetDB   endpoint.SQLConfig
	AppDB     endpoint.SQLConfig
	Heartbeat heartbeat.Config
}

func getConfig() Config {
	heartbeatConf, err := heartbeat.NewConfigFromEnv()
	util.CheckErrFatal(err)
	return Config{
		TweetDB:   endpoint.NewPGConfig(TweetDBName),
		AppDB:     endpoint.NewPGConfig(AppDBName),
		Heartbeat: heartbeatConf,
	}
}
