package main

import (
  "database/sql"
  "log"
  "fmt"
  _ "github.com/lib/pq"
)

func connectPostgres(config pgConfig) *sql.DB {
  connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.user, config.pwd, config.db)
  //fmt.Println(connStr)
  db, err := sql.Open("postgres", connStr)
  if (err != nil) {
    log.Fatalln(err)
  }
  err = db.Ping()
  if (err != nil) {
    log.Fatalln("Error: Could not establish connection with the database")
  }
  return db
}
