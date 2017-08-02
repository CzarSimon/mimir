package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

//Article holds values for an article
type Article struct {
	Title             string   `json:"title"`
	Summary           string   `json:"summary"`
	URL               string   `json:"url"`
	Timestamp         string   `json:"timestamp"`
	Keywords          []string `json:"keywords"`
	TwitterReferences []int64  `json:"twitterReferences"`
}

// GetNews Retrives a ranked list of news based on parsed article params
func (env *Env) GetNews(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	params, err := parseArticleParams(ps, env.periodMonthMap)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	articles, err := getTopArticles(params, env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonBody, err := json.Marshal(articles)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// getTopArticles Gets the leading articles of the highest ranked article cluster
// for a given ticker and time preiod
func getTopArticles(params ArticleParams, db *sql.DB) ([]Article, error) {
	articles := make([]Article, 0)
	query := `SELECT a.TITLE, a.URL, a.SUMMARY, a.DATE_INSERTED, a.KEYWORDS, a.TWITTER_REFERENCES
						FROM ARTICLE_CLUSTER c
						INNER JOIN ARTICLE a ON c.LEADER = a.URL_HASH
						WHERE c.TICKER=$1 AND c.ARTICLE_DATE>=$2
						ORDER BY c.SCORE DESC LIMIT $3;`
	rows, err := db.Query(query, params.Ticker, params.FromDate, params.Limit)
	defer rows.Close()
	if err != nil {
		return articles, err
	}
	var article Article
	for rows.Next() {
		err = rows.Scan(
			&article.Title, &article.URL, &article.Summary, &article.Timestamp, &article.Keywords, &article.TwitterReferences)
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

// parseArticleParams Parses ArticleParams from a request
func parseArticleParams(ps httprouter.Params, periodMonthMap periodMonthMap) (ArticleParams, error) {
	ticker := strings.ToUpper(ps.ByName("ticker"))
	if ticker == "" {
		return ArticleParams{}, errors.New("missing ticker in request")
	}
	return ArticleParams{
		Ticker:     ticker,
		FromDate:   parseFromDate(ps, periodMonthMap),
		Limit:      parseLimit(ps),
		TimePeriod: parseTimePeriod(ps, periodMonthMap),
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
func parseTimePeriod(ps httprouter.Params, periodMonthMap periodMonthMap) string {
	period := strings.ToUpper(ps.ByName("period"))
	if _, present := periodMonthMap[period]; !present {
		return DefaultPeriod
	}
	return period
}

// parseFromDate Parses the starting date from a request
func parseFromDate(ps httprouter.Params, periodMonthMap periodMonthMap) time.Time {
	now := time.Now().UTC()
	period := parseTimePeriod(ps, periodMonthMap)
	months, found := periodMonthMap[period]
	if !found {
		return now
	}
	return now.AddDate(0, months, 0)
}

func printArticles(articles []Article) {
	for index, article := range articles {
		fmt.Println("---", index, "---")
		fmt.Println("Title:", article.Title)
		fmt.Println(article.Timestamp)
	}
}
