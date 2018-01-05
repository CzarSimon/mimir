package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type labeledTweet struct {
	Tweet, Label string
}

func (env *Env) getSpamCandidates(res http.ResponseWriter, req *http.Request) {
	if env.authenticate(res, req) != nil {
		return
	}
	spamCandidates, err := querySpamCandidates(env.pg)
	if err != nil {
		sendError(res, err)
		return
	}
	js, err := json.Marshal(spamCandidates)
	if err != nil {
		sendError(res, err)
		return
	}
	jsonRes(res, js)
}

func querySpamCandidates(pg *sql.DB) ([]string, error) {
	candidates := make([]string, 0)
	query := "SELECT tweet FROM stocktweets WHERE createdat::date=$1::date limit 100"
	rows, err := pg.Query(query, getRandomDate())
	defer rows.Close()
	if err != nil {
		return candidates, err
	}
	var candidate string
	for rows.Next() {
		err := rows.Scan(&candidate)
		if err != nil {
			checkErr(err)
		} else {
			candidates = append(candidates, candidate)
		}
	}
	return candidates, nil
}

func (env *Env) labelTweet(res http.ResponseWriter, req *http.Request) {
	if env.authenticate(res, req) != nil {
		return
	}
	decoder := json.NewDecoder(req.Body)
	var tweet labeledTweet
	err := decoder.Decode(&tweet)
	if err != nil {
		sendError(res, err)
		return
	}
	defer req.Body.Close()
	err = insertTweet(tweet, env.pg)
	if err != nil {
		sendError(res, err)
		return
	}
	jsonStringRes(res, "Successfully labeled tweet")
}

func insertTweet(tweet labeledTweet, pg *sql.DB) error {
	stmt, err := pg.Prepare("INSERT INTO SPAM_DATA (TWEET, LABEL) VALUES ($1, $2)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tweet.Tweet, tweet.Label)
	return err
}

func getRandomDate() time.Time {
	min, max, err := minMaxDates()
	if err != nil {
		checkErr(err)
		return time.Now().UTC()
	}
	dateRange := max - min
	randomUnixDate := rand.Int63n(dateRange) + min
	randomDate := time.Unix(randomUnixDate, 0)
	log.Println("Random date:", randomDate)
	return randomDate
}

func minMaxDates() (int64, int64, error) {
	dateFormat := "2006-01-02"
	min, err := time.Parse(dateFormat, "2016-05-30")
	if err != nil {
		return 0, 0, err
	}
	max := time.Now().UTC()
	return min.Unix(), max.Unix(), nil
}
