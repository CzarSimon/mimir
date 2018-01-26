package main

import (
	"os"

	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	SERVER_NAME          = "TWEET_HANDLER"
	FilterSpamKey        = "HANDLE_SPAM"
	ShoudFilterSpamValue = "TRUE"
	TweetHandlerPort     = "2000"
	DBHostKey            = "PG_HOST"
	DBPasswordKey        = "PG_PASSWORD"
	RankerHostKey        = "RANKER_HOST"
	RankerPortKey        = "RANKER_PORT"
	SpamFilterHostKey    = "SPAM_FILTER_HOST"
	SpamFilterPortKey    = "SPAM_FILTER_PORT"
)

//Config is the main configuration type
type Config struct {
	filterSpam bool
	ranker     endpoint.ServerAddr
	spamFilter endpoint.ServerAddr
	server     endpoint.ServerAddr
	db         endpoint.SQLConfig
}

// getConfig Sets up inital config
func getConfig() Config {
	return Config{
		filterSpam: getShouldFilterSpamConfig(),
		ranker:     endpoint.NewServerAddr("NEWS_RANKER"),
		spamFilter: endpoint.NewServerAddr("SPAM_FILTER"),
		server:     endpoint.NewServerAddr(SERVER_NAME),
		db:         endpoint.NewPGConfig("DB"),
	}
}

// getShouldFilterSpamConfig Gets config for whether the tweet handler
// should filter for spam
func getShouldFilterSpamConfig() bool {
	shouldFilterSpam := os.Getenv(FilterSpamKey)
	return shouldFilterSpam == ShoudFilterSpamValue
}
