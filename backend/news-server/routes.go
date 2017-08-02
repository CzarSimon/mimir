package main

import "github.com/julienschmidt/httprouter"

func setupRoutes(env *Env) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/news/:ticker/:top/:period", env.GetNews)
	return router
}
