package stock

// Stock Contains information about a stock and issuing company
type Stock struct {
  Ticker      string   `json:"ticker"`
  Name        string   `json:"name,omitempty"`
  Description string   `json:"description,omitempty"`
  ImageURL    string   `json:"imageUrl,omitempty"`
  Website     string   `json:"website,omitempty"`
  Keywords    []string `json:"keyworks,omitempty"`
}
