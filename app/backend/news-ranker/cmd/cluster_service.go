package main

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

func (e *env) clusterArticle(article news.Article) {
	fmt.Printf("Clustering article: %s\n", article)
}
