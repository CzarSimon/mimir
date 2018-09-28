package news

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Article contains article info.
type Article struct {
	URLHash      string    `json:"urlHash"`
	URL          string    `json:"url"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	Body         string    `json:"body"`
	DateInserted time.Time `json:"dateInserted"`
	Keywords     []string  `json:"keywords"`
}

// NewArticle creates a new article.
func NewArticle(URL string) Article {
	return Article{
		URLHash: CreateURLHash(URL),
		URL:     URL,
	}
}

// CreateURLHash creates a sha256 hash from a URL.
func CreateURLHash(URL string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(URL)))
}

// Author contains info about an article refererer.
type Author struct {
	ID            int64 `json:"id"`
	FollowerCount int64 `json:"followerCount"`
}

// RankObject contains info to scrape and rank an article.
type RankObject struct {
	Urls     []string  `json:"urls"`
	Subjects []Subject `json:"subjects"`
	Author   Author    `json:"author"`
	Language string    `json:"language"`
}

// Subject subject which to look for in an article.
type Subject struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}
