package main

import (
	"log"
	"net/http"

	r "gopkg.in/gorethink/gorethink.v2"
)

//Env is the environment variables struct
type Env struct {
	db *r.Session
}

func setupEnvironment(config Config) *Env {
	return &Env{
		db: connectDB(config.db),
	}
}

func main() {
	/* --- Environment setup --- */
	config := getConfig()
	env := setupEnvironment(config)
	defer env.db.Close()

	server := &http.Server{
		Addr:    ":" + config.server.port,
		Handler: setupRoutes(env),
	}
	log.Println("Started news-server on port: " + config.server.port)
	err := server.ListenAndServe()
	checkErr(err)
}
