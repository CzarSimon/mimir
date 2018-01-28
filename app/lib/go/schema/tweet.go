package schema

import "time"

// Tweet holds tweet information.
type Tweet struct {
	Text     string        `json:"text"`
	Entities TweetEntities `json:"entities"`
	ID       string        `json:"id_str"`
	Language string        `json:"lang"`
	Date     string        `json:"created_at"`
	User     TwitterUser   `json:"user"`
}

// TweetEntities entities refereced in tweet.
type TweetEntities struct {
	Hashtags []Hashtag `json:"hashtags"`
	Symbols  []Symbol  `json:"symbols"`
	URLs     []URL     `json:"urls"`
}

// Hashtag hashtag referenced in tweet.
type Hashtag struct {
	Text string `json:"text"`
}

// Symbol symbol referenced in tweet.
type Symbol struct {
	Text string `json:"text"`
}

// URL linked URL in tweet.
type URL struct {
	ExpandedURL string `json:"expanded_url"`
	ShortURL    string `json:"url"`
}

// Get gets the ticker url, expanded if possible.
func (url URL) Get() string {
	if url.ExpandedURL != "" {
		return url.ExpandedURL
	}
	return url.ShortURL
}

// TwitterUser information about user who created tweet.
type TwitterUser struct {
	ID        string `json:"id_str"`
	Followers int64  `json:"followers_count"`
}

// GetDate Parses tweet date into a timestamp
func (tweet Tweet) GetDate() time.Time {
	timestamp, err := time.Parse(time.RubyDate, tweet.Date)
	if err != nil {
		return time.Now().UTC()
	}
	return timestamp
}

// GetURLS Gets the tweet urls as strings
func (tweet Tweet) GetURLs() []string {
	urls := make([]string, len(tweet.Entities.URLs))
	for i, url := range tweet.Entities.URLs {
		urls[i] = url.Get()
	}
	return urls
}
