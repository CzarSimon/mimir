package main

import (
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/handler"
)

// SetupRoutes Sets up routes and route handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	setupUserRoutes(mux, env)
	setupTwitterDataRoutes(mux, env)
	setupStockRoutes(mux, env)
	mux.Handle("/health", handler.HealthCheck)
	mux.HandleFunc("/readiness", env.ReadinessCheck)
	return mux
}

// setupUserRoutes Sets up API routes related to users
func setupUserRoutes(mux *http.ServeMux, env *Env) {
	mux.Handle("/api/app/user", handler.New(env.HandleUserRequest))
	mux.Handle("/api/app/user/session", handler.New(env.HandleSessionRequest))
	mux.Handle("/api/app/user/search", handler.New(env.HandleUserSearch))
	mux.Handle("/api/app/user/ticker", handler.New(env.HandleTickerRequest))
}

// setupTwitterDataRoutes Sets up API routes related to twitter data
func setupTwitterDataRoutes(mux *http.ServeMux, env *Env) {
	mux.Handle("/api/app/twitter-data", handler.New(env.HandleGetTwitterDataRequest))
	mux.Handle("/api/app/twitter-data/volumes", handler.New(env.HandleNewVolumesRequest))
	mux.Handle("/api/app/twitter-data/mean-and-stdev", handler.New(env.HandleNewStatsRequest))
}

// setupStockRoutes Sets up API routes related to stock data
func setupStockRoutes(mux *http.ServeMux, env *Env) {
	mux.Handle("/api/app/stock", handler.New(env.HandleStockRequest))
	mux.Handle("/api/app/stocks", handler.New(env.HandleStocksRequest))
	mux.Handle("/api/app/stock/description", handler.New(env.HandleStockRequest))
}

// ReadinessCheck checks that the server is ready to recieve traffic.
func (env *Env) ReadinessCheck(w http.ResponseWriter, r *http.Request) {
	err := env.db.Ping()
	if err != nil {
		httputil.SendErr(w, httputil.InternalServerError)
		return
	}
	httputil.SendOK(w)
}
