package api

import (
	"fmt"
	"testing"
)

func TestGetUntrackedTickers(t *testing.T) {
	tickers := GetUntrackedTickers()
	for _, ticker := range tickers {
		fmt.Println(ticker)
	}
}

func TestGetStocks(t *testing.T) {
	for _, stock := range GetStocks() {
		fmt.Println(stock)
		fmt.Println("")
	}
}
