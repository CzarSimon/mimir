package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/CzarSimon/util"
)

// spamHandler Handles request related to spam candidates and labeleing resource
func (env *Env) spamHandler(res http.ResponseWriter, req *http.Request) {
	status := http.StatusMethodNotAllowed
	err := METHOD_NOT_ALLOWED
	switch req.Method {
	case http.MethodGet:
		status, err = env.getSpamCandidates(res, req)
	case http.MethodPost:
		status, err = env.labelSpam(res, req)
	}
	if err != nil {
		util.SendErrStatus(res, err, status)
	}
}

// getSpamCandidates Sends a number of spam candidates to label
func (env *Env) getSpamCandidates(res http.ResponseWriter, req *http.Request) (int, error) {
	candidates, err := queryForSpamCandidates(env.TweetDB)
	if err != nil {
		log.Println(err)
		return InternalServerError()
	}
	err = SendJSON(res, candidates)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// queryForSpamCandidates Get spam candidates from database
func queryForSpamCandidates(db *sql.DB) ([]spam.Candidate, error) {
	rows, err := db.Query("SELECT TWEET FROM STOCKTWEETS LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return constructSpamCandidatesList(rows)
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
func (env *Env) labelSpam(res http.ResponseWriter, req *http.Request) (int, error) {
	var candidate spam.Candidate
	err := util.DecodeJSON(req.Body, &candidate)
	if err != nil {
		log.Println(err)
		return BadRequest()
	}
	err = storeSpamLabel(candidate, env.TweetDB)
	if err != nil {
		log.Println(err)
		return InternalServerError()
	}
	util.SendOK(res)
	return http.StatusOK, nil
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
