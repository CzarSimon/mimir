package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

func connectPostgres(config pgConfig) *sql.DB {
  connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.user, config.pwd, config.db)
  db, err := sql.Open("postgres", connStr)
  checkErr(err)
  err = db.Ping()
  checkErr(err)
  return db
}
