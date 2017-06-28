package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	db        *sql.DB
	rank      RankConfig
	clusterer util.ServerConfig
}

func setupEnvironment(config Config) Env {
	return Env{
		db:        util.ConnectPG(config.db),
		rank:      config.rank,
		clusterer: config.clusterer,
	}
}

func main() {
	config := getConfig()
	env := setupEnvironment(config)
	defer env.db.Close()
	http.HandleFunc("/api/rank-article", env.HandleArticleRanking)
	http.HandleFunc("/api/ranked-article", env.HandleRankedArticle)
	log.Println("Starting server at port: " + config.server.Port)
	http.ListenAndServe(":"+config.server.Port, nil)
}
