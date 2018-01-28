package main

import (
	"net/http"

	"github.com/CzarSimon/httputil/handler"
)

// SetupRoutes Sets up routes and route handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/search", handler.New(env.SearchStocks))
	mux.Handle("/api/search/sugestions", handler.New(env.GetSearchSugestions))
	return mux
}
