package main

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// untrackedTickerHandler Handles request for the resource untracked tickers
func (env *Env) untrackedTickerHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		env.getUntrackedTickers(res, req)
	default:
		util.SendErrStatus(res, METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

// getUntrackedTickers Gets a list of untracked tickers
func (env *Env) getUntrackedTickers(res http.ResponseWriter, req *http.Request) {
	util.SendOK(res)
}
