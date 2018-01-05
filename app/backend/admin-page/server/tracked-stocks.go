package main

import (
  "database/sql"
  "encoding/json"
  "io"
  "net/http"
  r "gopkg.in/gorethink/gorethink.v2"
)

type Stock struct {
  Name, Ticker, Description, ImageUrl, Website string
}

func (env *Env) sendStockInfo(res http.ResponseWriter, req *http.Request) {
  if env.authenticate(res, req) != nil {
    return
  }
  stocks := getTrackedStocks(env.rdb)
  js, err := json.Marshal(stocks)
  if (err != nil) {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }
  jsonRes(res, js)
}

func getTrackedStocks(session *r.Session) []Stock {
  stocks := make([]Stock, 0)
  rows, err := r.Table("stocks").Run(session)
  defer rows.Close()
  checkErrNice(err)
  var stock Stock
  for rows.Next(&stock) {
    stocks = append(stocks, stock)
  }
  return stocks
}

func (env *Env) untrackStock(res http.ResponseWriter, req *http.Request) {
  if env.authenticate(res, req) != nil {
    return
  }
  decoder := json.NewDecoder(req.Body)
  var stock Stock
  if err := decoder.Decode(&stock); err != io.EOF {
    checkErrNice(err)
  }
  responseMessage, err := deleteStock(stock.Ticker, env)
  checkErrNice(err)
  jsonStringRes(res, responseMessage)
}

func deleteStock(ticker string, env *Env) (string, error) {
  message := "Failed to untrack"
  err := deleteFromPg(ticker, env.pg)
  if err == nil {
    if err = deleteFromRdb(ticker, env.rdb); err == nil {
      message = "Stoped tracking ticker"
    }
  }
  return message, err
}

func deleteFromPg(ticker string, db *sql.DB) error {
  stmt, err := db.Prepare("UPDATE stocks SET is_tracked=FALSE WHERE ticker=$1")
  if err != nil {
    return err
  }
  _, err = stmt.Exec(ticker)
  return err
}

func deleteFromRdb(ticker string, session *r.Session) error {
  _, err := r.Table("stocks").GetAllByIndex("ticker", ticker).Delete().RunWrite(session)
  return err
}

type StockInfo struct {
  Ticker, Description string
}

func (env *Env) updateStockInfo(res http.ResponseWriter, req *http.Request) {
  if env.authenticate(res, req) != nil {
    return
  }
  decoder := json.NewDecoder(req.Body)
  var stockInfo StockInfo
  if err := decoder.Decode(&stockInfo); err != io.EOF {
    checkErrNice(err)
  }
  err := storeStockUpdate(stockInfo, env.rdb)
  checkErrNice(err)
  jsonStringRes(res, "Stock info updated")
}

func storeStockUpdate(stockInfo StockInfo, session *r.Session) error {
  _, err := r.Table("stocks").GetAllByIndex("ticker", stockInfo.Ticker).Update(map[string]interface{}{
    "description": stockInfo.Description,
  }).RunWrite(session)
  return err
}
