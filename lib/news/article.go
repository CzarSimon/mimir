package news

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Article contains article info
type Article struct {
	URLHash      string    `json:"urlHash"`
	URL          string    `json:"url"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	Body         string    `json:"body"`
	DateInserted time.Time `json:"dateInserted"`
	Keywords     []string  `json:"keywords"`
}

// NewArticle Creates a new article
func NewArticle(URL string) Article {
	return Article{
		URLHash: CreateURLHash(URL),
		URL:     URL,
	}
}

// CreateURLHash Creates a sha256 hash from a URL
func CreateURLHash(URL string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(URL)))
}
