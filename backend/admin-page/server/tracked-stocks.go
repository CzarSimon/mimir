package main

import (
  "database/sql"
  "encoding/json"
  "net/http"
  r "gopkg.in/gorethink/gorethink.v2"
)

type Stock struct {
  Name, Ticker, Description, ImageUrl string
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
  //parse Body
  var ticker string
  ticker = "O"
  //
  responseMessage, err := deleteStock(ticker, env)
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
  sqlString := "UPDATE stocks SET is_tracked=FALSE WHERE ticker=$1"
  err := db.QueryRow(sqlString, ticker).Scan()
  return err
}

func deleteFromRdb(ticker string, session *r.Session) error {
  _, err := r.Table("stocks").Delete().GetAllByIndex("ticker", ticker).RunWrite(session)
  return err
}
