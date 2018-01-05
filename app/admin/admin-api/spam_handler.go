package main

import (
	"net/http"

	"github.com/CzarSimon/util"
)

// spamHandler Handles request related to spam candidates and labeleing resource
func (env *Env) spamHandler(res http.ResponseWriter, req *http.Request) {
	status := http.StatusMethodNotAllowed
	err := METHOD_NOT_ALLOWED
	switch req.Method {
	case http.MethodGet:
		status, err = env.getSpamCandidates(res, req)
	case http.MethodPost:
		status, err = env.labelSpam(res, req)
	}
	if err != nil {
		util.SendErrStatus(res, err, status)
	}
}

// getSpamCandidates Sends a number of spam candidates to label
func (env *Env) getSpamCandidates(res http.ResponseWriter, req *http.Request) (int, error) {
	util.SendOK(res)
	return http.StatusOK, nil
}

// labelSpam Labels a text with with its spam class
func (env *Env) labelSpam(res http.ResponseWriter, req *http.Request) (int, error) {
	util.SendOK(res)
	return http.StatusOK, nil
}
