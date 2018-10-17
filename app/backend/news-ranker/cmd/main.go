package main

import (
	"sync"

	"github.com/CzarSimon/mimir/app/backend/pkg/mq"
)

func main() {
	conf := newConfig()
	e := setupEnv(conf)
	wg := &sync.WaitGroup{}

	rankObjectHandler := e.newSubscriptionHandler(conf.MQ.RankQueue, e.handleRankObjectMessage)
	articlesHandler := e.newSubscriptionHandler(conf.MQ.ScrapedQueue, emptyHandle)
	go handleSubscription(rankObjectHandler, wg)
	go handleSubscription(articlesHandler, wg)

	wg.Wait()
}

func emptyHandle(msg mq.Message) error {
	return nil
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
