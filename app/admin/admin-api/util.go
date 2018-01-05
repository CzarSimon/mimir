package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Common http errors
var (
	METHOD_NOT_ALLOWED = fmt.Errorf(
		http.StatusText(http.StatusMethodNotAllowed))
	INTERNAL_SERVER_ERROR = fmt.Errorf(
		http.StatusText(http.StatusInternalServerError))
	BAD_REQUEST = fmt.Errorf(http.StatusText(http.StatusBadRequest))
)

// InternalServerError Returns status code and text for an internal server error
func InternalServerError() (int, error) {
	return http.StatusInternalServerError, INTERNAL_SERVER_ERROR
}

// BadRequest Retusns a status code and text for a bad request
func BadRequest() (int, error) {
	return http.StatusBadRequest, BAD_REQUEST
}

// SendJSON Marshals a body and sends as response
func SendJSON(res http.ResponseWriter, v interface{}) error {
	jsonBody, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return INTERNAL_SERVER_ERROR
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(jsonBody)
	if err != nil {
		log.Println(err)
	}
	return nil
}
