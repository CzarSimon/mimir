package main

import (
	"fmt"
	"log"

	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

func main() {
	env := setupEnv(newConfig())
	env.listenForRankObjects()
}

func (env *env) listenForRankObjects() {
	msgChan, err := env.mqClient.Subscribe(env.config.MQ.RankQueue, SERVICE_NAME)
	if err != nil {
		log.Println(err)
		return
	}

	for msg := range msgChan {
		var ro news.RankObject
		err = msg.Decode(&ro)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(ro)
		msg.Ack()
		if err != nil {
			log.Println(err)
		}
	}
}

func newConfig() Config {
	return Config{
		MQ: MQConfig{
			Host:         "localhost",
			Port:         "5672",
			User:         "newsranker",
			Password:     "password",
			Exchange:     "x-news",
			ScrapeQueue:  "q-scrape-targets",
			ScrapedQueue: "q-scraped-articles",
			RankQueue:    "q-rank-objects",
		},
	}
}
