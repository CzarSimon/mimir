package main

import (
  "log"
  "net/http"
  r "gopkg.in/gorethink/gorethink.v2"
)

type Env struct {
  db *r.Session
}

func setupEnvironment(config Config) *Env {
  return &Env{
    db: connectDB(config.db),
  }
}

func main() {
  config := getConfig()
  env := setupEnvironment(config)
  defer env.db.Close()

  http.HandleFunc("/api/cluster-article", env.clusterArticle)

  log.Println("Started server running on port " + config.server.port)
  http.ListenAndServe(":" + config.server.port, nil)
}
