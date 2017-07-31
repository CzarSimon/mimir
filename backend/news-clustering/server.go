package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
	r "gopkg.in/gorethink/gorethink.v2"
)

// Env Holds common environment objects and database connections
type Env struct {
	db    *r.Session
	queue QueueMap
	pg    *sql.DB
}

// setupEnv Sets up environment for the server
func setupEnv(config Config) *Env {
	return &Env{
		db:    connectDB(config.db),
		queue: newQueueMap(),
		pg:    util.ConnectPG(config.pg),
	}
}

func main() {
	config := getConfig()
	env := setupEnv(config)
	defer env.db.Close()

	http.HandleFunc("/api/cluster-article", env.clusterArticle)

	log.Println("Started server running on port " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
