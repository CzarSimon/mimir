package main

import (
  "log"
  "encoding/json"
  "fmt"
  "net/http"
)

func checkErr(err error) {
  if err != nil {
    log.Println(err.Error())
  }
}

func checkErrFatal(err error) {
  if err != nil {
    log.Println("Fatal error, exiting.")
    log.Fatal(err.Error())
  }
}

func plainTextRes(res http.ResponseWriter, msg string) {
  res.Header().Set("Content-Type", "text/plain")
  res.Header().Set("Access-Control-Allow-Origin", "*")
  res.Write([]byte(fmt.Sprintf("%x", msg)))
}

type JsonRes struct {
  Response string
}

func jsonStringRes(res http.ResponseWriter, msg string) {
  jsonResponse := JsonRes{msg}
  js, err := json.Marshal(jsonResponse)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }
  jsonRes(res, js)
}

func jsonRes(res http.ResponseWriter, js []byte) {
  res.Header().Set("Content-Type", "application/json")
  res.Header().Set("Access-Control-Allow-Origin", "*")
  res.Write(js)
}
