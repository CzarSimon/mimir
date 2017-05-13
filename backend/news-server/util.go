package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func checkErrFatal(err error) {
	if err != nil {
		log.Println("Fatal error, exiting.")
		log.Fatal(err.Error())
	}
}

func plainTextRes(res http.ResponseWriter, msg string) {
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write([]byte(fmt.Sprintf("%x", msg)))
}

type JsonRes struct {
	Response string
}

func jsonStringRes(res http.ResponseWriter, msg string) {
	jsonResponse := JsonRes{msg}
	js, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonRes(res, js)
}

func jsonRes(res http.ResponseWriter, js []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(js)
}

func errorResponse(res http.ResponseWriter) {
	jsonStringRes(res, "Error in handling the request")
}

func getDate() string {
	dateFormat := "2006-01-02"
	return time.Now().UTC().Format(dateFormat)
}
