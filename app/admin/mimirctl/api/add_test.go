package api

import (
	"fmt"
	"testing"
)

func TestGetNewStock(t *testing.T) {
	return
	ticker := "AAPL"
	stock := GetNewStock(ticker)
	if ticker != stock.Ticker {
		t.Errorf("Wrong ticker reviced Expected=%s Got=%s", ticker, stock.Ticker)
	} else {
		fmt.Println(stock)
	}
}
