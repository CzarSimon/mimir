package main

import (
	"net/http"

	"github.com/CzarSimon/httputil/handler"
)

// SetupRoutes Sets up routes and route handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	setupUserRoutes(mux, env)
	setupTwitterDataRoutes(mux, env)
	setupStockRoutes(mux, env)
	mux.Handle("/health", handler.HealthCheck)
	return mux
}

// setupUserRoutes Sets up API routes related to users
func setupUserRoutes(mux *http.ServeMux, env *Env) {
	mux.HandleFunc("/api/app/user", env.HandleUserRequest)
	mux.HandleFunc("/api/app/user/session", env.HandleSessionRequest)
	mux.HandleFunc("/api/app/user/search", env.HandleUserSearch)
	mux.HandleFunc("/api/app/user/ticker", env.HandleTickerRequest)
}

// setupTwitterDataRoutes Sets up API routes related to twitter data
func setupTwitterDataRoutes(mux *http.ServeMux, env *Env) {
	mux.HandleFunc("/api/app/twitter-data", env.HandleGetTwitterDataRequest)
	mux.HandleFunc("/api/app/twitter-data/volumes", env.HandleNewVolumesRequest)
	mux.HandleFunc("/api/app/twitter-data/mean-and-stdev", env.HandleNewStatsRequest)
}

// setupStockRoutes Sets up API routes related to stock data
func setupStockRoutes(mux *http.ServeMux, env *Env) {
	mux.HandleFunc("/api/app/stock", env.HandleStockRequest)
	mux.HandleFunc("/api/app/stocks", env.HandleStocksRequest)
	mux.HandleFunc("/api/app/stock/description", env.HandleStockRequest)
}
