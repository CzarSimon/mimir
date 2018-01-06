package main

import (
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/httputil"
)

// Handler Function for dealing with http request / responses
type Handler func(http.ResponseWriter, *http.Request)

// auth Authorizes a request and calls the supplied handler if successfull
func (env *Env) auth(fn http.Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if env.validAccessKey(r) {
			logAuthStatus("Auth Success", r)
			fn.ServeHTTP(w, r)
		} else {
			logAuthStatus("Auth failed", r)
			httputil.SendErr(w, httputil.NotAuthorized)
		}
	}
}

// validAccessKey Checks if the request contains a valid access key
func (env *Env) validAccessKey(req *http.Request) bool {
	accessKeyIsValid := env.Config.accessToken == req.Header.Get("Authorization")
	return accessKeyIsValid && !env.accessTokenExpired()

}

// accessTokenExpired Checks if access token has expired
func (env *Env) accessTokenExpired() bool {
	return env.Config.tokenValidTo.Before(time.Now().UTC())
}

// logAuthStatus Logs outcome of authorization challange
func logAuthStatus(msg string, req *http.Request) {
	log.Printf("%s from: %s, %s", req.URL.Path, req.RemoteAddr, msg)
}
