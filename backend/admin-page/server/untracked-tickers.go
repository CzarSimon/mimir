package main

import (
  "database/sql"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "log"
  "net/http"
  r "gopkg.in/gorethink/gorethink.v2"
)

type ticker struct {
  Name        string
  Observances int64
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

func (env *Env) sendTickers(res http.ResponseWriter, r *http.Request) {
  tickers := getUntrackedTickers(env.pg);
  js, err := json.Marshal(tickers);
  if (err != nil) {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }
  jsonRes(res, js)
}

type tickerInfo struct {
  Ticker, Name, Description, ImageUrl, Website string
}

func (env *Env) trackTicker(res http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var newTicker tickerInfo
  if err := decoder.Decode(&newTicker); err != io.EOF {
    checkErr(err)
  }
  defer req.Body.Close()
  if err := insertNewTicker(newTicker, env.rdb); err == nil {
    jsonStringRes(res, fmt.Sprintf("%s successfully added", newTicker.Name))
  } else {
    log.Println("No ticker added error:", err.Error())
    jsonStringRes(res, err.Error())
  }
}

func insertNewTicker(ticker tickerInfo, session *r.Session) error {
  if tickerStored(ticker.Ticker, session) {
    return errors.New(fmt.Sprintf("Ticker %s already added", ticker.Ticker))
  }
  res, err := r.Table("stocks").Insert(map[string]interface{}{
    "ticker": ticker.Ticker,
    "name": ticker.Name,
    "description": ticker.Description,
    "website": ticker.Website,
    "imageUrl": ticker.ImageUrl,
  }).RunWrite(session)
  if res.Inserted > 0 {
    log.Println("Storing ticker", ticker.Ticker, "in rdb database")
  }
  return err
}

func tickerStored(ticker string, session *r.Session) bool {
  res, err := r.Table("stocks").GetAll(ticker).OptArgs(r.GetAllOpts{
    Index: "ticker",
  }).Run(session)
  defer res.Close()
  checkErrNice(err)
  return !res.IsNil()
}
