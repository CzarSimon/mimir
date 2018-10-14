package news

import (
	"crypto/sha256"
	"fmt"
	"strings"
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

// RankObject contains info to scrape and rank an article.
type RankObject struct {
	Urls     []string  `json:"urls"`
	Subjects []Subject `json:"subjects"`
	Author   Author    `json:"author"`
	Language string    `json:"language"`
}

func (ro RankObject) String() string {
	urls := strings.Join(ro.Urls, ",")

	subjects := make([]string, 0, len(ro.Subjects))
	for _, subject := range ro.Subjects {
		subjects = append(subjects, subject.String())
	}

	return fmt.Sprintf("RankObject(urls=[%s] subjects=[%s], author=%s, language=%s)",
		urls, strings.Join(subjects, ","), ro.Author, ro.Language)
}

// Author contains info about an article refererer.
type Author struct {
	ID            string `json:"id"`
	FollowerCount int64  `json:"followerCount"`
}

func (a Author) String() string {
	return fmt.Sprintf("Author(id=%s followerCount=%d)", a.ID, a.FollowerCount)
}

// Subject subject which to look for in an article.
type Subject struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

func (s Subject) String() string {
	return fmt.Sprintf("Subject(name=%s ticker=%s)", s.Name, s.Ticker)
}
