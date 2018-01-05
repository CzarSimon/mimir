package main

import (
	"database/sql"
	"log"
	"net/http"

	endpoint "github.com/CzarSimon/go-endpoint"
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
		TweetDB: connectToDB(config.tweetDB),
		AppDB:   connectToDB(config.appDB),
		Config:  config,
	}
}

// connectToDB Connect to a database, exits if unsuccessfull
func connectToDB(config endpoint.SQLConfig) *sql.DB {
	db, err := config.Connect()
	util.CheckErrFatal(err)
	return db
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
	//fmt.Println(config)
	env := NewEnv(config)

	server := NewServer(env)

	log.Printf("Starting server: %s on port: %s", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	util.CheckErrFatal(err)
}
