package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CzarSimon/mimir/lib/stock"
	"github.com/CzarSimon/util"
)

// stockHandler Handles request related to the stock resource
func (env *Env) stockHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		util.SendOK(res)
	case http.MethodPost:
		env.storeNewStock(res, req)
	default:
		util.SendErrStatus(res, METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

// storeNewStock Derializes and stores a new stock
func (env *Env) storeNewStock(res http.ResponseWriter, req *http.Request) {
	var newStock stock.Stock
	err := util.DecodeJSON(req.Body, &newStock)
	if err != nil {
		log.Println(err)
		util.SendErrStatus(
			res, fmt.Errorf("Could not parse stock"), http.StatusBadRequest)
		return
	}
	log.Printf("Ticker=%s Name=%s", newStock.Ticker, newStock.Name)
	util.SendOK(res)
}
