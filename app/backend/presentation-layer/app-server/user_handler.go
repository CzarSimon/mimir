package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/query"
	"github.com/CzarSimon/mimir/app/lib/go/schema"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

const USER_ID_KEY = "id"

// HandleUserRequest Handles retrival of user or creation of new user
func (env *Env) HandleUserRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.MethodNotAllowed
	}
	userID, err := query.ParseValue(r, USER_ID_KEY)
	log.Println(userID)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	user, err := getUser(userID, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	err = storeUserSession(schema.NewSession(user), env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, user)
}

// getUser Retrives user from datababase if exist, returns new user if not
func getUser(userID string, db *sql.DB) (schema.User, error) {
	var user schema.User
	err := db.QueryRow("SELECT ID, TICKERS, JOIN_DATE FROM APP_USER WHERE ID=$1", userID).Scan(
		&user.ID, &user.Tickers, &user.JoinDate)
	if err == sql.ErrNoRows {
		newUser := schema.NewUser(userID)
		err = StoreNewUser(newUser, db)
		return newUser, err
	}
	return user, err
}

// StoreNewUser Stores a newly created user in the supplied database
func StoreNewUser(user schema.User, db *sql.DB) error {
	stmt, err := db.Prepare(
		"INSERT INTO APP_USER(ID, TICKERS, JOIN_DATE) VALUES ($1, $2, $3)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID, pq.Array(user.Tickers), user.JoinDate)
	return err
}

// parseUserFromBody Parses a user struct from a request body
func parseUserFromBody(req *http.Request) (schema.User, error) {
	var user schema.User
	err := util.DecodeJSON(req.Body, &user)
	return user, err
}

// TickerRequest Holds ticker and user info
type TickerRequest struct {
	UserID string `json:"id"`
	Ticker string `json:"ticker"`
}

// HandleTickerRequest Handles requests regarding a users ticker
func (env *Env) HandleTickerRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		return httputil.MethodNotAllowed
	}
	var tickerRequest TickerRequest
	err := util.DecodeJSON(r.Body, &tickerRequest)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	if !checkTickerIsValid(tickerRequest.Ticker, env.db) {
		return httputil.BadRequest
	}
	return env.handleTickerUpdate(w, tickerRequest, r.Method)
}

// checkTickerIsValid Checks if ticker exists in database
func checkTickerIsValid(ticker string, db *sql.DB) bool {
	var foundTicker string
	err := db.QueryRow("SELECT TICKER FROM STOCK WHERE TICKER=$1", ticker).Scan(&foundTicker)
	if err == nil && foundTicker == ticker {
		return true
	}
	return false
}

// handleTickerUpdate Handles the update of a users tickers
func (env *Env) handleTickerUpdate(
	w http.ResponseWriter, tickerRequest TickerRequest, method string) error {
	tickers, err := getUserTickers(tickerRequest.UserID, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	err = alterTickers(&tickers, tickerRequest.Ticker, method)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	err = updateUserTickers(tickers, tickerRequest.UserID, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendOK(w)
	return nil
}

// alterTickers Changes a slice of tickers based on the method passed
func alterTickers(tickers *stock.Tickers, ticker, method string) error {
	switch method {
	case http.MethodPost:
		return tickers.Add(ticker)
	case http.MethodDelete:
		return tickers.Remove(ticker)
	default:
		return httputil.MethodNotAllowed
	}
}

// getUserTickers Retrives the tickers of a given user
func getUserTickers(userID string, db *sql.DB) (stock.Tickers, error) {
	tickers := make(stock.Tickers, 0)
	err := db.QueryRow("SELECT TICKERS FROM APP_USER WHERE ID=$1", userID).Scan(&tickers)
	if err != nil {
		return tickers, err
	}
	return tickers, nil
}

// updateUserTickers Stores the update of a user tickers in the database
func updateUserTickers(tickers stock.Tickers, userID string, db *sql.DB) error {
	stmt, err := db.Prepare("UPDATE APP_USER SET TICKERS=$1 WHERE ID=$2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(pq.Array(tickers), userID)
	return err
}
