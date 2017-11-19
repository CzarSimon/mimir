package main

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tweet/tickers", env.GetTrackedTickers)
	mux.HandleFunc("/api/tweet", util.PlaceholderHandler)
	return mux
}
