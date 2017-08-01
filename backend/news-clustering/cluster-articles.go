package main

import (
	"database/sql"
	"net/http"

	"github.com/CzarSimon/util"
)

// ParseArticleAndCluster Handles a cluster request. Pararses before passing on to cluster handler
func (env *Env) ParseArticleAndCluster(res http.ResponseWriter, req *http.Request) {
	var article Article
	err := util.DecodeJSON(req.Body, &article)
	if err != nil {
		util.PrintErr(err)
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONStringRes(res, "Article sent for clustering")
	env.HandleClustering(article)
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
