package news

// Subject subject which to look for in an article.
type Subject struct {
	Name     string   `json:"name"`
	Ticker   string   `json:"ticker"`
	Keywords []string `json:"keywords,omitempty"`
}
