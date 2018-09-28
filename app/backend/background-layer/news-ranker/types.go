package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/CzarSimon/util"
)

//Subject contains descriptive information about an articles subject
type Subject struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
}

//Author contains info about an article refererer.
type Author struct {
	ID            int64 `json:"id"`
	FollowerCount int64 `json:"followerCount"`
}

//RankObject contains info to scrape and rank an article
type RankObject struct {
	Urls     []URL     `json:"urls"`
	Subjects []Subject `json:"subjects"`
	Author   Author    `json:"author"`
	Language string    `json:"language"`
}

//Copy coipies a rank object
func (rankObject RankObject) Copy() RankObject {
	return RankObject{
		Urls:     rankObject.Urls,
		Subjects: rankObject.Subjects,
		Author:   rankObject.Author,
		Language: rankObject.Language,
	}
}

//NewTickerSet creates a ticker set based on a slice of scores
func NewTickerSet(scores []TickerScore) map[string]bool {
	tickerSet := make(map[string]bool)
	for _, score := range scores {
		tickerSet[score.Ticker] = true
	}
	return tickerSet
}

//Filter removes subjects already seen by the ranker and indicates if there are no sujects remaining
func (rankObject RankObject) Filter(subjectScores []TickerScore) (bool, RankObject) {
	filteredObject := rankObject.Copy()
	filteredObject.Subjects = make([]Subject, 0)
	tickerSet := NewTickerSet(subjectScores)
	for _, subject := range rankObject.Subjects {
		if _, present := tickerSet[subject.Ticker]; !present {
			filteredObject.Subjects = append(filteredObject.Subjects, subject)
		}
	}
	empty := len(filteredObject.Subjects) == 0
	return empty, filteredObject
}

//URL locates and article on the web
type URL string

//Hash turns a url into a shaq56 hash
func (url URL) Hash() string {
	byteHash := sha256.Sum256([]byte(url))
	return fmt.Sprintf("%x", byteHash)
}

//TwitterReferences is the list of twitter users that has refered an article
type TwitterReferences []int64

//Scan implements scanning a postgres array to TwitterReferences
func (tr *TwitterReferences) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	intSlice, err := util.BytesToIntSlice(bytes, 64)
	if util.IsErr(err) {
		return err
	}
	(*tr) = intSlice
	return nil
}

//Keywords is the list of keywords computed for an article
type Keywords []string

//Scan implements scanning a postgres arrya into a string slice
func (keywords *Keywords) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	strSlice := util.BytesToStrSlice(bytes)
	(*keywords) = strSlice
	return nil
}

//Article contains article info
type Article struct {
	URLHash           string            `json:"urlHash"`
	URL               URL               `json:"url"`
	Title             string            `json:"title"`
	ReferenceScore    float64           `json:"referenceScore"`
	Summary           string            `json:"summary"`
	Body              string            `json:"body"`
	DateInserted      time.Time         `json:"dateInserted"`
	TwitterReferences TwitterReferences `json:"twitterReferences"`
	Keywords          []string          `json:"keywords"`
	IsScraped         bool              `json:"isScraped"`
}

//UpdateReferenceScore updates the article reference score
func (article *Article) UpdateReferenceScore(author Author, config RankConfig) {
	additonalScore := config.CalcReferenceScore(author)
	article.ReferenceScore = article.ReferenceScore + additonalScore
}

//New creates a new article with URL, URLHash, ReferenceScore and TwitterReferences initalized
func (article *Article) New(url URL, rankObject RankObject, config RankConfig) {
	if article.URLHash == "" {
		article.URLHash = url.Hash()
	}
	article.IsScraped = false
	article.URL = url
	article.ReferenceScore = config.CalcReferenceScore(rankObject.Author)
	article.TwitterReferences = []int64{rankObject.Author.ID}
}

//HasNewReference checks if author has already referenced the article
func (article *Article) HasNewReference(authorID int64) bool {
	for _, reference := range article.TwitterReferences {
		if reference == authorID {
			return false
		}
	}
	return true
}

//AddReference adds a new author to the list of TwitterReferences
func (article *Article) AddReference(authorID int64) {
	article.TwitterReferences = append(article.TwitterReferences, authorID)
}

