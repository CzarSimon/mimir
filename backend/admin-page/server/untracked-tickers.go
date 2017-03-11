package main

import (
  "database/sql"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "log"
  "time"
  "net/http"
  "github.com/lib/pq"
  r "gopkg.in/gorethink/gorethink.v2"
)

type ticker struct {
  Name        string
  Observances int64
}

func (env *Env) sendTickers(res http.ResponseWriter, req *http.Request) {
  if env.authenticate(res, req) != nil {
    return
  }
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
  if env.authenticate(res, req) != nil {
    return
  }
  decoder := json.NewDecoder(req.Body)
  var newTicker tickerInfo
  if err := decoder.Decode(&newTicker); err != io.EOF {
    checkErr(err)
  }
  defer req.Body.Close()
  if err := insertNewTicker(newTicker, env.rdb); err == nil {
    insertTickerInPostgres(newTicker, env.pg)
    jsonStringRes(res, fmt.Sprintf("%s successfully added", newTicker.Name))
  } else {
    log.Println("No ticker added error:", err.Error())
    jsonStringRes(res, err.Error())
  }
}

func getUntrackedTickers(pg *sql.DB) []ticker {
  rows, err := pg.Query("SELECT ticker, COUNT(*) FROM untracked_tickers WHERE ticker NOT IN (SELECT ticker FROM stocks WHERE is_tracked=TRUE) GROUP BY ticker ORDER BY COUNT(*) DESC LIMIT 100")
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

func insertNewTicker(ticker tickerInfo, session *r.Session) error {
  if tickerStored(ticker.Ticker, session) {
    return errors.New(fmt.Sprintf("Ticker %s already added", ticker.Ticker))
  }
  stats := emptyStats()
  res, err := r.Table("stocks").Insert(map[string]interface{}{
    "ticker": ticker.Ticker,
    "name": ticker.Name,
    "description": ticker.Description,
    "website": ticker.Website,
    "imageUrl": ticker.ImageUrl,
    "mean": stats,
    "stdev": stats,
  }).RunWrite(session)
  if res.Inserted > 0 {
    log.Println("Storing ticker", ticker.Ticker, "in rdb database")
  }
  return err
}

func insertTickerInPostgres(ticker tickerInfo, db *sql.DB) {
  date := time.Now().UTC().AddDate(0, 0, 1)
  stmt, err := db.Prepare("INSERT INTO stocks(ticker, name, storedat) VALUES ($1,$2,$3)")
  if err != nil {
    checkErrNice(err)
    return
  }
  _, err = stmt.Exec(ticker.Ticker, ticker.Name, date)
  if err != nil {
    if pqErr := err.(*pq.Error); pqErr.Code == "23505" {
      stmt, err = db.Prepare("UPDATE stocks SET is_tracked=TRUE, storedat=$1 WHERE ticker=$2")
      _, err = stmt.Exec(date, ticker.Ticker)
      checkErrNice(err)
    }
  }
}

func tickerStored(ticker string, session *r.Session) bool {
  res, err := r.Table("stocks").GetAll(ticker).OptArgs(r.GetAllOpts{
    Index: "ticker",
  }).Run(session)
  defer res.Close()
  checkErrNice(err)
  return !res.IsNil()
}

func emptyStats() map[string]map[int]float64 {
  stats := make(map[int]float64)
  allStats := make(map[string]map[int]float64)
  for i := 0; i < 24; i++ { stats[i] = 0.0 }
  allStats["busdays"], allStats["weekend_days"] = stats, stats
  return allStats
}
