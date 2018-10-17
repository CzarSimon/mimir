package main

import (
	"log"

	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/repository"
	"github.com/CzarSimon/mimir/app/backend/pkg/mq"
	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

func (e *env) handleRankObjectMessage(msg mq.Message) error {
	rankObject, err := parseRankObject(msg)
	if err != nil {
		return nil
	}

	for _, URL := range rankObject.URLs {
		article, err := e.articleRepo.FindByURL(URL)
		if err == repository.ErrNoSuchArticle {
			e.rankNewArticle(news.NewArticle(URL), rankObject)
			continue
		} else if err != nil {
			log.Println(err)
			continue
		}
		e.rankExistingArticle(article, rankObject)
	}
	return nil
}

func (e *env) rankNewArticle(article news.Article, rankObject news.RankObject) {

}

func (e *env) rankExistingArticle(article news.Article, rankObject news.RankObject) {

}

func parseRankObject(msg mq.Message) (news.RankObject, error) {
	var ro news.RankObject
	err := msg.Decode(&ro)
	return ro, err
}
