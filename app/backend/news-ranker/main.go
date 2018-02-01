package main

import (
	"database/sql"
	"log"
	"net/http"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	db        *sql.DB
	rank      RankConfig
	clusterer endpoint.ServerAddr
}

// SetupEnv sets up environment.
func SetupEnv(config Config) *Env {
	db, err := config.db.Connect()
	util.CheckErrFatal(err)
	return &Env{
		db:        db,
		rank:      config.rank,
		clusterer: config.clusterer,
	}
}

// SetupServer sets up server and routes.
func SetupServer(env *Env, config Config) *http.Server {
	return &http.Server{
		Addr:    ":" + config.server.Port,
		Handler: SetupRoutes(env),
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)
	defer env.db.Close()

	server := SetupServer(env, config)

	log.Println("Starting server at port: " + config.server.Port)
	err := server.ListenAndServe()
	util.CheckErr(err)
}
