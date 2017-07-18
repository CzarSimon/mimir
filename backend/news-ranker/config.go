package main

import "github.com/CzarSimon/util"

//Config is the main configuration type
type Config struct {
	server    util.ServerConfig
	clusterer util.ServerConfig
	db        util.PGConfig
	rank      RankConfig
}

func getConfig() Config {
	pgHost := util.GetEnvVar("PG_HOST", "localhost")
	pgPwd := util.GetEnvVar("PG_PASSWORD", "pwd")
	return Config{
		server:    getServerConfig(),
		clusterer: getClustererConfig(),
		db:        util.GetPGConfig(pgHost, pgPwd, "simon", "mimirprod"),
		rank:      getRankConfig(),
	}
}

func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: "5000",
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

func getClustererConfig() util.ServerConfig {
	return util.ServerConfig{
		Host:     util.GetEnvVar("cluster_host", "localhost"),
		Port:     util.GetEnvVar("cluster_port", "6000"),
		Protocol: "http",
	}
}
