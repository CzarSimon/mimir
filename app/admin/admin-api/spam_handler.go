package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/CzarSimon/util"
)

// spamHandler Handles request related to spam candidates and labeleing resource
func (env *Env) spamHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return env.getSpamCandidates(w, r)
	case http.MethodPost:
		return env.labelSpam(w, r)
	default:
		return httputil.MethodNotAllowed
	}
}

// getSpamCandidates Sends a number of spam candidates to label
func (env *Env) getSpamCandidates(w http.ResponseWriter, r *http.Request) error {
	candidates, err := queryForSpamCandidates(env.TweetDB)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, candidates)
}

// queryForSpamCandidates Get spam candidates from database
func queryForSpamCandidates(db *sql.DB) ([]spam.Candidate, error) {
	rows, err := db.Query(getSpamCandidatesQuery(10))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return constructSpamCandidatesList(rows)
}

// getSpamCandidatesQuery constructs and returns the query for spam retrival.
func getSpamCandidatesQuery(limit int) string {
	return fmt.Sprintf(
		`SELECT DISTINCT T.TWEET
  		FROM STOCKTWEETS T
  		WHERE T.TWEET NOT IN (
    		SELECT S.TWEET FROM SPAM_DATA S
  		) LIMIT %d`, limit)
}

// constructSpamCandidatesList Structures a result set into a slice of spam candidates
func constructSpamCandidatesList(rows *sql.Rows) ([]spam.Candidate, error) {
	candidates := make([]spam.Candidate, 0, 10)
	var text string
	for rows.Next() {
		err := rows.Scan(&text)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, spam.NewCandidate(text))
	}
	return candidates, nil
}

// labelSpam Labels a text with with its spam class
func (env *Env) labelSpam(w http.ResponseWriter, r *http.Request) error {
	var candidate spam.Candidate
	err := util.DecodeJSON(r.Body, &candidate)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	err = storeSpamLabel(candidate, env.TweetDB)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendOK(w)
	return nil
}

// storeSpamLabel Stores the text and spam label of a supplied candidate
func storeSpamLabel(candidate spam.Candidate, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO SPAM_DATA (TWEET, LABEL) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(candidate.Text, candidate.Label)
	return err
}
