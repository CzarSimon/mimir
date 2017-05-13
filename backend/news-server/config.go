package main

import (
	"fmt"
	"os"
)

//Config contains internal config variables for the server
type Config struct {
	server ServerConfig
	db     DBConfig
}

func getConfig() Config {
	return Config{
		server: getServerConfig(),
		db:     getDBConfig(),
	}
}

//ServerConfig contains info about the current server
type ServerConfig struct {
	port string
}

func getServerConfig() ServerConfig {
	return ServerConfig{"5050"}
}

//DBConfig contains info to connect to a rethinkdb instance
type DBConfig struct {
	host, port, db string
}

func getDBConfig() DBConfig {
	return DBConfig{
		host: getEnvVar("db_host", "localhost"),
		port: "28015",
		db:   "mimir_news_db",
	}
}

func getEnvVar(varKey, nilValue string) string {
	envVar := os.Getenv(varKey)
	fmt.Println(envVar)
	if envVar != "" {
		return envVar
	}
	return nilValue
}
