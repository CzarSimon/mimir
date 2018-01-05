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
	return &Env{
		db: util.ConnectPG(config.DB),
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)

	server := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: SetupRoutes(env),
	}

	log.Printf("Running app-server on port: %s\n", config.Server.Port)
	err := server.ListenAndServe()
	util.LogErr(err)
}
