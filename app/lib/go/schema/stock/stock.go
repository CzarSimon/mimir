package stock

import (
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
    "Ticker=%s Name=%s\nDescription=%s\nImageURL=%s\nWebsite=%s",
    stock.Ticker, stock.Name, stock.Description, stock.ImageURL, stock.Website)
}
