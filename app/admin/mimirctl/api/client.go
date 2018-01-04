package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Ping Attempts to authenticate towards the admin-api using
// the supplied config, returns an error if unsuccessfull
func Ping(config Config) error {
	res, err := performRequest(newGetRequest("admin/ping", config))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"Unable to verify credentials. ResponseCode: %d", res.StatusCode)
	}
	return nil
}

// performRequest Performs an http request and returs the result
func performRequest(req *http.Request) (*http.Response, error) {
	client := newClient()
	return client.Do(req)
}

// newGetRequest Creates a new GET request to the admin-api and
// embends the api access key
func newGetRequest(route string, config Config) *http.Request {
	return newRequest(http.MethodGet, route, config, nil)
}

// newGetRequest Creates a new Post request to the admin-api and
// embends the api access key
func newPostRequest(route string, config Config, body io.Reader) *http.Request {
	req := newRequest(http.MethodPost, route, config, body)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// newGetRequest Creates a new http request of a supplied method
// to the admin-api and embends the api access key
func newRequest(method, route string, config Config, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, config.API.ToURL(route), body)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal()
	}
	req.Header.Set("Authorization", config.Auth.AccessKey)
	return req
}

// newClient Sets up a new client for use
func newClient() *http.Client {
	const TIMEOUT_SECONDS = 5
	return &http.Client{
		Timeout: time.Second * TIMEOUT_SECONDS,
	}
}
