package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/CzarSimon/httputil/query"
	"github.com/CzarSimon/util"
)

// Search Holds search info
type Search struct {
	UserID       string    `json:"id"`
	Query        string    `json:"query"`
	DateInserted time.Time `json:"dateInserted"`
}

// IsPopulated Checks if a search is successfully populated
func (search Search) IsPopulated() bool {
	return search.UserID != "" && search.Query != ""
}

// HandleUserSearch Handles requests regarding a users search history
func (env *Env) HandleUserSearch(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		env.HandleSearchHistoryRequest(res, req)
	case http.MethodPost:
		env.HandleNewSearch(res, req)
	case http.MethodDelete:
		env.HandleDeleteSearchHistoryRequest(res, req)
	default:
		util.SendErrStatus(
			res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
	}
}

// HandleNewSearch Stores a new search in the user search history
func (env *Env) HandleNewSearch(res http.ResponseWriter, req *http.Request) {
	var search Search
	err := util.DecodeJSON(req.Body, &search)
	if err != nil || !search.IsPopulated() {
		util.SendErrStatus(
			res, errors.New("Could not parse search"), http.StatusBadRequest)
		return
	}
	err = storeSearch(search, env.db)
	if err != nil {
		util.SendErrRes(res, errors.New("Could not record the search"))
		return
	}
	util.SendOK(res)
}

// storeSearch Records a search in the supplied search history
func storeSearch(search Search, db *sql.DB) error {
	stmt, err := db.Prepare(`
    INSERT INTO SEARCH_HISTORY (
      USER_ID, SEARCH_TERM, DATE_INSERTED
    ) VALUES ($1, $2, CURRENT_TIMESTAMP)`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(search.UserID, search.Query)
	if err != nil {
		return err
	}
	return nil
}

// HandleSearchHistoryRequest Retrives and sends a supplied users search history
func (env *Env) HandleSearchHistoryRequest(res http.ResponseWriter, req *http.Request) {
	userID, err := query.ParseValue(req, USER_ID_KEY)
	if err != nil {
		util.SendErrStatus(res, err, http.StatusBadRequest)
		return
	}
	searchHistory, err := getSearchHistory(userID, env.db)
	if err != nil {
		util.LogErr(err)
		util.SendErrRes(res, errors.New("Unable to get users search history"))
		return
	}
	jsonBody, err := json.Marshal(searchHistory)
	if err != nil {
		util.SendErrRes(res, errors.New("Unable to get users search history"))
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// getSearchHistory Retrives a users search history from the database
func getSearchHistory(userID string, db *sql.DB) ([]string, error) {
	searchHistory := make([]string, 0)
	rows, err := db.Query(getSeachHistoryQuery(), userID)
	defer rows.Close()
	if err != nil {
		return searchHistory, err
	}
	var search Search
	for rows.Next() {
		err = rows.Scan(&search.Query, &search.DateInserted)
		if err != nil {
			return searchHistory, err
		}
		searchHistory = append(searchHistory, search.Query)
	}
	return searchHistory, nil
}

// getSeachHistoryQuery Returns the query for getting distinct search terms for a user
func getSeachHistoryQuery() string {
	return `SELECT SEARCH_TERM, MAX(DATE_INSERTED) AS DATE_INSERTED
            FROM SEARCH_HISTORY WHERE USER_ID=$1
            GROUP BY SEARCH_TERM
            ORDER BY DATE_INSERTED DESC`
}

// HandleDeleteSearchHistoryRequest Handles the deletrion of a users search history
func (env *Env) HandleDeleteSearchHistoryRequest(res http.ResponseWriter, req *http.Request) {
	userID, err := query.ParseValue(req, USER_ID_KEY)
	if err != nil {
		util.SendErrStatus(res, err, http.StatusBadRequest)
		return
	}
	err = clearSearchHistory(userID, env.db)
	if err != nil {
		util.SendErrRes(res, errors.New("Could not clear search history"))
		return
	}
	util.SendOK(res)
}

// clearSearchHistory Removes the search history for a given user
func clearSearchHistory(userID string, db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM SEARCH_HISTORY WHERE USER_ID=$1")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}
	return nil
}
