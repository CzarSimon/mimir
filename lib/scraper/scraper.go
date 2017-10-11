package scraper

import (
  "time"
)

// ScraperInterface Main interface for scraper
type ScraperInterface interface {
  Get(URL string)    (string, error)
  Parse(html string) (Page, error)
}

// Scraper Implementation of ScraperInterface
type Scraper struct {
  UserAgent string
  Subjects  []Subject
}

// Get Retrives the raw html from a page
func (scraper Scraper) Get(URL string) (string, error) {
  return URL, nil
}

// Parse Parses raw html and returns a Page struct if successfull
func (scraper Scraper) Parse(html string) (Page, error) {
  return Page{}, nil
}

// Page Struct for scraped page
type Page struct {
  Title string    `json:"title"`
  Text  string    `json:"text"`
  Date  time.Time `json:"date"`
  Links []string  `json:"links"`
}

// Subject Subject which to look for on page
type Subject struct {
  Name     string   `json:"name"`
  Ticker   string   `json:"ticker"`
  Keywords []string `json:"keywords"`
}
