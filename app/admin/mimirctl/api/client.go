package api

import "errors"

// TODO: Actually implement ping
// Ping Attempts to authenticate towards the admin-api using
// the supplied config, returns an error if unsuccessfull
func Ping(config Config) error {
	if config.Auth.AccessKey == "" {
		return errors.New("Invalid Access key")
	}
	return nil
}
