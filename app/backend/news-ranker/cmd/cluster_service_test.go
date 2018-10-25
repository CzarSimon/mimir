package main

import (
	"errors"
	"testing"
	"time"

	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/domain"
	"github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/repository"
	"github.com/CzarSimon/mimir/app/backend/pkg/mq"
	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
	"github.com/stretchr/testify/assert"
)

var mockError = errors.New("mock error")
var emptyCluster = domain.ArticleCluster{}

type mockClusterRepo struct {
	findByHashCluster domain.ArticleCluster
	findByHashErr     error
	findByHashArg     string
	saveReturn        error
	saveArg           domain.ArticleCluster
	updateReturn      error
	updateArg         domain.ArticleCluster
}

func (r *mockClusterRepo) FindByHash(arg string) (domain.ArticleCluster, error) {
	r.findByHashArg = arg
	return r.findByHashCluster, r.findByHashErr
}

func (r *mockClusterRepo) Save(arg domain.ArticleCluster) error {
	r.saveArg = arg
	return r.saveReturn
}

func (r *mockClusterRepo) Update(arg domain.ArticleCluster) error {
	r.updateArg = arg
	return r.updateReturn
}

func TestCreateNewCluster(t *testing.T) {
	assert := assert.New(t)

	articleDate, err := time.Parse("2006-01-02", "2018-10-25")
	assert.Nil(err)

	article := news.Article{
		ID:             "a-0",
		URL:            "http://url.com",
		Title:          "t-0",
		ReferenceScore: 0.5,
		ArticleDate:    articleDate,
	}
	subject := news.Subject{
		Symbol:    "smbl",
		Score:     0.3,
		ArticleID: "a-0",
	}
	clusterHash := domain.CalcClusterHash(article.Title, subject.Symbol, articleDate)

	clusterRepo := &mockClusterRepo{
		findByHashCluster: emptyCluster,
		findByHashErr:     repository.ErrNoSuchCluster,
		saveReturn:        nil,
	}
	mockEnv := newMockEnv(nil, clusterRepo, nil)

	mockEnv.clusterArticleWithSubject(article, subject)

	assert.Equal(clusterHash, clusterRepo.findByHashArg)
	cluster := clusterRepo.saveArg

	assert.Equal(clusterHash, cluster.ClusterHash)
	assert.Equal(article.ID, cluster.LeadArticleID)
	assert.Equal(article.Title, cluster.Title)
	assert.Equal(subject.Symbol, cluster.Symbol)
	assert.Equal(articleDate, cluster.ArticleDate)
	assert.Equal(0.8, cluster.Score)
	assert.Equal(1, len(cluster.Members))

	member := cluster.Members[0]
	assert.NotEqual("", member.ID)
	assert.Equal(article.ID, member.ArticleID)
	assert.Equal(subject.Score, member.SubjectScore)
	assert.Equal(article.ReferenceScore, member.ReferenceScore)

	clusterRepo = &mockClusterRepo{
		findByHashCluster: emptyCluster,
		findByHashErr:     repository.ErrNoSuchCluster,
		saveReturn:        mockError,
	}
	mockEnv = newMockEnv(nil, clusterRepo, nil)

	mockEnv.clusterArticleWithSubject(article, subject)

	assert.Equal(clusterHash, clusterRepo.findByHashArg)
	cluster = clusterRepo.saveArg
	assert.Equal(clusterHash, cluster.ClusterHash)

	clusterRepo = &mockClusterRepo{
		findByHashCluster: emptyCluster,
		findByHashErr:     mockError,
	}
	mockEnv = newMockEnv(nil, clusterRepo, nil)

	mockEnv.clusterArticleWithSubject(article, subject)

	assert.Equal(clusterHash, clusterRepo.findByHashArg)
	cluster = clusterRepo.saveArg
	assert.Equal("", cluster.ClusterHash)

}

func TestUpdateArticleCluster(t *testing.T) {

}

func newMockEnv(
	articleRepo repository.ArticleRepo,
	clusterRepo repository.ClusterRepo,
	mqClient mq.Client) *env {
	return &env{
		config: Config{
			TwitterUsers: 1000,
			MQ: MQConfig{
				Exchange:     "x-news",
				ScrapeQueue:  "q-scrape-targets",
				ScrapedQueue: "q-scraped-articles",
				RankQueue:    "q-rank-objects",
			},
		},
		articleRepo: articleRepo,
		clusterRepo: clusterRepo,
		mqClient:    mqClient,
	}
}
