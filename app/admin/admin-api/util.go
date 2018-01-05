package main

import (
	"fmt"
	"net/http"
)

var METHOD_NOT_ALLOWED = fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed))

// HandlerFunc Handler function for a http request
// returns a status code and error reference
type HandlerFunc func(http.ResponseWriter, *http.Request) (int, error)
