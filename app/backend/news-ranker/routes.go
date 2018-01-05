package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CzarSimon/util"
)

// HandleArticleRanking is the entrypoint for ranking and article
func (env *Env) HandleArticleRanking(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var rankObject RankObject
	err := decoder.Decode(&rankObject)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONStringRes(res, "Sent to ranking")
	env.rankArticles(rankObject)
}

// HandleRankedArticle stores the result of a ranked article
func (env *Env) HandleRankedArticle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling Ranked Article")
	var ranked RankReturn
	err := json.NewDecoder(req.Body).Decode(&ranked)
	if util.IsErr(err) {
		util.CheckErr(err)
		return
	}
	go storeRankReturn(ranked, env.db, env.clusterer)
	util.SendJSONStringRes(res, "success")
}
