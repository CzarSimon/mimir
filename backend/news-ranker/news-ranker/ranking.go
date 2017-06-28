package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os/exec"

	"github.com/CzarSimon/util"
)

func (env *Env) rankArticles(rankObject RankObject) {
	for _, url := range rankObject.Urls {
		env.rankArticle(url, rankObject)
	}
}

func (env *Env) rankArticle(url URL, rankObject RankObject) {
	articleFound, article := checkForArticle(url, env.db)
	if articleFound {
		fmt.Println("Existing article")
		env.handleExistingArticle(article, rankObject)
	} else {
		article.New(url, rankObject, env.rank)
		fmt.Println("New article")
		go sendToRanker(article, rankObject, env.rank)
	}
}

func (env *Env) handleExistingArticle(article Article, rankObject RankObject) {
	subjectScores, err := checkForSubjects(article.URLHash, env.db)
	if err != nil {
		util.CheckErr(err)
		return
	}
	hasNewReference := article.HasNewReference(rankObject.Author.ID)
	if hasNewReference {
		fmt.Println("New reference")
		article.Update(rankObject.Author, env.rank)
		go updateScores(article, subjectScores, env.db)
	}
	noNew, filteredRankObject := rankObject.Filter(subjectScores)
	if !noNew {
		fmt.Println("New refernece sending to ranking")
		go sendToRanker(article, filteredRankObject, env.rank)
		go sendToClustering(article, subjectScores, env.clusterer)
	}
}

func checkForArticle(url URL, db *sql.DB) (bool, Article) {
	article := Article{
		URLHash:   url.Hash(),
		IsScraped: true,
	}
	query := "SELECT URL, TITLE, BODY, DATE_INSERTED, TWITTER_REFERENCES, REFERENCE_SCORE FROM ARTICLE WHERE URL_HASH=$1"
	err := db.QueryRow(query, article.URLHash).Scan(
		&article.URL, &article.Title, &article.Body, &article.DateInserted, &article.TwitterReferences, &article.ReferenceScore)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		article.IsScraped = false
		return false, article
	}
	return true, article
}

func checkForSubjects(urlHash string, db *sql.DB) ([]TickerScore, error) {
	subjectScores := make([]TickerScore, 0)
	stmt, err := db.Prepare("SELECT TICKER, SCORE FROM SUBJECT_SCORE WHERE URL_HASH=$1")
	defer stmt.Close()
	if err != nil {
		return subjectScores, err
	}
	rows, err := stmt.Query(urlHash)
	defer rows.Close()
	if err != nil {
		return subjectScores, err
	}
	var score TickerScore
	score.URLHash = urlHash
	for rows.Next() {
		err = rows.Scan(&score.Ticker, &score.Score)
		if err != nil {
			return subjectScores, err
		}
		subjectScores = append(subjectScores, score)
	}
	return subjectScores, nil
}

func sendToRanker(article Article, rankObject RankObject, config RankConfig) {
	rankArg, err := NewRankArgument(rankObject, article).ToString()
	if err != nil {
		log.Println(err.Error())
		return
	}
	cmd := exec.Command(config.Command, config.Path+config.Script, rankArg)
	var out bytes.Buffer
	cmd.Stdout = &out
	var errOut bytes.Buffer
	cmd.Stderr = &errOut
	err = cmd.Run()
	util.CheckErr(err)
	printOutStream(out.String())
	printOutStream(errOut.String())
}

func printOutStream(stream string) {
	if stream != "" {
		fmt.Println(stream)
	}
}
