package main


import (
  "fmt"
  r "gopkg.in/gorethink/gorethink.v2"
)

func connectRethink(config rdbConfig) *r.Session {
  address := fmt.Sprintf("%s:%s", config.host, config.port)
  session, err := r.Connect(r.ConnectOpts{
    Address: address,
    Database: config.db,
  })
  checkErr(err)
  return session
}
