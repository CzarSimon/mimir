package main

import (
	"log"

	"github.com/CzarSimon/mimir/app/backend/pkg/mq"
)

func (e *env) handleScrapedArticleMessage(msg mq.Message) error {
	scrapedArticle, err := parseScrapedArticle(msg)
	if err != nil {
		return err
	}

	e.updateAndStoreScrapedArticle(scrapedArticle)
	return nil
}

func (e *env) updateAndStoreScrapedArticle(scrapedArticle news.ScrapedArticle) {
	referers, err := e.articleRepo.FindArticleReferers(scrapedArticle.Article.ID)
	if err != nil {
		log.Println(err)
		return
	}

	referenceScore := calcReferenceScore(e.config.TwitterUsers, scrapedArticle.Referer, referers...)
	scrapedArticle.Article.ReferenceScore = referenceScore

	err = e.articleRepo.SaveScrapedArticle(scrapedArticle)
	if err != nil {
		log.Println(err)
	}
}

func parseScrapedArticle(msg mq.Message) (news.ScrapedArticle, error) {
	var sa news.ScrapedArticle
	err := msg.Decode(&sa)
	return sa, err
}
