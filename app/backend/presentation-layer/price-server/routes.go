package main

import (
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/handler"
)

// SetupRoutes Sets up routes and route handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/price/historical", handler.New(env.GetHistoricalPrices))
	mux.Handle("/api/price/latest", handler.New(env.GetLatestPrices))
	mux.Handle("/health", handler.HealthCheck)
	mux.HandleFunc("/readiness", env.ReadinessCheck)
	return mux
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
