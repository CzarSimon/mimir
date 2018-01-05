package main

import (
	"log"
	"os"
	"strconv"
	"time"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
)

const (
	TWEET_DB_NAME         = "TWEET_DB"
	APP_DB_NAME           = "APP_DB"
	SERVER_NAME           = "ADMIN_API"
	ACCESS_TOKEN_KEY      = SERVER_NAME + "_ACCESS_KEY"
	TOKEN_EXPIRIY_MINUTES = "TOKEN_EXPIRIY_MINUTES"
)

// Config is the main configuration type
type Config struct {
	server       endpoint.ServerAddr
	tweetDB      endpoint.SQLConfig
	appDB        endpoint.SQLConfig
	accessToken  string
	tokenValidTo time.Time
}

// NewConfig Sets up a new configuration struct
func NewConfig() Config {
	return Config{
		server:       endpoint.NewServerAddr(SERVER_NAME),
		tweetDB:      endpoint.NewPGConfig(TWEET_DB_NAME),
		appDB:        endpoint.NewPGConfig(APP_DB_NAME),
		accessToken:  getAccessToken(),
		tokenValidTo: getTokenValidTo(),
	}
}

// getAccessToken Tries to get AccessToken from environment
func getAccessToken() string {
	token, ok := os.LookupEnv(ACCESS_TOKEN_KEY)
	if !ok {
		log.Fatal("Could not find access key")
	}
	return token
}

// getTokenValidTo Returns that expiry time of the access token
func getTokenValidTo() time.Time {
	now := time.Now().UTC()
	value, ok := os.LookupEnv(TOKEN_EXPIRIY_MINUTES)
	if !ok {
		log.Fatal("Could not find token expiry time")
	}
	validMinutes, err := strconv.Atoi(value)
	util.CheckErrFatal(err)
	return now.Add(time.Minute * time.Duration(validMinutes))
}
