package main

import (
  "database/sql"
  "encoding/json"
  "net/http"
  "log"
  "io"
  "fmt"
)


type ticker struct {
  Name        string
  Observances int64
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
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(js)
}


type tickerInfo struct {
  Ticker, Name, Description, Token string
}


func (env *Env) trackTicker(w http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var newTicker tickerInfo
  if err := decoder.Decode(&newTicker); err != io.EOF {
    checkErr(err)
  }
  defer req.Body.Close()
  log.Println("Storing ticker", newTicker.Ticker, "in rdb database")

  w.Header().Set("Content-Type", "text/plain")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write([]byte(fmt.Sprintf("%s successfully added", newTicker.Name)))
}
