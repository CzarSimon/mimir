package main

import (
	"os"

	"github.com/CzarSimon/util"
)

const (
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
	ranker     util.ServerConfig
	spamFilter util.ServerConfig
	server     util.ServerConfig
	db         util.PGConfig
}

// getConfig Sets up inital config
func getConfig() Config {
	return Config{
		filterSpam: getShouldFilterSpamConfig(),
		ranker:     getRankerConfig(),
		spamFilter: getSpamFilterConfig(),
		server:     getServerConfig(),
		db:         getDBConfig(),
	}
}

// getShouldFilterSpamConfig Gets config for whether the tweet handler
// should filter for spam
func getShouldFilterSpamConfig() bool {
	shouldFilterSpam := os.Getenv(FilterSpamKey)
	return shouldFilterSpam == ShoudFilterSpamValue
}

// getRankerConfig Gets config info for news ranker server
func getRankerConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     util.GetEnvVar(RankerHostKey, "localhost"),
		Port:     util.GetEnvVar(RankerPortKey, "5000"),
	}
}

// getSpamFilterConfig Gets config info for spam filter
func getSpamFilterConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     util.GetEnvVar(SpamFilterHostKey, "localhost"),
		Port:     util.GetEnvVar(SpamFilterPortKey, "1000"),
	}
}

// getServerConfig Sets up server config
func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: TweetHandlerPort,
	}
}

// getDBConfig Sets up database configuration
func getDBConfig() util.PGConfig {
	pgHost := util.GetEnvVar(DBHostKey, "localhost")
	pgPwd := util.GetEnvVar(DBPasswordKey, "pwd")
	return util.GetPGConfig(pgHost, pgPwd, "simon", "mimirprod")
}
