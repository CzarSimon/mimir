package main

import (
	"log"

	"github.com/CzarSimon/mimir/lib/news"
)

// SendLinksToRanker Sends attached links to news ranker along with mentioned subjects
func (env *Env) SendLinksToRanker(tweet Tweet, subjects []news.Subject) {
	log.Println(env.Config.ranker.ToURL("api/rank-article"))
	log.Println(tweet.Text)
}

// createRankObject Truns a tweet and list of subjects to a rankt object
func createRankObject(tweet Tweet, subjects []news.Subject) (news.RankObject, error) {
	return news.RankObject{
		Subjects: subjects,
		Language: tweet.Language,
	}, nil
}
