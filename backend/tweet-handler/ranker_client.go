package main

import (
	"log"
	"strconv"

	"github.com/CzarSimon/mimir/lib/news"
)

// SendLinksToRanker Sends attached links to news ranker along with mentioned subjects
func (env *Env) SendLinksToRanker(tweet Tweet, subjects []news.Subject) {
	log.Println(env.Config.ranker.ToURL("api/rank-article"))
	log.Println(tweet.Text)
}

// twitterUserToAuthor Converts a TwitterUser to Author type
func twitterUserToAuthor(user TwitterUser) (news.Author, error) {
	var author news.Author
	author.ID, err := strconv.Atoi(user.ID)
	if err != nil {
		return author, err
	}
	author.FollowerCount, err = strconv.Atoi(user.Followers)
	return author, err
}
