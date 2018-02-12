package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects.
type Env struct {
	db     *sql.DB
	config Config
}

// SetupEnv sets up the service enironment.
func SetupEnv(config Config) *Env {
	db, err := config.db.Connect()
	util.CheckErrFatal(err)
	return &Env{
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

	log.Printf("Starting %s at port: %s\n", config.server.Port, SERVER_NAME)
	err := server.ListenAndServe()
	util.CheckErr(err)
}
