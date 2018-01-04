package main

import (
	"fmt"
	"log"
	"os"

	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	SERVER_NAME      = "ADMIN_API"
	ACCESS_TOKEN_KEY = SERVER_NAME + "_ACCESS_KEY"
)

// Config is the main configuration type
type Config struct {
	server      endpoint.ServerAddr
	accessToken string
}

// NewConfig Sets up a new configuration struct
func NewConfig() Config {
	return Config{
		server:      endpoint.NewServerAddr(SERVER_NAME),
		accessToken: getAccessToken(),
	}
}

// getAccessToken Tries to get AccessToken from environment
func getAccessToken() string {
	token, ok := os.LookupEnv(ACCESS_TOKEN_KEY)
	if !ok {
		log.Fatal("Could not find access key")
	}
	fmt.Println(token)
	return token
}
