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
