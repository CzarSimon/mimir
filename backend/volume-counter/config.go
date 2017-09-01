package main

import "github.com/CzarSimon/util"

//Config holds configurations and environment variables
type Config struct {
	DB     util.PGConfig
	Timing TimingConfig
	Server util.ServerConfig
}

func getConfig() Config {
	pgHost := util.GetEnvVar("PG_HOST", "localhost")
	pgPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	return Config{
		DB:     util.GetPGConfig(pgHost, pgPwd, "simon", "mimirprod"),
		Timing: getTimingConfig(),
		Server: getServerConfig(),
	}
}

//TimingConfig holds the values for executions times of scheduled jobs
type TimingConfig struct {
	Sleep       int
	VolumeCount string
	TotalCount  string
}

func getTimingConfig() TimingConfig {
	return TimingConfig{
		Sleep:       10,
		VolumeCount: ":30",
		TotalCount:  "07:00",
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Protocol: "http",
		Host:     util.GetEnvVar("server_ip", "localhost"),
		Port:     util.GetEnvVar("server_port", "3000"),
	}
}
