package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/CzarSimon/util"
)

func sendToClustering(article Article, subjectScores []TickerScore, clusterer util.ServerConfig) {
	fmt.Println("Sending to clusterer")
	clusterArticles := createClusterArticles(article, subjectScores)
	clustererURL := clusterer.ToURL("api/cluster-article")
	for _, clusterArticle := range clusterArticles {
		err := sendArticleToClustering(clusterArticle, clustererURL)
		if util.IsErr(err) {
			util.LogErr(err)
			return
		}
	}
}

func rankReturnToClustering(rankReturn RankReturn, clusterer util.ServerConfig) {
	clusterArticle := rankReturn.NewArticle.ToArticle(
		rankReturn.StoredArticle.URLHash, rankReturn.StoredArticle.ReferenceScore)
	sendToClustering(clusterArticle, rankReturn.NewArticle.SubjecScore, clusterer)
}

func sendArticleToClustering(clusterArticle ClusterArticle, endpoint string) error {
	jsonStr, err := json.Marshal(clusterArticle)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("Non 200 response")
}

func createClusterArticles(article Article, subjectScores []TickerScore) []ClusterArticle {
	clusterArticles := make([]ClusterArticle, 0)
	clusterArticle := NewClusterArticle(article)
	for _, subject := range subjectScores {
		clusterArticle.Score.SubjectScore = subject.Score
		clusterArticle.Ticker = subject.Ticker
		clusterArticles = append(clusterArticles, clusterArticle)
	}
	return clusterArticles
}
