package stock

import (
	"errors"

	"github.com/CzarSimon/util"
)

// InitalTickers default list of inital tickers.s
var InitalTickers = Tickers{"AAPL", "FB", "TSLA", "TWTR", "AMZN"}

// Tickers slice of tickers that can be queried and inserted into a postgres database.
type Tickers []string

// Scan scans a slice of tickers.
func (tickers *Tickers) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	(*tickers) = util.BytesToStrSlice(bytes)
	return nil
}

// Add adds new ticker if it is not yet added.
func (tickers *Tickers) Add(ticker string) error {
	for _, currentTicker := range *tickers {
		if currentTicker == ticker {
			return errors.New("Ticker already added")
		}
	}
	*tickers = append(*tickers, ticker)
	return nil
}

// Remove removes ticker is it is present.
func (tickers *Tickers) Remove(ticker string) error {
	filteredTickers := make([]string, 0)
	tickerPresent := false
	for _, currentTicker := range *tickers {
		if currentTicker != ticker {
			filteredTickers = append(filteredTickers, currentTicker)
		} else {
			tickerPresent = true
		}
	}
	if !tickerPresent {
		return errors.New("Ticker not present")
	}
	*tickers = filteredTickers
	return nil
}
