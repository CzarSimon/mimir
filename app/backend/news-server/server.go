package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
	_ "github.com/lib/pq"
)

//Env is the environment variables struct
type Env struct {
	db             *sql.DB
	config         Config
	periodMonthMap periodMonthMap
}

// setupEnv Sets up the service environment
func setupEnv(config Config) *Env {
	db, err := config.db.Connect()
	util.CheckErrFatal(err)
	return &Env{
		db:             db,
		config:         config,
		periodMonthMap: newPeriodMonthMap(),
	}
}

// SetupServer creates a server with a route handler.
func SetupServer(env *Env) *http.Server {
	return &http.Server{
		Addr:    ":" + env.config.server.Port,
		Handler: SetupRoutes(env),
	}
}

func main() {
	/* --- Environment setup --- */
	config := getConfig()
	env := setupEnv(config)
	defer env.db.Close()

	server := SetupServer(env)

	log.Printf("Started %s on port: %s\n", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	util.LogErr(err)
}
