package api

import (
	"fmt"
	"testing"
)

func TestGetUntrackedTickers(t *testing.T) {
	return
	tickers := GetUntrackedTickers()
	for _, ticker := range tickers {
		fmt.Println(ticker)
	}
}

func TestGetStocks(t *testing.T) {
	return
	for _, stock := range GetStocks() {
		fmt.Println(stock)
		fmt.Println()
	}
}

func TestGetSpamCandidates(t *testing.T) {
	return
	for _, candidate := range GetSpamCandidates() {
		fmt.Println(candidate)
	}
}
