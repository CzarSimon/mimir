package main

import (
	"database/sql"
	"log"

	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/repository"
    "github.com/CzarSimon/mimir/app/backend/pkg/dbutil"
    "github.com/CzarSimon/mimir/app/backend/pkg/mq"
)

type env struct {
	config      Config
	mqClient    mq.Client
	articleRepo repository.ArticleRepo
	db          *sql.DB
}

func setupEnv(config Config) *env {
	mqClient, err := mq.NewClient(config.MQConfig())
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.DB.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}
	runMigrations(db)

	articleRepo := repository.NewArticleRepo(db)

	return &env{
		config:      config,
		mqClient:    mqClient,
		articleRepo: articleRepo,
		db:          db,
	}
}

func runMigrations(db *sql.DB) error {
	err := dbutil.Migrate("./migrations", "postgres", db)
	if err != nil {
		log.Fatal(err)
	}
}

func (e *env) close() {
	err := e.mqClient.Close()
	if err != nil {
		log.Println(err)
	}

	err = e.db.Close()
	if err != nil {
		log.Println(err)
	}
}

func (e *env) newSubscriptionHandler(queue string, fn handlerFunc) handler {
	return newHandler(queue, e.mqClient, fn)
}

func (e *env) exchange() string {
	return e.config.MQ.Exchange
}

func (e *env) rankQueue() string {
	return e.config.MQ.RankQueue
}

func (e *env) scrapeQueue() string {
	return e.config.MQ.ScrapeQueue
}

func (e *env) scrapedQueue() string {
	return e.config.MQ.ScrapedQueue
}
