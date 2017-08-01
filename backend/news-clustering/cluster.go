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
		ClusterHash: member.ClusterHash,
		Title:       article.Title,
		Ticker:      article.Ticker,
		Date:        article.Date,
		Members:     []ClusterMember{member},
	}
}

// GetCluster Retrives cluster if exists and returns new cluser if not
func GetCluster(tx *sql.Tx, article Article, member ClusterMember) (Cluster, error) {
	var cluster Cluster
	query := `SELECT CLUSTER_HASH, LEADER, SCORE FROM ARTICLE_CLUSTER WHERE CLUSTER_HASH=$1`
	err := tx.QueryRow(query, member.ClusterHash).Scan(
		&cluster.ClusterHash, &cluster.Leader, &cluster.Score)
	if err == sql.ErrNoRows {
		newCluster := NewCluster(article, member)
		err = InsertNewCluster(tx, newCluster)
		if err != nil {
			return cluster, err
		}
		return newCluster, nil
	}
	if err != nil {
		return cluster, err
	}
	members, err := GetMembers(tx, member)
	if err != nil {
		return cluster, err
	}
	cluster.Members = members
	return cluster, nil
}

// InsertNewCluster Inserts Cluster hash and other metatdata for a new cluster
func InsertNewCluster(tx *sql.Tx, cluster Cluster) error {
	query := "INSERT INTO ARTICLE_CLUSTERS(CLUSTER_HASH, TITLE, TICKER, DATE) VALUES($1,$2,$3,$4)"
	stmt, err := tx.Prepare(query)
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

// ElectLeaderAndScore Elects a cluster leader and calculates the cluster score
func (cluster *Cluster) ElectLeaderAndScore() {
	leaderScore := 0.0
	leaderSubjectScore := 0.0
	clusterScore := 0.0
	newLeader := cluster.Leader
	var candidateScore float64
	for _, member := range cluster.Members {
		clusterScore += member.ReferenceScore
		candidateScore = CalcCompoundScore(member)
		if candidateScore > leaderScore {
			leaderScore = candidateScore
			leaderSubjectScore = member.SubjectScore
			newLeader = member.URLHash
		}
	}
	cluster.Score = leaderSubjectScore + clusterScore
	cluster.Leader = newLeader
}

// StoreClusterAndMember Stores updated cluster and new members
func StoreClusterAndMember(tx *sql.Tx, cluster Cluster, member ClusterMember) error {
	err := UpdateCluster(tx, cluster)
	if err != nil {
		return err
	}
	err = StoreMember(tx, member)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCluster Stores updated cluster
func UpdateCluster(tx *sql.Tx, cluster Cluster) error {
	query := "UPDATE ARTICLE_CLUSTER SET LEADER=$1, SCORE=$2 WHERE CLUSTER_HASH=$3"
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cluster.Leader, cluster.Score, cluster.ClusterHash)
	if err != nil {
		return err
	}
	return nil
}
