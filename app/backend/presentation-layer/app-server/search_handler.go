package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/httputil"
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
func (env *Env) HandleUserSearch(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return env.HandleSearchHistoryRequest(w, r)
	case http.MethodPost:
		return env.HandleNewSearch(w, r)
	case http.MethodDelete:
		return env.HandleDeleteSearchHistoryRequest(w, r)
	default:
		return httputil.MethodNotAllowed
	}
}

// HandleNewSearch Stores a new search in the user search history
func (env *Env) HandleNewSearch(w http.ResponseWriter, r *http.Request) error {
	var search Search
	err := util.DecodeJSON(r.Body, &search)
	if err != nil || !search.IsPopulated() {
		return httputil.BadRequest
	}
	err = storeSearch(search, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendOK(w)
	return nil
}

// storeSearch Records a search in the supplied search history
func storeSearch(search Search, db *sql.DB) error {
	stmt, err := db.Prepare(`
    INSERT INTO SEARCH_HISTORY (
      USER_ID, SEARCH_TERM, DATE_INSERTED
    ) VALUES ($1, $2, CURRENT_TIMESTAMP)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(search.UserID, search.Query)
	return err
}

// HandleSearchHistoryRequest Retrives and sends a supplied users search history
func (env *Env) HandleSearchHistoryRequest(w http.ResponseWriter, r *http.Request) error {
	userID, err := query.ParseValue(r, USER_ID_KEY)
	if err != nil {
		return httputil.BadRequest
	}
	searchHistory, err := getSearchHistory(userID, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, searchHistory)
}

// getSearchHistory Retrives a users search history from the database
func getSearchHistory(userID string, db *sql.DB) ([]string, error) {
	rows, err := db.Query(getSeachHistoryQuery(), userID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	searchHistory := make([]string, 0)
	var search Search
	for rows.Next() {
		err = rows.Scan(&search.Query, &search.DateInserted)
		if err != nil {
			return nil, err
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
func (env *Env) HandleDeleteSearchHistoryRequest(w http.ResponseWriter, r *http.Request) error {
	userID, err := query.ParseValue(r, USER_ID_KEY)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	err = clearSearchHistory(userID, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendOK(w)
	return nil
}

// clearSearchHistory Removes the search history for a given user
func clearSearchHistory(userID string, db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM SEARCH_HISTORY WHERE USER_ID=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID)
	return err
}
