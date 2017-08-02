package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

// Article holds the cluster representation of an article
type Article struct {
	Title   string    `json:"title"`
	Ticker  string    `json:"ticker"`
	Date    time.Time `json:"date"`
	URLHash string    `json:"urlHash"`
	Score   Score     `json:"score"`
}

type tempArticle struct {
	Title   string `json:"title"`
	Ticker  string `json:"ticker"`
	URLHash string `json:"urlHash"`
	Score   Score  `json:"score"`
}

// Score holds the subject and reference score of a cluster member
type Score struct {
	SubjectScore   float64 `json:"subjectScore"`
	ReferenceScore float64 `json:"referenceScore"`
}

// calcClusterHash Calculates the cluster hash based on title, ticker and date
func calcClusterHash(title, ticker, date string) string {
	byteHash := sha256.Sum256([]byte(title + ticker + date))
	return fmt.Sprintf("%x", byteHash)
}

// Format Formats the content of an article according to excpectations
func (article *Article) Format() {
	article.Title = strings.ToLower(article.Title)
	article.Ticker = strings.ToUpper(article.Ticker)
}

// formatArticleDate Returns a string formated date in order to calculate cluster hash
func formatArticleDate(date time.Time) string {
	return date.Format("2006-01-02")
}

// ToClusterMember Turns an article into a cluster member
func (article *Article) ToClusterMember() ClusterMember {
	article.Format()
	dateStr := formatArticleDate(article.Date)
	return ClusterMember{
		ClusterHash:    calcClusterHash(article.Title, article.Ticker, dateStr),
		URLHash:        article.URLHash,
		ReferenceScore: article.Score.ReferenceScore,
		SubjectScore:   article.Score.SubjectScore,
	}
}
