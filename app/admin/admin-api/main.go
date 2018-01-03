package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

// Env is the struct for environment objects
type Env struct {
	TweetDB *sql.DB
	AppDB   *sql.DB
	Config  Config
}

// NewEnv Sets up handler environment
func NewEnv(config Config) *Env {
	return &Env{
		Config: config,
	}
}

// NewServer Creates a server with a route handler
func NewServer(env *Env) *http.Server {
	return &http.Server{
		Addr:    ":" + env.Config.server.Port,
		Handler: SetupRoutes(env),
	}
}

func main() {
	config := NewConfig()
	env := NewEnv(config)

	server := NewServer(env)

	log.Printf("Starting server: %s on port: %s", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	util.CheckErrFatal(err)
}
