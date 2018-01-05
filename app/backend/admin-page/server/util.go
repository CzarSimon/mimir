package main

import (
  "log"
)

func checkErr(err error) {
  if err != nil {
    log.Fatal(err.Error())
  }
}

func checkErrNice(err error) {
  if err != nil {
    log.Println(err.Error())
  }
}
