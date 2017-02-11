package main

import (
  "flag"
  "log"
  "net/http"
  "database/sql"
  r "gopkg.in/gorethink/gorethink.v2"
)

type Env struct{
  pg      *sql.DB
  rdb     *r.Session
  auth    authConfig
  devMode bool
}

func parseFlags() bool {
  var devMode bool
  flag.BoolVar(&devMode, "dev", false, "sets development mode")
  flag.Parse()
  return devMode
}

func setupEnvironment(conf config) *Env {
  env := &Env{
    pg: connectPostgres(conf.pg),
    rdb: connectRethink(conf.rdb),
    auth: conf.auth,
    devMode: conf.server.devMode,
  }
  return env
}

func main() {
  devMode := parseFlags()
  config := getConfig(devMode)

  /* ---- Environment setup ---- */
  env := setupEnvironment(config)
  defer env.pg.Close()
  defer env.rdb.Close()

  /* ---- Routes ---- */
  http.Handle("/", http.FileServer(http.Dir(config.server.staticFolder)))
  http.HandleFunc("/login", env.login)
  http.HandleFunc("/tracked-stocks", env.sendStockInfo)
  http.HandleFunc("/untrack-stock", env.untrackStock)
  http.HandleFunc("/untracked-tickers", env.sendTickers)
  http.HandleFunc("/track-ticker", env.trackTicker)

  /* ---- Starting Server ---- */
  log.Println("Starting server on port " + config.server.port)
  checkErr(http.ListenAndServe(":" + config.server.port, nil))
}
