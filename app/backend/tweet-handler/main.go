package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

// SetupServer Creates a server with a route handler
func SetupServer(env *Env) *http.Server {
	return &http.Server{
		Addr:    ":" + env.Config.server.Port,
		Handler: SetupRoutes(env),
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)

	server := SetupServer(env)
	log.Println(env.Config.ranker.ToURL(RankRoute))
	log.Println(env.Config.spamFilter.ToURL(ClassifyRoute))
	log.Printf("Starting %s on port: %s\n", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	util.CheckErrFatal(err)
}
