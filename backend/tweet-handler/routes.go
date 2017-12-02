package main

import (
	"net/http"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tweet/tickers", env.GetTrackedTickers)
	mux.HandleFunc("/api/tweet", env.ReciveNewTweet)
	return mux
}
