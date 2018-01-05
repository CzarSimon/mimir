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
	queue QueueMap
	db    *sql.DB
}

// setupEnv Sets up environment for the server
func setupEnv(config Config) *Env {
	return &Env{
		queue: newQueueMap(),
		db:    util.ConnectPG(config.db),
	}
}

func main() {
	config := getConfig()
	env := setupEnv(config)
	defer env.db.Close()

	http.HandleFunc("/api/cluster-article", env.ParseArticleAndCluster)

	log.Println("Started server running on port " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
