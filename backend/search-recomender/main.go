package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	db     *sql.DB
	config Config
}

// SetupEnv Sets up the service enironment
func SetupEnv(config Config) Env {
	return Env{
		db:     util.ConnectPG(config.db),
		config: config,
	}
}

// SetupServer Creates a server with a route handler
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

	server := SetupServer(&env)

	log.Println("Starting server at port: " + config.server.Port)
	err := server.ListenAndServe()
	util.CheckErr(err)
}
