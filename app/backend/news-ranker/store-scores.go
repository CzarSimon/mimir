package main

import (
	"database/sql"
	"fmt"
	"log"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

func storeRankReturn(rankReturn RankReturn, db *sql.DB, clusterer endpoint.ServerAddr) {
	var err error
	if !rankReturn.StoredArticle.IsScraped {
		err = insertArticle(rankReturn, db)
	}
	if err != nil {
		log.Println(err)
		return
	}
	insertScores(rankReturn.NewArticle, db)
	go rankReturnToClustering(rankReturn, clusterer)
}

func insertArticle(rankReturn RankReturn, db *sql.DB) error {
	query := getArticleInsert()
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	article := rankReturn.StoredArticle
	rank := rankReturn.NewArticle
	return storeArticle(stmt, article, rank)
}

func getArticleInsert() string {
	return `INSERT INTO ARTICLE(
						URL, URL_HASH, TITLE, REFERENCE_SCORE, SUMMARY,
						BODY, DATE_INSERTED, TWITTER_REFERENCES, KEYWORDS)
						VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
}

func storeArticle(stmt *sql.Stmt, article Article, rank RankResult) error {
	_, err := stmt.Exec(
		string(article.URL),
		article.URLHash,
		rank.Title,
		article.ReferenceScore,
		rank.Summary,
		rank.Body,
		rank.Timestamp,
		pq.Array(article.TwitterReferences),
		pq.Array(rank.Keywords))
	return err
}

func insertScores(rankResult RankResult, db *sql.DB) {
	tx, err := db.Begin()
	if util.CheckErrAndRollback(err, tx) {
		return
	}
	err = insertScore("SUBJECT_SCORE", rankResult.SubjecScore, tx)
	if util.CheckErrAndRollback(err, tx) {
		return
	}
	err = insertScore("COMPOUND_SCORE", rankResult.CompundScore, tx)
	if util.CheckErrAndRollback(err, tx) {
		return
	}
	tx.Commit()
}

func insertScore(table string, subjects []TickerScore, tx *sql.Tx) error {
	query := fmt.Sprintf("INSERT INTO %s(URL_HASH, TICKER, SCORE) VALUES ($1,$2,$3)", table)
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	for _, subject := range subjects {
		_, err = stmt.Exec(subject.URLHash, subject.Ticker, subject.Score)
		if err != nil {
			fmt.Println(query)
			return err
		}
	}
	return nil
}

func updateScores(article Article, subjectScores []TickerScore, db *sql.DB) {
	tx, err := db.Begin()
	if util.CheckErrAndRollback(err, tx) {
		return
	}
	err = updateReferenceScore(article, tx)
	if util.CheckErrAndRollback(err, tx) {
		return
	}
	err = updateCompoundScores(article.ReferenceScore, subjectScores, tx)
	if util.CheckErrAndRollback(err, tx) {
		return
	}
	err = tx.Commit()
	util.CheckErr(err)
}

func updateReferenceScore(article Article, tx *sql.Tx) error {
	stmt, err := tx.Prepare("UPDATE ARTICLE SET REFERENCE_SCORE=$1, TWITTER_REFERENCES=$2 WHERE URL_HASH=$3")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(article.ReferenceScore, pq.Array(article.TwitterReferences), article.URLHash)
	if err != nil {
		return err
	}
	return nil
}

func updateCompoundScores(referenceScore float64, subjectScores []TickerScore, tx *sql.Tx) error {
	query := "UPDATE COMPOUND_SCORE SET SCORE=$1 WHERE URL_HASH=$2 AND TICKER=$3"
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	var compundScore float64
	for _, score := range subjectScores {
		compundScore = referenceScore + score.Score
		_, err = stmt.Exec(compundScore, score.URLHash, score.Ticker)
		if err != nil {
			return err
		}
	}
	return nil
}
