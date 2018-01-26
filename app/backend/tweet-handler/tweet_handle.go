package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/CzarSimon/mimir/app/lib/go/schema/news"
	"github.com/CzarSimon/util"
)

// ReciveNewTweet Recives and handles new tweet
func (env *Env) ReciveNewTweet(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		util.SendErrStatus(res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	var tweet Tweet
	err := util.DecodeJSON(req.Body, &tweet)
	if err != nil {
		util.LogErr(err)
		util.SendErrStatus(res, errors.New("Could not parse tweet"), http.StatusMethodNotAllowed)
		return
	}
	go env.HandleNewTweet(tweet)
	util.SendOK(res)
}

// HandleNewTweet Checks if tweets if spam and handles it if not
func (env *Env) HandleNewTweet(tweet Tweet) {
	if isSpam := env.FilterSpam(tweet); isSpam {
		return
	}
	env.HandleNonSpamTweet(tweet)
}

// HandleNonSpamTweet Stores tweet, untracked tickers and sends urls to news-ranker
func (env *Env) HandleNonSpamTweet(tweet Tweet) {
	trackedSubjects, untrackedTickers := env.separateTrackedAndUntrackedTickers(tweet)
	env.storeTweet(tweet, trackedSubjects)
	env.SendLinksToRanker(tweet, trackedSubjects)
	env.storeUntrackedTickers(tweet, untrackedTickers)
}

// separateTrackedAndUntrackedTickers Separates tracked subjects and untracked ticker from a tweet
func (env *Env) separateTrackedAndUntrackedTickers(tweet Tweet) ([]news.Subject, []string) {
	tickers := env.mapSymbolListToTickers(tweet)
	trackedSubjects := env.mapTrackedTickersToSubjects(tickers)
	untrackedTickers := env.filterUntrackedTickers(tickers)
	return trackedSubjects, untrackedTickers

}

// mapSymbolList Maps list of symbols in a tweet to tickers, converting aliases to tickers
func (env *Env) mapSymbolListToTickers(tweet Tweet) []string {
	tickerSet := NewTickerSet()
	for _, ticker := range tweet.Entities.Symbols {
		if mappedAlias, isAnAlias := env.Aliases[ticker.Text]; isAnAlias {
			tickerSet.Add(mappedAlias)
		} else {
			tickerSet.Add(ticker.Text)
		}
	}
	return tickerSet.ToSlice()
}

// mapTrackedTickersToSubjects Filters tracked tickers and maps to Subjects
func (env *Env) mapTrackedTickersToSubjects(tickers []string) []news.Subject {
	trackedSubjects := make([]news.Subject, 0)
	for _, ticker := range tickers {
		if subject, isTracked := env.Tickers[ticker]; isTracked {
			trackedSubjects = append(trackedSubjects, subject)
		}
	}
	return trackedSubjects
}

// filterUntrackedTickers Filters supplied list of tickers to only inlcude untracked
func (env *Env) filterUntrackedTickers(tickers []string) []string {
	untrackedTickers := make([]string, 0)
	for _, ticker := range tickers {
		if _, isTracked := env.Tickers[ticker]; !isTracked {
			untrackedTickers = append(untrackedTickers, ticker)
		}
	}
	return untrackedTickers
}

// storeTweet Stores tweet in database
func (env *Env) storeTweet(tweet Tweet, subjects []news.Subject) {
	insertStmt, err := getStoreTweetStmt(env.DB)
	if err != nil {
		util.LogErr(err)
		return
	}
	defer insertStmt.Close()
	for _, subject := range subjects {
		err = insertTweet(tweet, subject, insertStmt)
		if err != nil {
			util.LogErr(err)
		}
	}
}

// insertTweet Inserts a tweet into a database through a supplied prepared statement
func insertTweet(tweet Tweet, subject news.Subject, insertStmt *sql.Stmt) error {
	_, err := insertStmt.Exec(
		tweet.ID,
		tweet.User.ID,
		tweet.GetDate(),
		tweet.Text,
		subject.Ticker,
		tweet.Language,
		tweet.User.Followers)
	return err
}

// getStoreTweetStmt Creates insertet statment for tweet
func getStoreTweetStmt(db *sql.DB) (*sql.Stmt, error) {
	return db.Prepare(`INSERT INTO STOCKTWEETS(
		tweetID, userID, createdAt, tweet, ticker, lang, followers
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
}

// storeUntrackedTickers Stores occurances of untracked tickers
func (env *Env) storeUntrackedTickers(tweet Tweet, untrackedTickers []string) {
	timestamp := tweet.GetDate()
	stmt, err := env.DB.Prepare(
		"INSERT INTO UNTRACKED_TICKERS(ID, TICKER, TIMESTAMP) VALUES ($1, $2, $3)")
	if err != nil {
		util.LogErr(err)
		return
	}
	defer stmt.Close()
	for _, ticker := range untrackedTickers {
		_, err = stmt.Exec(tweet.ID, ticker, timestamp)
		if err != nil {
			util.LogErr(err)
			return
		}
	}
}
