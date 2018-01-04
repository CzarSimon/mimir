package main

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/admin/ping", env.auth(env.healthCheck))
	mux.HandleFunc("/admin/health", env.healthCheck)
	return mux
}

func (env *Env) healthCheck(res http.ResponseWriter, req *http.Request) {
	util.SendOK(res)
}
