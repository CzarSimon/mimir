package main

import (
	"log"

	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/domain"
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
	article.ReferenceScore = calcReferenceScore(e.config.TwitterUsers, rankObject.Author)
	scrapeTarget := newScrapeTarget(article, rankObject)

	err := e.mqClient.Send(scrapeTarget, e.exchange(), e.scrapeQueue())
	if err != nil {
		log.Println(err)
	}
}

func (e *env) rankExistingArticle(article news.Article, rankObject news.RankObject) {
	update, err := e.getArticleUpdate(article, rankObject)
	if err != nil {
		log.Println(err)
		return
	}

	switch update.Type {
	case domain.NEW_SUBJECTS_AND_REFERENCES:
		e.rankWithNewSubjectsAndReferences(update)
	case domain.NEW_SUBJECTS:
		e.rankWithNewSubjects(update)
	case domain.NEW_REFERENCES:
		e.rankWithNewReferences(update)
	default:
		log.Printf("Taking no action on update type: %d for article: %s\n",
			update.Type, article.ID)
	}
}

func (e *env) rankWithNewSubjectsAndReferences(update domain.ArticleUpdate) {

}

func (e *env) rankWithNewSubjects(update domain.ArticleUpdate) {

}

func (e *env) rankWithNewReferences(update domain.ArticleUpdate) {

}

func (e *env) getArticleUpdate(article news.Article, rankObject news.RankObject) (domain.ArticleUpdate, error) {

	return domain.ArticleUpdate{}, nil
}

func parseRankObject(msg mq.Message) (news.RankObject, error) {
	var ro news.RankObject
	err := msg.Decode(&ro)
	return ro, err
}

func calcReferenceScore(twitterUsers int64, references ...news.Author) float64 {
	var totalReferences int64
	for _, reference := range references {
		totalReferences += reference.FollowerCount
	}
	return float64(totalReferences) / float64(twitterUsers)
}

func newScrapeTarget(article news.Article, ro news.RankObject) news.ScrapeTarget {
	return news.ScrapeTarget{
		URL:            article.URL,
		Subjects:       ro.Subjects,
		ReferenceScore: article.ReferenceScore,
		ArticleID:      article.ID,
	}
}
