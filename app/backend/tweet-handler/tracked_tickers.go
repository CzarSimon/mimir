package main

import "encoding/json"

// TrackedTickers Slice of all tracked tickers and aliases
type TrackedTickers []string

// ToJSON Serializes TrackedTickers to JSON
func (tickers TrackedTickers) ToJSON() ([]byte, error) {
	return json.Marshal(&tickers)
}

// CreateTrackedTickersList Combines tickers and aliases into TrackedTickers
func CreateTrackedTickersList(tickers Tickers, aliases Aliases) TrackedTickers {
	trackedTickers := make(TrackedTickers, 0, len(tickers)+len(aliases))
	for ticker, _ := range tickers {
		trackedTickers = append(trackedTickers, ticker)
	}
	for alias, _ := range aliases {
		trackedTickers = append(trackedTickers, alias)
	}
	return trackedTickers
}
