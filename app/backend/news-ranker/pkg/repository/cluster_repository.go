package repository

import "github.com/CzarSimon/mimir/app/backend/news-ranker/pkg/domain"

type ClusterRepository interface {
	FindByHash(clusterHash string) (domain.ArticleCluster, error)
}
