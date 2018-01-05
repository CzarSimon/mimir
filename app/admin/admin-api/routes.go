package main

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/admin/stock", env.stockHandler)
	mux.HandleFunc("/api/admin/untracked-tickers", env.untrackedTickerHandler)
	mux.HandleFunc("/api/admin/spam", env.auth(env.spamHandler))
	mux.HandleFunc("/api/admin/ping", env.auth(env.healthCheck))
	mux.HandleFunc("/api/admin/health", env.healthCheck)
	return mux
}

// healthCheck Handler function for confirming that the admin-api is active
func (env *Env) healthCheck(res http.ResponseWriter, req *http.Request) {
	util.SendOK(res)
}
