package main

import (
	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	SERVER_NAME = "NEWS_RANKER"
)

//Config is the main configuration type
type Config struct {
	server    endpoint.ServerAddr
	clusterer endpoint.ServerAddr
	db        endpoint.SQLConfig
	rank      RankConfig
}

func getConfig() Config {
	return Config{
		server:    endpoint.NewServerAddr(SERVER_NAME),
		clusterer: endpoint.NewServerAddr("NEWS_CLUSTERER"),
		db:        endpoint.NewPGConfig("DB"),
		rank:      getRankConfig(),
	}
}

func getRankConfig() RankConfig {
	return RankConfig{
		Path:            "article_ranker/",
		Script:          "scrape_and_rank.pyc",
		Command:         "python3",
		TwitterUsers:    328000000.0,
		ReferenceWeight: 800.0,
	}
}

//RankConfig contains info to start scoring an article
type RankConfig struct {
	Script, Command, Path         string
	TwitterUsers, ReferenceWeight float64
}

//CalcReferenceScore calculates a refence score based on a referers followers and configuration parameters
func (rank RankConfig) CalcReferenceScore(referer Author) float64 {
	return rank.ReferenceWeight * float64(referer.FollowerCount) / rank.TwitterUsers
}
