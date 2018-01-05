package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/util"
)

// Handler Function for dealing with http request / responses
type Handler func(http.ResponseWriter, *http.Request)

// auth Authorizes a request and calls the supplied handler if successfull
func (env *Env) auth(handler Handler) Handler {
	return func(res http.ResponseWriter, req *http.Request) {
		if env.validAccessKey(req) {
			logAuthStatus("Auth Success", req)
			handler(res, req)
		} else {
			logAuthStatus("Auth failed", req)
			util.SendErrStatus(
				res, fmt.Errorf("Not authorized"), http.StatusUnauthorized)
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
