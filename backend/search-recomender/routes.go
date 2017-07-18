package main

import "net/http"

// SetupRoutes Sets up routes and route handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/search-sugestions", env.GetSearchSugestions)
	return mux
}
