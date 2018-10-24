package main

import (
	"log"

	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/domain"
	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/repository"
	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

func (e *env) clusterArticle(article news.Article) {
	subjects, err := e.articleRepo.FindArticleSubjects(article.ID)
	if err != nil && err != repository.ErrNoSubjects {
		log.Println(err)
		return
	}

	for _, subject := range subjects {
		e.clusterArticleWithSubject(article, subject)
	}
}

func (e *env) clusterArticleWithSubject(article news.Article, subject news.Subject) {
	clusterHash := domain.CalcClusterHash(article.Title, subject.Symbol, article.ArticleDate)

	cluster, err := e.clusterRepo.FindByHash(clusterHash)
	if err == repository.ErrNoSuchCluster {
		e.createNewCluster(clusterHash, article, subject)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	e.updateArticleCluster(cluster, article, subject)
}

func (e *env) createNewCluster(clusterHash string, article news.Article, subject news.Subject) {
	members := createNewClusterMemebers(clusterHash, article, subject)

	cluster := domain.NewArticleCluster(
		article.Title, subject.Symbol, article.ArticleDate,
		article.ID, members[0].Score(), members)

	err := e.clusterRepo.Save(*cluster)
	if err != nil {
		log.Println(err)
	}
}

func (e *env) updateArticleCluster(cluster domain.ArticleCluster, article news.Article, subject news.Subject) {
	updateClusterMembers(&cluster, article, subject)
	cluster.ElectLeaderAndScore()
	e.clusterRepo.Save(cluster)
}

func updateClusterMembers(cluster *domain.ArticleCluster, article news.Article, subject news.Subject) {
	newMember := createNewClusterMember(cluster, article, subject)
	members := make([]domain.ClusterMember, len(cluster.Members))
	copy(members, cluster.Members)

	for i, _ := range cluster.Members {
		members[i].ReferenceScore = article.ReferenceScore
		members[i].SubjectScore = subject.Score
	}
	cluster.Members = members
	cluster.AddMember(newMember)
}

func createNewClusterMember(c *domain.ArticleCluster, a news.Article, s news.Subject) domain.ClusterMember {
	return *domain.NewClusterMember(c.ClusterHash, a.ID, a.ReferenceScore, s.Score)
}

func createNewClusterMemebers(clusterHash string, article news.Article, subject news.Subject) []domain.ClusterMember {
	return []domain.ClusterMember{
		*domain.NewClusterMember(clusterHash, article.ID, article.ReferenceScore, subject.Score),
	}
}
