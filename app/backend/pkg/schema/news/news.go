package news

import (
	"fmt"
	"strings"
	"time"

	"github.com/CzarSimon/mimir/app/backend/pkg/id"
)

// Article contains article info.
type Article struct {
	ID             string    `json:"id"`
	URL            string    `json:"url"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Keywords       []string  `json:"keywords"`
	ReferenceScore float64   `json:"referenceScore"`
	ArticleDate    time.Time `json:"articleDate"`
	CreatedAt      time.Time `json:"createdAt"`
}

// NewArticle creates a new article.
func NewArticle(URL string) Article {
	return Article{
		ID:  id.New(),
		URL: URL,
	}
}

func (a Article) String() string {
	return fmt.Sprintf(
		"Article(id=%s url=%s title=%s keywords=[%s] referenceScore=%f)",
		a.ID, a.URL, a.Title, strings.Join(a.Keywords, ","), a.ReferenceScore)
}

// RankObject contains info to scrape and rank an article.
type RankObject struct {
	URLs     []string  `json:"urls"`
	Subjects []Subject `json:"subjects"`
	Author   Author    `json:"author"`
	Language string    `json:"language"`
}

func (ro RankObject) String() string {
	urls := strings.Join(ro.URLs, ",")
	subjects := joinSubjects(ro.Subjects)

	return fmt.Sprintf("RankObject(urls=[%s] subjects=[%s], author=%s, language=%s)",
		urls, subjects, ro.Author, ro.Language)
}

// Author contains info about an article refererer.
type Author struct {
	ID            string `json:"id"`
	FollowerCount int64  `json:"followerCount"`
}

func (a Author) String() string {
	return fmt.Sprintf("Author(id=%s followerCount=%d)", a.ID, a.FollowerCount)
}

// Subject subject which to look for in an article.
type Subject struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (s Subject) String() string {
	return fmt.Sprintf("Subject(name=%s ticker=%s)", s.Name, s.Symbol)
}

// ScrapeTarget article info needed to scrape and score an article.
type ScrapeTarget struct {
	URL            string    `json:"url"`
	Subjects       []Subject `json:"subjects"`
	ReferenceScore float64   `json:"referenceScore"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	ArticleID      string    `json:"articleId"`
}

func (s ScrapeTarget) String() string {
	return fmt.Sprintf(
		"ScrapeTarget(url=%s subjects=[%s] referenceScore=%f title=%s body=%s articleId=%s)",
		s.URL, joinSubjects(s.Subjects), s.ReferenceScore, s.Title, s.Body, s.ArticleID)
}

// ScrapedArticle result of scraping and scoring an article.
type ScrapedArticle struct {
	Article  Article        `json:"article"`
	Subjects []SubjectScore `json:"subjects"`
}

func (s ScrapedArticle) String() string {
	subjectList := make([]string, 0, len(s.Subjects))
	for _, subject := range s.Subjects {
		subjectList = append(subjectList, subject.String())
	}
	subjects := strings.Join(subjectList, ",")

	return fmt.Sprintf("ScrapedArticle(article=%s subjects=[%s])", s.Article.String(), subjects)
}

// SubjectScore scored subject in relation to an article.
type SubjectScore struct {
	ID        string  `json:"id"`
	Symbol    string  `json:"symbol"`
	Score     float64 `json:"score"`
	ArticleID string  `json:"articleId"`
}

func (s SubjectScore) String() string {
	return fmt.Sprintf(
		"SubjectScore(id=%s symbol=%s score=%f articleId=%s)",
		s.ID, s.Symbol, s.Score, s.ArticleID)
}

func joinSubjects(subjects []Subject) string {
	subjectList := make([]string, 0, len(subjects))
	for _, subject := range subjects {
		subjectList = append(subjectList, subject.String())
	}
	return strings.Join(subjectList, ",")
}
