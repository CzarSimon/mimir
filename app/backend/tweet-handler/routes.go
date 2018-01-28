package main

import (
	"net/http"

	"github.com/CzarSimon/httputil/handler"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/tweet/tickers", handler.New(env.GetTrackedTickers))
	mux.Handle("/api/tweet", handler.New(env.ReciveNewTweet))
	return mux
}
