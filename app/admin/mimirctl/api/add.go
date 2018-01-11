package api

import (
	"fmt"
	"net/http"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
)

// AddLabeledSpam instructs the admin api to store a labeled spam candidate.
func AddLabeledSpam(candidate spam.Candidate) {
	resp, err := makePostRequest(createRoute("spam"), candidate)
	if err != nil {
		checkErr(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		checkErr(fmt.Errorf("Error status. %d", resp.StatusCode))
	}
	fmt.Println("Spam labeled")
}

// AddStock instructs the admin api to store and track a new candidate
func AddStock(newStock stock.Stock) {
	resp, err := makePostRequest(createRoute("stock"), newStock)
	if err != nil {
		checkErr(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		checkErr(fmt.Errorf("Error status. %d", resp.StatusCode))
	}
	fmt.Println("Stock stored")
}

// GetNewStock creates a stock with default
// sugested values based on the supplied ticker.
func GetNewStock(ticker string) stock.Stock {
	return stock.Stock{
		Ticker: ticker,
	}
}
