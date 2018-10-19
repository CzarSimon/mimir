package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

const (
	keywordDelimiter = ","
)

var (
	ErrNoSuchArticle = errors.New("No such article")
	ErrNoSubjects    = errors.New("No subjects found")
)

type ArticleRepo interface {
	FindByURL(url string) (news.Article, error)
	FindArticleSubjects(articleID string) ([]news.Subject, error)
	FindArticleReferers(articleID string) ([]news.Referer, error)
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
	} else if err != nil {
		return a, err
	}
	a.Keywords = splitKeywords(joinedKeywords)
	return a, nil
}

const findArticleSubjectsQuery = `
  SELECT id, symbol, name, score, article_id FROM subject_score
  WHERE article_id = $1`

func (r *pgArticleRepo) FindArticleSubjects(articleID string) ([]news.Subject, error) {
	rows, err := r.db.Query(findArticleSubjectsQuery, articleID)
	if err == sql.ErrNoRows {
		return nil, ErrNoSubjects
	} else if err != nil {
		return nil, err
	}
	return mapRowsToSubjects(rows)
}

func mapRowsToSubjects(rows *sql.Rows) ([]news.Subject, error) {
	subjects := make([]news.Subject, 0)
	for rows.Next() {
		var s news.Subject
		err := rows.Scan(&s.ID, &s.Symbol, &s.Name, &s.Score, &s.ArticleID)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}

const findArticleReferersQuery = `
  SELECT id, twitter_author, follower_count, article_id FROM twitter_references
  WHERE article_id = $1`

func (r *pgArticleRepo) FindArticleReferers(articleID string) ([]news.Referer, error) {
	rows, err := r.db.Query(findArticleReferersQuery, articleID)
	if err == sql.ErrNoRows {
		return nil, ErrNoSubjects
	} else if err != nil {
		return nil, err
	}
	return mapRowsToReferers(rows)
}

func mapRowsToReferers(rows *sql.Rows) ([]news.Referer, error) {
	referers := make([]news.Referer, 0)
	for rows.Next() {
		var r news.Referer
		err := rows.Scan(&r.ID, &r.ExternalID, &r.FollowerCount)
		if err != nil {
			return nil, err
		}
		referers = append(referers, r)
	}
	return referers, nil
}

func joinKeywords(keywords []string) string {
	return strings.Join(keywords, keywordDelimiter)
}

func splitKeywords(joinedKeywords string) []string {
	return strings.Split(joinedKeywords, keywordDelimiter)
}
