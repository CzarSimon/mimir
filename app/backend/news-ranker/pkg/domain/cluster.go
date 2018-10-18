package domain

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CzarSimon/mimir/app/backend/pkg/id"
)

const dateFormat = "2006-01-02"

// ArticleCluster is a collection of articles.
type ArticleCluster struct {
	ClusterHash   string
	Title         string
	Symbol        string
	ArticleDate   time.Time
	LeadArticleId string
	Score         float64
	Members       []ClusterMember
}

// AddMember add an additional member to the article cluster.
func (a *ArticleCluster) AddMember(newMember ClusterMember) {
	for _, member := range a.Members {
		if member.ArticleId == newMember.ArticleId {
			log.Printf("Member=[%s] already present in cluster=[%s]\n", newMember.String(), a.String())
			return
		}
	}
	a.Members = append(a.Members, newMember)
}

// ElectLeaderAndScore finds highes scoring member and sums up the total cluster score.
func (a *ArticleCluster) ElectLeaderAndScore() {
	leader := selectHighestScoreMember(a.Members)
	referenceSum := sumReferenceScore(a.Members)
	a.LeadArticleId = leader.ArticleId
	a.Score = leader.SubjectScore + referenceSum
}

func selectHighestScoreMember(members []ClusterMember) ClusterMember {
	var highScoreMember ClusterMember
	highScore := 0.0
	for _, member := range members {
		if member.Score() >= highScore {
			highScore = member.Score()
			highScoreMember = member
		}
	}
	return highScoreMember
}

func sumReferenceScore(members []ClusterMember) float64 {
	var referenceSum float64
	for _, member := range members {
		referenceSum += member.ReferenceScore
	}
	return referenceSum
}

func NewArticleCluster(title, symbol string, articleDate time.Time, leadArticleId string,
	score float64, members []ClusterMember) *ArticleCluster {
	return &ArticleCluster{
		ClusterHash:   CalcClusterHash(title, symbol, articleDate),
		Title:         title,
		Symbol:        symbol,
		ArticleDate:   articleDate,
		LeadArticleId: leadArticleId,
		Score:         score,
		Members:       members,
	}
}

// CalcClusterHash calculates sha256 digest of a title, symbol and date.
func CalcClusterHash(title, symbol string, date time.Time) string {
	lowerTitle := strings.ToLower(title)
	lowerSymbol := strings.ToLower(symbol)
	dateStr := date.Format(dateFormat)
	byteHash := sha256.Sum256([]byte(lowerTitle + lowerSymbol + dateStr))
	return fmt.Sprintf("%x", byteHash)
}

func (c *ArticleCluster) String() string {
	return fmt.Sprintf(
		"ArticleCluster(ClusterHash=%s Title=%s Symbol=%s ArticleDate=%s LeadArticleId=%s Score=%f)",
		c.ClusterHash, c.Title, c.Symbol, c.ArticleDate, c.LeadArticleId, c.Score)
}

// ClusterMember is a scored article that is part of a cluster.
type ClusterMember struct {
	Id             string
	ClusterHash    string
	ArticleId      string
	ReferenceScore float64
	SubjectScore   float64
}

func NewClusterMember(clusterHash, articleId string, referenceScore, subjectScore float64) *ClusterMember {
	return &ClusterMember{
		Id:             id.New(),
		ClusterHash:    clusterHash,
		ArticleId:      articleId,
		ReferenceScore: referenceScore,
		SubjectScore:   subjectScore,
	}
}

func (m *ClusterMember) Score() float64 {
	return m.ReferenceScore + m.SubjectScore
}

func (m *ClusterMember) String() string {
	return fmt.Sprintf(
		"ClusterMember(Id=%s ClusterHash=%s ArticleId=%s ReferenceScore=%f SubjectScore=%f)",
		m.Id, m.ClusterHash, m.ArticleId, m.ReferenceScore, m.SubjectScore)
}
