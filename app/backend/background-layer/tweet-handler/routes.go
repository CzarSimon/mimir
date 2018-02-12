package main

import (
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/handler"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/tweet/tickers", handler.New(env.GetTrackedTickers))
	mux.Handle("/api/tweet", handler.New(env.ReciveNewTweet))
	mux.Handle("/health", handler.HealthCheck)
	mux.HandleFunc("/readiness", env.ReadinessCheck)
	return mux
}

// ReadinessCheck checks that the server is ready to recieve traffic.
func (env *Env) ReadinessCheck(w http.ResponseWriter, r *http.Request) {
	err := env.DB.Ping()
	if err != nil {
		httputil.SendErr(w, httputil.InternalServerError)
		return
	}
	httputil.SendOK(w)
}
