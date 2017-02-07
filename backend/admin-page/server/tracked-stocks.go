package main

import (
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
  checkErrNice(err)
  var stock Stock
  for rows.Next(&stock) {
    stocks = append(stocks, stock)
  }
  return stocks
}
