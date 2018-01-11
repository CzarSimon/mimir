package api

import (
	"fmt"
	"net/http"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
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
	company, err := getCompanyInfo(ticker)
	if err != nil || company.Symbol != ticker {
		return stock.Stock{Ticker: ticker}
	}
	return company.Stock()
}

// getCompanyInfo gets company info from the external IEX trading api.
func getCompanyInfo(ticker string) (*companyInfo, error) {
	apiEndpoint := fmt.Sprintf("https://ws-api.iextrading.com/1.0/stock/%s/company", ticker)
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var company companyInfo
	err = util.DecodeJSON(resp.Body, &company)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

// companyInfo helper struct to recive company info from the IEX api.
type companyInfo struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

// Stock converts companyInfo to the stock structure.
func (company companyInfo) Stock() stock.Stock {
	return stock.Stock{
		Ticker:      company.Symbol,
		Name:        company.CompanyName,
		Description: company.Description,
		Website:     company.Website,
	}
}
