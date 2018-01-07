package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
)

// stockHandler Handles request related to the stock resource
func (env *Env) stockHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		httputil.SendOK(w)
		return nil
	case http.MethodPost:
		return env.storeNewStock(w, r)
	default:
		return httputil.MethodNotAllowed
	}
}

// storeNewStock Derializes and stores a new stock
func (env *Env) storeNewStock(w http.ResponseWriter, r *http.Request) error {
	var newStock stock.Stock
	err := util.DecodeJSON(r.Body, &newStock)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	log.Printf("Ticker=%s Name=%s", newStock.Ticker, newStock.Name)
	httputil.SendOK(w)
	return nil
}
