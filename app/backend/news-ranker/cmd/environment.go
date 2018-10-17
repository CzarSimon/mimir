package main

import (
	"log"

	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/repository"
	"github.com/CzarSimon/mimir/app/backend/pkg/mq"
)

type env struct {
	config      Config
	mqClient    mq.Client
	articleRepo repository.ArticleRepo
}

func setupEnv(config Config) *env {
	mqClient, err := mq.NewClient(config.MQConfig())
	if err != nil {
		log.Fatal(err)
	}
	return &env{
		config:   config,
		mqClient: mqClient,
	}
}

func (e *env) close() {
	err := e.mqClient.Close()
	if err != nil {
		log.Println(err)
	}
}

func (e *env) newSubscriptionHandler(queue string, fn handlerFunc) handler {
	return newHandler(queue, e.mqClient, fn)
}
