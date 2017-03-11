package main

import (
  "encoding/json"
  "errors"
  "net/http"
  r "gopkg.in/gorethink/gorethink.v2"
)

/* --- Handles a cluster request. Pararses before passing on to cluster handler --- */
func (env *Env) clusterArticle(res http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var article Article
  if err := decoder.Decode(&article); err != nil {
    checkErr(err)
    jsonStringRes(res, "Article parsing failed")
    return
  }
  article.Format()
  clusterHash := calcClusterHash(article.Title, article.Ticker, article.Date)
  jsonStringRes(res, "Article sent for clustering")
  env.lockCluster(clusterHash, article)
}

/* --- Handles the update or creation of a cluster --- */
func handleClustering(clusterHash string, article Article, session *r.Session) {
  cluster, err := lookupCluster(clusterHash, session)
  if err == nil {
    cluster = updateCluster(cluster, article)
    err = storeUpdatedCluster(cluster, session)
    checkErr(err)
  } else {
    cluster = newCluster(clusterHash, article)
    cluster = updateCluster(cluster, article)
    err = storeNewCluster(cluster, session)
    checkErr(err)
  }
}

/* --- Retrives the cluster from the database, returns nil if not found --- */
func lookupCluster(clusterHash string, session *r.Session) (Cluster, error) {
  var cluster Cluster
  res, err := r.Table("article_clusters").GetAll(clusterHash).Run(session)
  if err != nil {
    checkErr(nil)
    return Cluster{}, errors.New("no such cluster")
  }
  defer res.Close()
  if res.IsNil() {
    return Cluster{}, errors.New("no such cluster")
  }
  err = res.One(&cluster)
  if err != nil {
    checkErr(nil)
    return Cluster{}, errors.New("no such cluster")
  }
  return cluster, nil
}

/* --- Updates cluster in db --- */
func storeUpdatedCluster(cluster Cluster, session *r.Session) error {
  _, err := r.Table("article_clusters").GetAll(cluster.Id).Update(map[string]interface{}{
    "members": cluster.Members,
    "leader": cluster.Leader,
    "score": cluster.Score,
  }).RunWrite(session)
  return err
}

/* --- Adds new cluster to db --- */
func storeNewCluster(cluster Cluster, session *r.Session) error {
  _, err := r.Table("article_clusters").Insert(map[string]interface{}{
    "id": cluster.Id,
    "title": cluster.Title,
    "ticker": cluster.Ticker,
    "date": cluster.Date,
    "members": cluster.Members,
    "leader": cluster.Leader,
    "score": cluster.Score,
  }).RunWrite(session)
  return err
}

/* --- Updates and existing cluster with a new article --- */
func updateCluster(cluster Cluster, article Article) Cluster {
  newMember := Member{
    UrlHash: article.UrlHash,
    Score: article.Score,
  }
  cluster.addMember(newMember)
  cluster.electLeader()
  cluster.calculateScore()
  return cluster
}

/* --- Locks the cluster while an update is carried out --- */
func (env *Env) lockCluster(clusterHash string, article Article) {
  env.queue.addCluster(clusterHash)
  env.queue.lockCluster(clusterHash)

  handleClustering(clusterHash, article, env.db)

  env.queue.unlockCluster(clusterHash)
  env.queue.removeCluster(clusterHash)
}
