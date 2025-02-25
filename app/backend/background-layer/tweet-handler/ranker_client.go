package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema"
	"github.com/CzarSimon/mimir/app/lib/go/schema/news"
)

const (
	RankRoute = "api/rank-article"
)

// SendLinksToRanker Sends attached links to news ranker along with mentioned subjects
func (env *Env) SendLinksToRanker(tweet schema.Tweet, subjects []news.Subject) {
	rankObject, err := env.createJsonRankObject(tweet, subjects)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.Post(
		env.Config.ranker.ToURL(RankRoute), httputil.JSON, bytes.NewBuffer(rankObject))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to send object to ranker. Status = %d\n", resp.StatusCode)
	}
}

// createJsonRankObject Creates json serialized rank object
func (env *Env) createJsonRankObject(tweet schema.Tweet, subjects []news.Subject) ([]byte, error) {
	rankObject, err := env.createRankObject(tweet, subjects)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(rankObject)
}

// createRankObject Truns a tweet and list of subjects to a rankt object
func (env *Env) createRankObject(tweet schema.Tweet, subjects []news.Subject) (news.RankObject, error) {
	author, err := twitterUserToAuthor(tweet.User)
	return news.RankObject{
		Urls:     env.Config.forbiddenURLs.FilterURLs(tweet.GetURLs()),
		Subjects: subjects,
		Author:   author,
		Language: tweet.Language,
	}, err
}

// twitterUserToAuthor Converts a TwitterUser to Author type
func twitterUserToAuthor(user schema.TwitterUser) (news.Author, error) {
	id, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return news.Author{}, err
	}
	return news.Author{
		ID:            id,
		FollowerCount: user.Followers,
	}, nil
}
