package main

import (
	"time"
	"strings"
)

// Article holds the cluster representation of an article
type Article struct {
	Title	string    `json:"title"`
	Ticker	string	  `json:"ticker"`	
	Date	time.Time `json:"date"`
	UrlHash string    `json:"urlHash"`
	Score	Score     `json:"score"`
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
		ClusterHash: 	calcClusterHash(article.Title, article.Ticker, dateStr),
		UrlHash: 	article.UrlHash,
		ReferenceScore: article.Score.ReferenceScore,
		SubjectScore: 	article.Score.SubjectScore,
	}
}
