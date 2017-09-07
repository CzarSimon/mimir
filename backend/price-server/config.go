package main

import "github.com/CzarSimon/util"

// Config Holds configuration values
type Config struct {
	DB     util.PGConfig
	Server util.ServerConfig
}

// GetConfig Returns a new Config struct based on evnirionment variables
func GetConfig() Config {
	pdbHost := util.GetEnvVar("PG_HOST", "localhost")
	pdbPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	return Config{
		DB:     util.GetPGConfig(pdbHost, pdbPwd, "simon", "mimirprod"),
		Server: getServerConfig(),
	}
}

// getServerConfig Retruns a new server configuration
func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: "4000",
	}
}
