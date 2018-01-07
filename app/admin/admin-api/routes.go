package main

import (
	"net/http"

	"github.com/CzarSimon/httputil/auth"
	"github.com/CzarSimon/httputil/handler"
)

// SetupRoutes sets up routes and handlers.
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	check := auth.NewWrapper(env.validAccessKey)
	mux.Handle("/api/admin/stock", handler.New(env.stockHandler))
	mux.Handle("/api/admin/untracked-tickers", handler.New(env.untrackedTickerHandler))
	mux.Handle("/api/admin/spam", check.Wrap(handler.New(env.spamHandler)))
	mux.Handle("/api/admin/ping", check.Wrap(handler.HealthCheck))
	mux.Handle("/api/admin/health", handler.HealthCheck)
	return mux
}
