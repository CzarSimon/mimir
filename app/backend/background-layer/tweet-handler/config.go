package main

import (
	"os"

	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	SERVER_NAME          = "TWEET_HANDLER"
	FilterSpamKey        = "HANDLE_SPAM"
	ShoudFilterSpamValue = "TRUE"
)

//Config is the main configuration type
type Config struct {
	filterSpam    bool
	ranker        endpoint.ServerAddr
	spamFilter    endpoint.ServerAddr
	server        endpoint.ServerAddr
	db            endpoint.SQLConfig
	forbiddenURLs ForbiddenURLs
}

// getConfig Sets up inital config
func getConfig() Config {
	return Config{
		filterSpam:    getShouldFilterSpamConfig(),
		ranker:        endpoint.NewServerAddr("NEWS_RANKER"),
		spamFilter:    endpoint.NewServerAddr("SPAM_FILTER"),
		server:        endpoint.NewServerAddr(SERVER_NAME),
		db:            endpoint.NewPGConfig("DB"),
		forbiddenURLs: GetForbiddenURLs(),
	}
}

// getShouldFilterSpamConfig Gets config for whether the tweet handler
// should filter for spam
func getShouldFilterSpamConfig() bool {
	shouldFilterSpam := os.Getenv(FilterSpamKey)
	return shouldFilterSpam == ShoudFilterSpamValue
}
