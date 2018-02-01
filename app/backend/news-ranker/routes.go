package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/handler"
	"github.com/CzarSimon/util"
)

func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/rank-article", handler.New(env.HandleArticleRanking))
	mux.Handle("/api/ranked-article", handler.New(env.HandleRankedArticle))
	mux.Handle("/health", handler.HealthCheck)
	mux.HandleFunc("/readiness", env.ReadinessCheck)
	return mux
}

// HandleArticleRanking is the entrypoint for ranking and article
func (env *Env) HandleArticleRanking(w http.ResponseWriter, r *http.Request) error {
	var rankObject RankObject
	err := util.DecodeJSON(r.Body, &rankObject)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	httputil.SendOK(w)
	env.rankArticles(rankObject)
	return nil
}

// HandleRankedArticle stores the result of a ranked article
func (env *Env) HandleRankedArticle(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Ranked Article")
	var ranked RankReturn
	err := util.DecodeJSON(r.Body, &ranked)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	go storeRankReturn(ranked, env.db, env.clusterer)
	httputil.SendOK(w)
	return nil
}

// ReadinessCheck checks that the server is ready to recieve traffic.
func (env *Env) ReadinessCheck(w http.ResponseWriter, r *http.Request) {
	err := env.db.Ping()
	if err != nil {
		httputil.SendErr(w, httputil.InternalServerError)
		return
	}
	httputil.SendOK(w)
}
