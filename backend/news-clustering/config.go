package main

import (
  "fmt"
  "os"
)

type Config struct {
  server ServerConfig
  db DBConfig
}

func getConfig() Config {
  return Config{
    server: getServerConfig(),
    db: getDBConfig(),
  }
}

type ServerConfig struct {
  port string
}

func getServerConfig() ServerConfig {
  return ServerConfig{"6000"}
}

type DBConfig struct {
  host, port, db string
}

func getDBConfig() DBConfig {
  return DBConfig{
    host: getEnvVar("db_host", "localhost"),
    port: "28015",
    db: "mimir_news_db",
  }
}

func getEnvVar(varKey, nilValue string) string {
  envVar := os.Getenv(varKey)
  fmt.Println(envVar)
  if envVar != "" {
    return envVar
  } else {
    return nilValue
  }
}
