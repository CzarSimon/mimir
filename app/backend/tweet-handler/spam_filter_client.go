package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/CzarSimon/util"
)

const (
	ClassifyRoute = "classify"
)

// FilterSpam Querys spam filter to check if tweet is spam
func (env *Env) FilterSpam(tweet Tweet) bool {
	if !env.Config.filterSpam {
		return false
	}
	endpoint := env.Config.spamFilter.ToURL(ClassifyRoute)
	spamResult, err := queryForSpam(spam.NewCandidate(tweet.Text), endpoint)
	if err != nil {
		log.Println(err)
	}
	log.Println(spamResult)
	return spamResult.IsSpam()
}

// queryForSpam Resuests the spam filter to evaluate a request
func queryForSpam(candidate spam.Candidate, URL string) (spam.Candidate, error) {
	body, err := json.Marshal(candidate)
	if err != nil {
		return candidate, err
	}
	resp, err := http.Post(URL, BodyType, bytes.NewBuffer(body))
	if err != nil {
		return candidate, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return candidate, fmt.Errorf("Failed spam filter call. Status=%d", resp.StatusCode)
	}
	err = util.DecodeJSON(resp.Body, &candidate)
	return candidate, err
}
