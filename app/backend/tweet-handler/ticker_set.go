package main

// TickerSet Set of tickers
type TickerSet map[string]bool

// Add Adds a ticker to TickerSet
func (tickerSet TickerSet) Add(ticker string) {
	tickerSet[ticker] = true
}

// NewTickerSet Creates a new empty TickerSet
func NewTickerSet() TickerSet {
	return make(TickerSet)
}

// ToSlice Converts a ticker set to a slice
func (tickerSet TickerSet) ToSlice() []string {
	tickerSlice := make([]string, 0)
	for ticker, _ := range tickerSet {
		tickerSlice = append(tickerSlice, ticker)
	}
	return tickerSlice
}
