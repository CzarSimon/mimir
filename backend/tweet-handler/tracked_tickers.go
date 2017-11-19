package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/github.com/CzarSimon/mimir/lib/news"
)

// TrackedTickers Slice of all tracked tickers and aliases
type TrackedTickers []string

// ToJSON Serializes TrackedTickers to JSON
func (tickers TrackedTickers) ToJSON() ([]byte, error) {
	return json.Marshal(&tickers)
}

// CreateTrackedTickersList Combines tickers and aliases into TrackedTickers
func CreateTrackedTickersList(tickers TickerSet, aliases Aliases) TrackedTickers {
	trackedTickers := make(TrackedTickers, len(tickers)+len(aliases))
	i := 0
	for ticker, _ := range tickers {
		trackedTickers[i] = ticker
		i++
	}
	for alias, _ := range aliases {
		trackedTickers[i] = alias
		i++
	}
	return trackedTickers
}

// Aliases Map of aliases to tickers
type Aliases map[string]string

// Add Adds alias and ticker to an alias map
func (aliases Aliases) Add(alias, ticker string) error {
	if currentTicker, present := aliases[alias]; present {
		return errors.New(
			fmt.Sprintf("Alias: %s already present, ticker = %s", alias, currentTicker))
	}
	aliases[alias] = ticker
	return nil
}

// TickerSet Set of tracked tickers
type TickerSet map[string]news.Subject

// Add Adds ticker to ticker set
func (tickers TickerSet) Add(ticker string) error {
	if _, present := tickers[ticker]; present {
		return errors.New(fmt.Sprintf("Ticker: %s already present", ticker))
	}
	tickers[ticker] = true
	return nil
}
