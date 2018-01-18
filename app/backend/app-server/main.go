package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Env is the environment variables struct
type Env struct {
	db *sql.DB
}

// SetupEnv Sets up environment
func SetupEnv(config Config) *Env {
	db, err := config.DB.Connect()
	util.CheckErrFatal(err)
	return &Env{
		db: db,
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)

	server := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: SetupRoutes(env),
	}

	log.Printf("Running %s on port: %s\n", SERVER_NAME, config.Server.Port)
	err := server.ListenAndServe()
	util.LogErr(err)
}
