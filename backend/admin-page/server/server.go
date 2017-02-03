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

func main() {
  var devMode bool
  flag.BoolVar(&devMode, "dev", false, "sets development mode")
  flag.Parse()
  config := getConfig(devMode)

  /* ---- DB setup ---- */
  env := &Env{
    pg: connectPostgres(config.pg),
    rdb: connectRethink(config.rdb),
    auth: config.auth,
    devMode: config.server.devMode,
  }
  defer env.pg.Close()
  defer env.rdb.Close()

  /* ---- Routes ---- */
  http.Handle("/", http.FileServer(http.Dir(config.server.staticFolder)))
  http.HandleFunc("/login", env.login)
  http.HandleFunc("/untracked-tickers", env.sendTickers)
  http.HandleFunc("/track-ticker", env.trackTicker)

  /* ---- Starting Server ---- */
  log.Println("Starting server on port " + config.server.port)
  log.Println("Serving files from " + config.server.staticFolder)
  err := http.ListenAndServe(":" + config.server.port, nil)
  checkErr(err)
}
