package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Env is the environment variables struct
type Env struct {
	db             *sql.DB
	periodMonthMap periodMonthMap
}

// setupEnv Sets up the service environment
func setupEnv(config Config) *Env {
	return &Env{
		db:             util.ConnectPG(config.db),
		periodMonthMap: newPeriodMonthMap(),
	}
}

func main() {
	/* --- Environment setup --- */
	config := getConfig()
	env := setupEnv(config)
	defer env.db.Close()

	server := &http.Server{
		Addr:    ":" + config.server.Port,
		Handler: setupRoutes(env),
	}
	log.Println("Started news-server on port: " + config.server.Port)
	err := server.ListenAndServe()
	util.LogErr(err)
}
