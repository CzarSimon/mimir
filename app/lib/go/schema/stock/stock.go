package stock

import (
	"database/sql"
	"fmt"
)

// Stock Contains information about a stock and issuing company
type Stock struct {
	Ticker      string   `json:"ticker"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	ImageURL    string   `json:"imageUrl,omitempty"`
	Website     string   `json:"website,omitempty"`
	Keywords    []string `json:"keyworks,omitempty"`
}

// String Returns a string representation of a stock
func (stock Stock) String() string {
	return fmt.Sprintf(
		"Ticker=%s \nName=%s\nDescription=%s\nImageURL=%s\nWebsite=%s",
		stock.Ticker, stock.Name, stock.Description, stock.ImageURL, stock.Website)
}

// NullStock nullable version of the stock type to be used for database querying.
type NullStock struct {
	Ticker      sql.NullString
	Name        sql.NullString
	Description sql.NullString
	ImageURL    sql.NullString
	Website     sql.NullString
}

// Stock turns a nullable stock into a stock struct
func (ns NullStock) Stock() Stock {
	return Stock{
		Ticker:      ns.Ticker.String,
		Name:        ns.Name.String,
		Description: ns.Description.String,
		ImageURL:    ns.ImageURL.String,
		Website:     ns.Website.String,
	}
}
