package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/util"
	"github.com/julienschmidt/httprouter"
)

// ArticleParams Holds values for determining what articles to query for
type ArticleParams struct {
	Ticker     string
	TimePeriod string
	FromDate   time.Time
	Limit      int
}

// GetNews Retrives a ranked list of news based on parsed article params
func (env *Env) GetNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params, err := parseArticleParams(ps, env.periodMonthMap)
	if err != nil {
		httputil.SendErr(w, httputil.BadRequest)
		return
	}
	articles, err := getTopArticles(params, env.db)
	if err != nil {
		log.Println(err)
		httputil.SendErr(w, httputil.InternalServerError)
		return
	}
	httputil.SendJSON(w, articles)
}

// getTopArticles Gets the leading articles of the highest ranked article cluster
// for a given ticker and time preiod
func getTopArticles(params ArticleParams, db *sql.DB) ([]Article, error) {
	query := getSelectArticlesQuery()
	rows, err := db.Query(
		query, params.Ticker, params.FromDate, time.Now().UTC(), params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return constructArticles(rows)
}

func constructArticles(rows *sql.Rows) ([]Article, error) {
	articles := make([]Article, 0)
	var article Article
	for rows.Next() {
		err := populateArticle(&article, rows)
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func populateArticle(article *Article, rows *sql.Rows) error {
	err := rows.Scan(
		&article.Title,
		&article.URL,
		&article.Summary,
		&article.Timestamp,
		&article.Keywords,
		&article.TwitterReferences)
	return err
}

func getSelectArticlesQuery() string {
	return `SELECT a.TITLE, a.URL, a.SUMMARY, a.DATE_INSERTED, a.KEYWORDS, a.TWITTER_REFERENCES
						FROM ARTICLE_CLUSTER c
						INNER JOIN ARTICLE a
						ON c.LEADER = a.URL_HASH
						WHERE c.TICKER=$1
						AND c.ARTICLE_DATE>=$2
						AND c.ARTICLE_DATE<=$3
						ORDER BY c.SCORE DESC LIMIT $4;`
}

// parseArticleParams Parses ArticleParams from a request
func parseArticleParams(ps httprouter.Params, periodMonthMap periodMonthMap) (ArticleParams, error) {
	ticker := strings.ToUpper(ps.ByName("ticker"))
	if ticker == "" {
		return ArticleParams{}, errors.New("missing ticker in request")
	}
	return ArticleParams{
		Ticker:     ticker,
		FromDate:   parseFromDate(ps),
		Limit:      parseLimit(ps),
		TimePeriod: parseTimePeriod(ps),
	}, nil
}

// parseLimit Parses the limit parameter from a request
func parseLimit(ps httprouter.Params) int {
	limit, err := strconv.Atoi(ps.ByName("top"))
	if err != nil {
		util.LogErr(err)
		return 5
	}
	return limit
}

// parseTimePeriod Parses the TimePeriod from a request
func parseTimePeriod(ps httprouter.Params) string {
	period := strings.ToUpper(ps.ByName("period"))
	return period
}

// parseFromDate Parses the starting date from a request
func parseFromDate(ps httprouter.Params) time.Time {
	period := parseTimePeriod(ps)
	adjustedDate := calcDateAdjustment(period)
	return adjustedDate
}

// FIXME: Move away from switch statement
// calcDateAdjustment Adjusts from date according to the priod supplied
func calcDateAdjustment(period string) time.Time {
	now := time.Now().UTC()
	switch period {
	case "TODAY":
		return now
	case "1W":
		return now.AddDate(0, 0, -7)
	case "1M":
		return now.AddDate(0, -1, 0)
	case "3M":
		return now.AddDate(0, -3, 0)
	default:
		return now
	}
}

func printArticles(articles []Article) {
	for index, article := range articles {
		fmt.Println("---", index, "---")
		fmt.Println("Title:", article.Title)
		fmt.Println(article.Timestamp)
	}
}
