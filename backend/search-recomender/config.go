package main

import (
	"fmt"

	"github.com/CzarSimon/util"
)

//Config is the main configuration type
type Config struct {
	server util.ServerConfig
	db     util.PGConfig
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: "7000",
	}
}

func getConfig() Config {
	pgHost := util.GetEnvVar("PG_HOST", "localhost")
	pgPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	fmt.Println(pgHost, pgPwd)
	return Config{
		server: getServerConfig(),
		db:     util.GetPGConfig(pgHost, pgPwd, "simon", "mimirprod"),
	}
}
