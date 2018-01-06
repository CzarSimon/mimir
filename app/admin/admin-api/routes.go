package main

import (
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/auth"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	check := auth.NewWrapper(env.validAccessKey)
	mux.HandleFunc("/api/admin/stock", env.stockHandler)
	mux.HandleFunc("/api/admin/untracked-tickers", env.untrackedTickerHandler)
	mux.Handle("/api/admin/spam", check.Wrap(httputil.NewHandler(env.spamHandler)))
	mux.Handle("/api/admin/ping", check.Wrap(httputil.HealthCheck))
	mux.Handle("/api/admin/health", httputil.HealthCheck)
	return mux
}
