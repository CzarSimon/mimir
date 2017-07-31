package main

import (
	"crypto/sha256"
	"fmt"
)

// Cluster is the object representing the score and members of a cluster
type Cluster struct {
	Title, Ticker, Date, Id string
	Members                 map[string]Member
	Leader                  Member
	Score                   float64
}

// Member represent an article in a cluester and it score
type Member struct {
	UrlHash string
	Score   Score
}

// Score holds the subject and reference score of a cluster member
type Score struct {
	SubjectScore   float64 `json:"subjectScore"`
	ReferenceScore float64 `json:"referenceScore"`
}

func newCluster(clusterHash string, article Article) Cluster {
	return Cluster{
		Title:   article.Title,
		Ticker:  article.Ticker,
		Date:    article.Date,
		Id:      clusterHash,
		Members: make(map[string]Member),
	}
}

func calcClusterHash(title, ticker, date string) string {
	byteHash := sha256.Sum256([]byte(title + ticker + date))
	return fmt.Sprintf("%x", byteHash)
}

func (cluster *Cluster) addMember(newMember Member) {
	cluster.Members[newMember.UrlHash] = newMember
}

func (cluster *Cluster) electLeader() {
	var leader Member
	for _, member := range cluster.Members {
		if leader != (Member{}) {
			if member.Score.calcCompoundScore() > leader.Score.calcCompoundScore() {
				leader = member
			}
		} else {
			leader = member
		}
	}
	cluster.Leader = leader
}

func (score Score) calcCompoundScore() float64 {
	return score.SubjectScore + score.ReferenceScore
}

func (cluster *Cluster) calculateScore() {
	if cluster.Leader != (Member{}) {
		clusterScore := Score{
			SubjectScore:   cluster.Leader.Score.SubjectScore,
			ReferenceScore: 0.0,
		}
		for _, member := range cluster.Members {
			clusterScore.ReferenceScore += member.Score.ReferenceScore
		}
		cluster.Score = clusterScore.calcCompoundScore()
	} else {
		cluster.Score = 0.0
	}
}
