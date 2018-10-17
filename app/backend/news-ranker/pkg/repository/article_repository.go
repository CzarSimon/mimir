package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/CzarSimon/mimir/app/lib/go/schema/news"
)

const (
	keywordDelimiter = ","
)

var (
	ErrNoSuchArticle = errors.New("No such article")
)

type ArticleRepo interface {
	FindByURL(url string) (news.Article, error)
}

type pgArticleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) ArticleRepo {
	return &pgArticleRepo{
		db: db,
	}
}

const findByURLQuery = `SELECT
  id, url, title, body, keywords, reference_score, article_date, created_at
  FROM article WHERE url = $1`

func (r *pgArticleRepo) FindByURL(url string) (news.Article, error) {
	var a news.Article
	var joinedKeywords string
	err := r.db.QueryRow(findByURLQuery, url).Scan(
		&a.ID, &a.URL, &a.Title, &a.Body, joinedKeywords,
		&a.ReferenceScore, &a.ArticleDate, &a.CreatedAt)
	if err == sql.ErrNoRows {
		return a, ErrNoSuchArticle
	}
	if err != nil {
		return a, err
	}
	a.Keywords = splitKeywords(joinedKeywords)
	return a, nil
}

func joinKeywords(keywords []string) string {
	return strings.Join(keywords, keywordDelimiter)
}

func splitKeywords(joinedKeywords string) []string {
	return strings.Split(joinedKeywords, keywordDelimiter)
}
