package main

import (
	"errors"
	"fmt"
)

// Aliases Map of aliases to tickers
type Aliases map[string]string

// Add Adds alias and ticker to an alias map
func (aliases Aliases) Add(alias, ticker string) error {
	if currentTicker, present := aliases[alias]; present {
		return newAliasPresentError(alias, currentTicker)
	}
	aliases[alias] = ticker
	return nil
}

// newAliasPresentError Creates an error for duplicate alias
func newAliasPresentError(alias, ticker string) error {
	return errors.New(
		fmt.Sprintf("Alias: %s already present, ticker = %s", alias, ticker))
}
