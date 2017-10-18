package scraper

// Subject Subject which to look for on page
type Subject struct {
	Name     string   `json:"name"`
	Ticker   string   `json:"ticker"`
	Keywords []string `json:"keywords"`
}
