package main

import (
  "os"
  "fmt"
)

func getConfig(devMode bool) Config {
  server := getServerConfig(devMode)
  pg := getPgConfig()
  rdb := getRdbConfig()
  auth := getAuthConfig()
  cert := getCertConfig(devMode)
  return Config{server, pg, rdb, auth, cert}
}

func isDev() bool {
  return false
}

type Config struct{
  server  serverConfig
  pg      pgConfig
  rdb     rdbConfig
  auth    authConfig
  cert    certConfig
}

func getServerConfig(devMode bool) serverConfig {
  return serverConfig{"8000", "./../build", devMode}
}

type serverConfig struct{
  port, staticFolder  string
  devMode bool
}

func getPgConfig() pgConfig {
  return pgConfig{
    host: getEnvVar("pg_host", "localhost"),
    port: "5432",
    db: "mimirprod",
    pwd: getEnvVar("pg_pwd", "56error78"),
    user: "simon",
  }
}

type pgConfig struct {
  host, port, db, pwd, user string
}

func getRdbConfig() rdbConfig {
  return rdbConfig{
    host: getEnvVar("rdb_host", "localhost"),
    port: "28015",
    db: "mimir_app_server",
  }
}

type rdbConfig struct {
  host, port, db  string
}

type authConfig struct {
  pwdHash, token, salt string
}

func getAuthConfig() authConfig {
  return authConfig{
    pwdHash: "14e403b056aae6c4ab42f114e93024064aeb1f793aecbe08efdf28eba905661b",
    token: generateToken(),
    salt: "46b05499459ee36074658bc0d3736dfa6d65cc1bccfb82f58a45397cce1ccbdc",
  }
}

type certConfig struct {
  folder, host string
  devMode bool
}

func getCertConfig(devMode bool) certConfig {
  return certConfig{
    folder: "./certs",
    host: "mimirapp.co",
    devMode: devMode,
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
