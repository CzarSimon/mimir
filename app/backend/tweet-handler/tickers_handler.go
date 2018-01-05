package main

import (
	"errors"
	"net/http"

	"github.com/CzarSimon/util"
)

// GetTrackedTickers Returns tracked tickers to a requestor
func (env *Env) GetTrackedTickers(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrStatus(res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	trackedTickers := CreateTrackedTickersList(env.Tickers, env.Aliases)
	jsonBody, err := trackedTickers.ToJSON()
	if err != nil {
		util.SendErrRes(res, errors.New("Could not retrive tracked tickers"))
		return
	}
	util.SendJSONRes(res, jsonBody)
}
