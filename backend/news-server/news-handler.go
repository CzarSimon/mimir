package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	r "gopkg.in/gorethink/gorethink.v2"
)

type articleParams struct {
	ticker, date string
	limit        int
}

//Article holds values for an article
type Article struct {
	Title, Summary, URL, Timestamp string
	Compound_Score                 map[string]float64
	Keywords                       []string
	Twitter_References             []int64
}

func (env *Env) getNews(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	limit, err := strconv.Atoi(ps.ByName("top"))
	if err != nil {
		checkErr(err)
		limit = 5
	}
	params := articleParams{
		ticker: ps.ByName("ticker"),
		limit:  limit,
		date:   getDate(),
	}
	articles := getTopArticles(params, env.db)
	js, err := json.Marshal(articles)
	if err != nil {
		errorResponse(res)
		return
	}
	jsonRes(res, js)
}

func getTopArticles(params articleParams, session *r.Session) []Article {
	articles := make([]Article, 0)
	rows, err := r.Table("articles").GetAll(r.Args(r.Table("article_clusters").GetAllByIndex("date", params.date).Filter(map[string]interface{}{
		"ticker": params.ticker,
	}).Filter(func(cluster r.Term) r.Term {
		return cluster.Field("leader").Field("Score").Field("SubjectScore").Gt(0)
	}).OrderBy(r.Desc("score")).Limit(params.limit).Field("leader").Field("UrlHash"))).Pluck("title", "url", "keywords", "twitter_references", "compound_score", "timestamp", "summary").Run(session)
	checkErr(err)
	defer rows.Close()
	var article Article
	for rows.Next(&article) {
		articles = append(articles, article)
	}
	return articles
}

func printArticles(articles []Article) {
	for index, article := range articles {
		fmt.Println("---", index, "---")
		fmt.Println("Title:", article.Title)
		for ticker, score := range article.Compound_Score {
			fmt.Println(ticker, ":", score)
		}
		fmt.Println(article.Timestamp)
	}
}

type ArticleCluster struct {
	Title, ID string
	Score     float64
}

func getClusters(params articleParams, session *r.Session) {
	clusters := make([]ArticleCluster, 0)
	rows, err := r.Table("article_clusters").GetAllByIndex("date", params.date).Filter(map[string]interface{}{
		"ticker": params.ticker,
	}).Filter(func(cluster r.Term) r.Term {
		return cluster.Field("leader").Field("Score").Field("SubjectScore").Gt(0)
	}).OrderBy(r.Desc("score")).Limit(params.limit).Run(session)
	checkErr(err)
	defer rows.Close()

	var cluster ArticleCluster
	for rows.Next(&cluster) {
		clusters = append(clusters, cluster)
	}
	for _, clust := range clusters {
		fmt.Println(clust)
	}
}
