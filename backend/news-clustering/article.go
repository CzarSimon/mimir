package main

import (
  "strings"
)

type Article struct {
  Title, Ticker, Date, UrlHash string
  Score Score
}

func (article *Article) Format() {
  article.Title = strings.ToLower(article.Title)
  article.Ticker = strings.ToUpper(article.Ticker)
}
