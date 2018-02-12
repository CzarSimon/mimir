package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
	_ "github.com/lib/pq"
)

// Env Holds common environment objects and database connections
type Env struct {
	queue  QueueMap
	db     *sql.DB
	config Config
}

// SetupEnv Sets up environment for the server
func SetupEnv(config Config) *Env {
	db, err := config.db.Connect()
	util.CheckErrFatal(err)
	return &Env{
		queue:  newQueueMap(),
		db:     db,
		config: config,
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
	config := getConfig()
	env := SetupEnv(config)
	defer env.db.Close()

	server := SetupServer(env)

	log.Printf("Started %s on port: %s\n", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	util.CheckErr(err)
}
