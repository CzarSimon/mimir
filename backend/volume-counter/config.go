package main

import "os"

//Config holds configurations and environment variables
type Config struct {
	DB     DBConfig
	Timing TimingConfig
	Server ServerConfig
}

func getConfig() Config {
	return Config{
		DB:     getDBConfig(),
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

//DBConfig holds connections details to the database
type DBConfig struct {
	Host, Port, Password, User, DB string
}

func getDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnvVar("DB_HOST", "localhost"),
		Port:     getEnvVar("DB_PORT", "5432"),
		Password: getEnvVar("PG_PASSWORD", ""),
		User:     getEnvVar("DB_USER", "simon"),
		DB:       getEnvVar("DB_NAME", "mimirprod"),
	}
}

//ServerConfig holds values for the reciving server
type ServerConfig struct {
	IP, Port string
}

func getServerConfig() ServerConfig {
	return ServerConfig{
		IP:   getEnvVar("server_ip", "localhost"),
		Port: getEnvVar("server_port", "3000"),
	}
}

func getEnvVar(varKey, nilValue string) string {
	envVar := os.Getenv(varKey)
	if envVar != "" {
		return envVar
	}
	return nilValue
}
