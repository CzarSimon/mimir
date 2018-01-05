package main

import (
	"errors"
	"fmt"

	"github.com/CzarSimon/mimir/lib/news"
)

// Tickers Map of tracked tickers and their corresponding subjects
type Tickers map[string]news.Subject

// Add Adds ticker to ticker set
func (tickers Tickers) Add(subject news.Subject) error {
	if _, present := tickers[subject.Ticker]; present {
		return newTickerPresentError(subject.Ticker)
	}
	tickers[subject.Ticker] = subject
	return nil
}

// newTickerPresentError Creates an error indicating a ticker is already present
func newTickerPresentError(ticker string) error {
	return errors.New(fmt.Sprintf("Ticker: %s already present", ticker))
}
