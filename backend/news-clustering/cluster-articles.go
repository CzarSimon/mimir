package main

import (
  "log"
  "net/http"
)

func (env *Env) clusterArticle(res http.ResponseWriter, req *http.Request) {
  log.Println("Cluster article endpoint called")
  jsonStringRes(res, "Article clustered")
}
