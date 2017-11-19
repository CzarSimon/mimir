package main

import "github.com/CzarSimon/util"

//Config is the main configuration type
type Config struct {
	server util.ServerConfig
	db     util.PGConfig
}

// getServerConfig Sets up server config
func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: "2000",
	}
}

// getConfig Sets up inital config
func getConfig() Config {
	pgHost := util.GetEnvVar("PG_HOST", "localhost")
	pgPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	return Config{
		server: getServerConfig(),
		db:     util.GetPGConfig(pgHost, pgPwd, "simon", "mimirprod"),
	}
}
