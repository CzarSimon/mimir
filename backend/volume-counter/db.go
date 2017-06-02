package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectPostgres(config DBConfig) *sql.DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.User, config.Password, config.DB)
	db, err := sql.Open("postgres", connStr)
	checkErrFatal(err)
	err = db.Ping()
	checkErrFatal(err)
	return db
}
