package main

import (
  "fmt"
  "net/http"
  "database/sql"
)

type Env struct{
  pg *sql.DB
}

func main() {
  config := getConfig();
  /* ---- DB setup ---- */
  //Need to be added and handler funcs wrapped
  env := &Env{
    pg: connectPostgres(config.pg),
  }
  defer env.pg.Close()

  /* ---- Routes ---- */
  http.HandleFunc("/untracked-tickers", env.sendTickers);

  /* ---- Starting Server ---- */
  fmt.Println("Starting server on port " + config.server.port);
  err := http.ListenAndServe(":" + config.server.port, nil)
  if (err != nil) {
    fmt.Println(err)
  }
}
