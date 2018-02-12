package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/util"
)

// GetTrackedTickers Returns tracked tickers to a requestor
func (env *Env) GetTrackedTickers(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.MethodNotAllowed
	}
	trackedTickers := CreateTrackedTickersList(env.Tickers, env.Aliases)
	jsonBody, err := trackedTickers.ToJSON()
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	util.SendJSONRes(w, jsonBody)
	return nil
}
