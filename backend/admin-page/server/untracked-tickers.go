package main

import (
  "database/sql"
  "encoding/json"
  "net/http"
  "log"
)

type ticker struct {
  Name        string
  Observances int64
}

func mockTickers() []ticker {
  pg := ticker{"PG", 435};
  nvda := ticker{"NVDA", 332}
  gm := ticker{"GM", 78}
  crm := ticker{"CRM", 547}
  ibm := ticker{"IBM", 107}
  return []ticker{pg, nvda, gm, crm, ibm}
}

func checkErr(err error) {
  if (err != nil) {
    log.Fatal(err)
  }
}

func getUntrackedTickers(pg *sql.DB) []ticker {
  rows, err := pg.Query("SELECT ticker, COUNT(*) FROM untracked_tickers GROUP BY ticker ORDER BY COUNT(*) DESC")
  checkErr(err)
  defer rows.Close()
  tickers := make([]ticker, 0)
  var tickerName string
  var observances int64
  for rows.Next() {
    err = rows.Scan(&tickerName, &observances)
    checkErr(err)
    tickers = append(tickers, ticker{tickerName, observances})
  }
  return tickers
}

func (env *Env) sendTickers(w http.ResponseWriter, r *http.Request) {
  tickers := getUntrackedTickers(env.pg);

  js, err := json.Marshal(tickers);
  if (err != nil) {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/xml")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(js)
}
