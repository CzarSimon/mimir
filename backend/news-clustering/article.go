package main

import (
	"strings"
)

//Article holds the cluster representation of an article
type Article struct {
	Title, Ticker, Date, UrlHash string
	Score                        Score
}

func (article *Article) Format() {
	article.Title = strings.ToLower(article.Title)
	article.Ticker = strings.ToUpper(article.Ticker)
}
