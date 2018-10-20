package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/CzarSimon/mimir/app/backend/pkg/dbutil"
	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

const (
	keywordDelimiter = ","
)

var (
	ErrNoSuchArticle = errors.New("No such article")
	ErrNoSubjects    = errors.New("No subjects found")
	ErrNoReferers    = errors.New("No referers found")
	ErrFailedInsert  = errors.New("Insert failed")
)

type ArticleRepo interface {
	FindByURL(url string) (news.Article, error)
	FindArticleSubjects(articleID string) ([]news.Subject, error)
	FindArticleReferers(articleID string) ([]news.Referer, error)
	Update(article news.Article) error
	SaveReferer(referer news.Referer) error
}

type pgArticleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) ArticleRepo {
	return &pgArticleRepo{
		db: db,
	}
}

const findArticleByURLQuery = `SELECT
  id, url, title, body, keywords, reference_score, article_date, created_at
  FROM article WHERE url = $1`

func (r *pgArticleRepo) FindByURL(url string) (news.Article, error) {
	var a news.Article
	var joinedKeywords string
	err := r.db.QueryRow(findArticleByURLQuery, url).Scan(
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
	defer rows.Close()
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
		return nil, ErrNoReferers
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()
	return mapRowsToReferers(rows)
}

func mapRowsToReferers(rows *sql.Rows) ([]news.Referer, error) {
	referers := make([]news.Referer, 0)
	for rows.Next() {
		var r news.Referer
		err := rows.Scan(&r.ID, &r.ExternalID, &r.FollowerCount, &r.ArticleID)
		if err != nil {
			return nil, err
		}
		referers = append(referers, r)
	}
	return referers, nil
}

const updateArticleQuery = `
  UPDATE article SET reference_score = $1
  WHERE id = $2`

func (r *pgArticleRepo) Update(article news.Article) error {
	res, err := r.db.Exec(updateArticleQuery, article.ReferenceScore, article.ID)
	if err != nil {
		return err
	}

	var expectedUpdates int64 = 1
	return dbutil.AssertRowsAffected(res, expectedUpdates, ErrNoSuchArticle)
}

const insertReferencesQuery = `
  INSERT INTO twitter_references(id, twitter_author, follower_count, article_id)
  VALUES ($1, $2, $3, $4)`

func (r *pgArticleRepo) SaveReferer(referer news.Referer) error {
	if referer.ID == "" {
		referer.SetID()
	}

	res, err := r.db.Exec(
		insertReferencesQuery, referer.ID, referer.ExternalID,
		referer.FollowerCount, referer.ArticleID)
	if err != nil {
		return err
	}

	return dbutil.AssertRowsAffected(res, 1, ErrFailedInsert)
}

func joinKeywords(keywords []string) string {
	return strings.Join(keywords, keywordDelimiter)
}

func splitKeywords(joinedKeywords string) []string {
	return strings.Split(joinedKeywords, keywordDelimiter)
}
