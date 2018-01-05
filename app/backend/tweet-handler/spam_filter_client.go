package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

const (
	Spam    = "SPAM"
	NonSpam = "NON-SPAM"
)

// FilterSpam Querys spam filter to check if tweet is spam
func (env *Env) FilterSpam(tweet Tweet) bool {
	if !env.Config.filterSpam {
		return false
	}
	endpoint := getSpamFilterURL(env.Config.spamFilter)
	spamResult, err := queryForSpam(newSpamQuery(tweet), endpoint)
	if err != nil {
		util.LogErr(err)
		return false
	}
	spamResult.Log(tweet.Text)
	return spamResult.IsSpam()
}

// queryForSpam Resuests the spam filter to evaluate a request
func queryForSpam(spamQuery SpamQuery, url string) (SpamResult, error) {
	var spamResult SpamResult
	body, err := json.Marshal(spamQuery)
	if err != nil {
		return spamResult, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return spamResult, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return spamResult, errors.New("Unsuccessfull call to spam filter")
	}
	err = util.DecodeJSON(resp.Body, &spamResult)
	return spamResult, err
}

// getSpamFilterURL Gets the URL to which to send spam filter requests
func getSpamFilterURL(spamFilter util.ServerConfig) string {
	return spamFilter.ToURL("/classify")
}

// SpamQuery Query to send to spam filter for evaluation
type SpamQuery struct {
	Text string `json:"text"`
}

// newSpamQuery Creates a spam query from a tweet
func newSpamQuery(tweet Tweet) SpamQuery {
	return SpamQuery{
		Text: tweet.Text,
	}
}

// SpamResult Result from spam query
type SpamResult struct {
	Result string `json:"result"`
}

// IsSpam Evaluates if a spam result indicated spam or not
func (spamResult SpamResult) IsSpam() bool {
	return !(spamResult.Result == NonSpam)
}

// Log Logs a spam result
func (spamResult SpamResult) Log(text string) {
	log.Printf("%s : %s\n", spamResult.Result, text)
}
