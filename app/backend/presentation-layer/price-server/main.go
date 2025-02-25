package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Env is the struct for environment objects
type Env struct {
	db *sql.DB
}

// SetupEnv Sets up the service enironment
func SetupEnv(config Config) Env {
	db, err := config.DB.Connect()
	util.CheckErrFatal(err)
	return Env{
		db: db,
	}
}

// SetupServer Creates a server with a route handler
func SetupServer(env *Env, port string) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: SetupRoutes(env),
	}
}

func main() {
	config := GetConfig()
	env := SetupEnv(config)
	defer env.db.Close()

	server := SetupServer(&env, config.Server.Port)

	log.Printf("Starting %s at port: %s\n", SERVER_NAME, config.Server.Port)
	err := server.ListenAndServe()
	util.CheckErr(err)
}
