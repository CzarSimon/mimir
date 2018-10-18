package main

import (
	"github.com/CzarSimon/mimir/app/backend/pkg/mq"
)

const (
	SERVICE_NAME = "news-ranker"
)

type Config struct {
	MQ           MQConfig `env:"MQ"`
	TwitterUsers int64    `env:"TWITTER_USERS" envDefault:"320000000"`
}

type MQConfig struct {
	Host         string `env:"HOST"`
	Port         string `env:"PORT"`
	User         string `env:"USER"`
	Password     string `env:"PASSWORD"`
	Exchange     string `env:"EXCHANGE"`
	ScrapeQueue  string `env:"SCRAPE_QUEUE"`
	ScrapedQueue string `env:"SCRAPED_QUEUE"`
	RankQueue    string `env:"RANK_QUEUE"`
}

func (c Config) Exchange() string {
	return c.MQ.Exchange
}

func (c Config) MQConfig() mq.Config {
	return mq.NewConfig(c.MQ.Host, c.MQ.Port, c.MQ.User, c.MQ.Password)
}
