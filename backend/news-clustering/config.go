package main

import "github.com/CzarSimon/util"

// Config Holds service configurations
type Config struct {
	server util.ServerConfig
	db     DBConfig
	pg     util.PGConfig
}

func getConfig() Config {
	return Config{
		server: getServerConfig(),
		db:     getDBConfig(),
		pg:     util.GetPGConfig("localhost", "pwd", "simon", "mimirprod"),
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: "6000",
	}
}

func getPGConfig() util.PGConfig {
	return util.PGConfig{}
}

// DBConfig Holds connection info about a rethinkdb instance
type DBConfig struct {
	host, port, db string
}

func getDBConfig() DBConfig {
	return DBConfig{
		host: util.GetEnvVar("db_host", "localhost"),
		port: "28015",
		db:   "mimir_news_db",
	}
}