//Update adds a new article referer and updates the reference score
func (article *Article) Update(author Author, config RankConfig) {
	article.AddReference(author.ID)
	article.UpdateReferenceScore(author, config)
}

//RankArgument is the data sent to the python ranker
type RankArgument struct {
	ArticleInfo    RankObject `json:"articleInfo"`
	URL            URL        `json:"url"`
	ReferenceScore float64    `json:"referenceScore"`
	StoredArticle  Article    `json:"storedArticle"`
}

//Print outputs the contents of a rank argument to standard out
func (arg RankArgument) Print() {
	fmt.Println("articleInfo", arg.ArticleInfo)
	fmt.Println("url", arg.URL)
	fmt.Println("referenceScore", arg.ReferenceScore)
	fmt.Println("storedArticle", arg.StoredArticle)
}

//NewRankArgument creates a ranking argument object based on a RankObject and Article
func NewRankArgument(rankObject RankObject, article Article) RankArgument {
	return RankArgument{
		ArticleInfo:    rankObject,
		URL:            article.URL,
		ReferenceScore: article.ReferenceScore,
		StoredArticle:  article,
}

//ToString turns RankArgument to a json string
func (arg RankArgument) ToString() (string, error) {
	js, err := json.Marshal(arg)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

//TickerScore holds the ticker specific score of an article
type TickerScore struct {
	URLHash string  `json:"urlHash"`
	Ticker  string  `json:"ticker"`
	Score   float64 `json:"score"`
}

//CreateCompoundScore turns a subject score to a compund score
func (score TickerScore) CreateCompoundScore(referenceScore float64) TickerScore {
	return TickerScore{
		URLHash: score.URLHash,
		Ticker:  score.Ticker,
		Score:   score.Score + referenceScore,
	}
}

//IsTicker checks if the score is in reference to the supplied ticker
func (score TickerScore) IsTicker(ticker string) bool {
	return ticker == score.Ticker
}

//RankResult is the resulting score and content retrived and calulated by scrape_and_rank
type RankResult struct {
	Title        string        `json:"title"`
	Summary      string        `json:"summary"`
	Body         string        `json:"body"`
	Keywords     []string      `json:"keywords"`
	SubjecScore  []TickerScore `json:"subjectScore"`
	CompundScore []TickerScore `json:"compoundScore"`
	Timestamp    time.Time     `json:"timestamp"`
}

//ToArticle turns a RankResult to an article shell
func (result RankResult) ToArticle(URLHash string, referenceScore float64) Article {
	return Article{
		Title:          result.Title,
		Summary:        result.Summary,
		Body:           result.Body,
		URLHash:        URLHash,
		DateInserted:   result.Timestamp,
		ReferenceScore: referenceScore,
	}
}

//RankReturn is the complete result sent from scrape_and_rank
type RankReturn struct {
	NewArticle    RankResult `json:"newArticle"`
	StoredArticle Article    `json:"storedArticle"`
}

//ClusterArticle holds the cluster representation of an article
type ClusterArticle struct {
	Title   string          `json:"title"`
	Ticker  string          `json:"ticker"`
	Date    time.Time       `json:"date"`
	URLHash string          `json:"urlHash"`
	Score   ClusteringScore `json:"score"`
}

//Print outputs the contents of a cluster article to stdout
func (article ClusterArticle) Print() {
	fmt.Println("Title:", article.Title)
	fmt.Println("Ticker:", article.Ticker)
	fmt.Println("Date:", article.Date)
	fmt.Println("URLHash:", article.URLHash)
	article.Score.Print()
}

//NewClusterArticle creates a cluster artilce shell based on a artile
func NewClusterArticle(article Article) ClusterArticle {
	return ClusterArticle{
		Title:   article.Title,
		Date:    article.DateInserted,
		URLHash: article.URLHash,
		Score: ClusteringScore{
			ReferenceScore: article.ReferenceScore,
		},
	}
}

//ClusteringScore holds the cluster representation of an article socre
type ClusteringScore struct {
	SubjectScore   float64 `json:"subjectScore"`
	ReferenceScore float64 `json:"referenceScore"`
}

//Print outputs the content of score to stdout
func (score ClusteringScore) Print() {
	fmt.Println("SubjectScore:", score.SubjectScore)
	fmt.Println("ReferenceScore:", score.ReferenceScore)
}
