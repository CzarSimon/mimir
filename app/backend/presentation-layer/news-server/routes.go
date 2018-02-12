package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/julienschmidt/httprouter"
)

// SetupRoutes sets up routes for the server.
func SetupRoutes(env *Env) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/news/:ticker/:top/:period", env.GetNews)
	router.GET("/health", HealthCheck)
	router.GET("/readiness", env.ReadinessCheck)
	return router
}

// HealthCheck responsiveness check of the server.
func HealthCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httputil.SendOK(w)
}

// ReadinessCheck checks that the server is ready to recieve traffic.
func (env *Env) ReadinessCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := env.db.Ping()
	if err != nil {
		log.Println(err)
		httputil.SendErr(w, httputil.InternalServerError)
		return
	}
	httputil.SendOK(w)
}
