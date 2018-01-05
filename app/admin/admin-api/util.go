package main

import (
	"fmt"
	"net/http"
)

var METHOD_NOT_ALLOWED = fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed))
