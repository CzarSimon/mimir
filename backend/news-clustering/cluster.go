package main

import (
	"database/sql"
  	"time"
)

// Cluster is the object representing the score and members of a cluster
type Cluster struct {
  	ClusterHash string
	Title       string
  	Ticker      string
 	Date        time.Time
	Leader      string         
	Score       float64
  	Members     []ClusterMember
}

// NewCluster Creates a new cluster based on an inital article and its member representation
func NewCluster(article Article, member ClusterMember) Cluster {
	return Cluster{
		ClusterHash: 	member.ClusterHash,
		Title:  	article.Title,
		Ticker: 	article.Ticker,
		Date: 		article.Date,
		Members: 	[]ClusterMember{member},
	}
}

// GetCluster Retrives cluster if exists and returns new cluser if not
func GetCluster(db *sql.DB, article Article, member ClusterMember) (Cluster, error) {
	var cluster Cluster
	query := `SELECT CLUSTER_HASH, LEADER, SCORE FROM ARTICLE_CLUSTER
		  WHERE CLUSTER_HASH=$1`
	err := db.QueryRow(query, member.ClusterHash).Scan(
		&cluster.ClusterHash, &cluster.Leader, &cluster.Score)
	if err == sql.ErrNoRows {
		newCluster := NewCluster(article, member)
		err = InsertNewCluster(db, newCluster)
		if err != nil {
			return cluster, err
		}
		return newCluster, nil
	}
	if err != nil {
		return cluster, err
	}
	members, err := GetMembers(db, member)
	if err != nil {
		return cluster, err
	}
	cluster.Members = members
	return cluster, nil
}

// InsertNewCluster Inserts Cluster hash and other metatdata for a new cluster
func InsertNewCluster(db *sql.DB, cluster Cluster) error {
	query := "INSERT INTO ARTICLE_CLUSTERS(CLUSTER_HASH, TITLE, TICKER, DATE) VALUES($1,$2,$3,$4)"
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cluster.ClusterHash, cluster.Title, cluster.Ticker, cluster.Date)
	if err != nil {
		return err
	}
	return nil	
}
