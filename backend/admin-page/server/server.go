package main

import (
  "log"
  "net/http"
  "database/sql"
  r "gopkg.in/gorethink/gorethink.v2"
)

type Env struct{
  pg    *sql.DB
  rdb   *r.Session
  token string
}

func main() {
  config := getConfig()

  /* ---- DB setup ---- */
  env := &Env{
    pg: connectPostgres(config.pg),
    rdb: connectRethink(config.rdb),
    token: generateToken(),
  }
  defer env.pg.Close()
  defer env.rdb.Close()

  /* ---- Routes ---- */
  http.HandleFunc("/untracked-tickers", env.sendTickers)
  http.HandleFunc("/track-ticker", env.trackTicker)

  /* ---- Starting Server ---- */
  log.Println("Starting server on port " + config.server.port)
  err := http.ListenAndServe(":" + config.server.port, nil)
  checkErr(err)
}
