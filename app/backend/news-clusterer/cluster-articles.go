package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/util"
)

// ParseArticleAndCluster Handles a cluster request. Pararses before passing on to cluster handler
func (env *Env) ParseArticleAndCluster(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return httputil.MethodNotAllowed
	}
	var article Article
	err := util.DecodeJSON(r.Body, &article)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	httputil.SendOK(w)
	env.HandleClustering(article)
	return nil
}

// ClusterArticle Clusters an article and stores cluster and member
func ClusterArticle(tx *sql.Tx, article Article, member ClusterMember) error {
	cluster, err := GetCluster(tx, article, member)
	if err != nil {
		return err
	}
	cluster.ElectLeaderAndScore()
	err = StoreClusterAndMember(tx, cluster, member)
	if err != nil {
		return nil
	}
	return nil
}

// HandleClustering Locks the cluster while an update is carried out
func (env *Env) HandleClustering(article Article) {
	member := article.ToClusterMember()
	env.AddAndLockCluster(member.ClusterHash)
	tx, err := env.db.Begin()
	if err != nil {
		util.CheckErrAndRollback(err, tx)
		env.RemoveAndUnlockCluster(member.ClusterHash)
		return
	}
	err = ClusterArticle(tx, article, member)
	if err != nil {
		util.CheckErrAndRollback(err, tx)
		env.RemoveAndUnlockCluster(member.ClusterHash)
		return
	}
	err = tx.Commit()
	if err != nil {
		util.CheckErrAndRollback(err, tx)
	}
	env.RemoveAndUnlockCluster(member.ClusterHash)
}
