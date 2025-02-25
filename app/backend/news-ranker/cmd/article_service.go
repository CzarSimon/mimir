package main

import (
	"log"

	"github.com/CzarSimon/mimir/app/backend/pkg/mq"
	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

func (e *env) handleScrapedArticleMessage(msg mq.Message) error {
	scrapedArticle, err := parseScrapedArticle(msg)
	if err != nil {
		return err
	}

	article, err := e.updateAndStoreScrapedArticle(scrapedArticle)
	if err != nil {
		log.Println(err)
		return nil
	}

	e.clusterArticle(article)
	return nil
}

func (e *env) updateAndStoreScrapedArticle(scrapedArticle news.ScrapedArticle) (news.Article, error) {
	referers, err := e.articleRepo.FindArticleReferers(scrapedArticle.Article.ID)
	if err != nil {
		return news.Article{}, err
	}

	mergedReferers := mergeReferers(referers, scrapedArticle.Referer)
	referenceScore := calcReferenceScore(e.config.TwitterUsers, mergedReferers...)
	scrapedArticle.Article.ReferenceScore = referenceScore

	err = e.articleRepo.SaveScrapedArticle(scrapedArticle)
	if err != nil {
		return news.Article{}, err
	}
	return scrapedArticle.Article, nil
}

func parseScrapedArticle(msg mq.Message) (news.ScrapedArticle, error) {
	var sa news.ScrapedArticle
	err := msg.Decode(&sa)
	return sa, err
}

func mergeReferers(referers []news.Referer, newReferer news.Referer) []news.Referer {
	merged := make([]news.Referer, len(referers))
	copy(merged, referers)

	for _, referer := range referers {
		if referer.ExternalID == newReferer.ExternalID {
			return merged
		}
	}
	merged = append(merged, newReferer)
	return merged
}
