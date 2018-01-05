package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// Send Posts a json payload to a specified resource
func Send(payload []byte, URL string) error {
	response, err := http.Post(URL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	if response.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("Error code %d in response", response.StatusCode)
}
