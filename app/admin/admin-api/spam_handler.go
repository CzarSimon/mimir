package main

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// spamHandler Handles request related to spam candidates and labeleing resource
func (env *Env) spamHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		env.getSpamCandidates(res, req)
	case http.MethodPost:
		env.labelSpam(res, req)
	default:
		util.SendErrStatus(res, METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

// getSpamCandidates Sends a number of spam candidates to label
func (env *Env) getSpamCandidates(res http.ResponseWriter, req *http.Request) {
	util.SendOK(res)
}

// labelSpam Labels a text with with its spam class
func (env *Env) labelSpam(res http.ResponseWriter, req *http.Request) {
	util.SendOK(res)
}
