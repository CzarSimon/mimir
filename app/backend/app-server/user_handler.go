package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/query"
	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
	"github.com/CzarSimon/mimir/app/lib/go/schema/user"
	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

const USER_ID_KEY = "id"

// HandleUserRequest Handles retrival of user or creation of new user
func (env *Env) HandleUserRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		err := httputil.MethodNotAllowed
		util.SendErrStatus(res, err, err.Status)
		return
	}
	userID, err := query.ParseValue(req, USER_ID_KEY)
	log.Println(userID)
	if err != nil {
		util.SendErrStatus(res, err, http.StatusBadRequest)
		return
	}
	user, err := GetUser(userID, env.db)
	if err != nil {
		util.LogErr(err)
		util.SendErrRes(res, err)
		return
	}
	err = storeUserSession(NewSession(user), env.db)
	if err != nil {
		log.Println("Unable to store user session")
	}
	jsonBody, err := json.Marshal(user)
	if err != nil {
		util.LogErr(err)
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// GetUser Retrives user from datababase if exist, returns new user if not
func GetUser(userID string, db *sql.DB) (user.User, error) {
	var usr user.User
	err := db.QueryRow("SELECT ID, TICKERS, JOIN_DATE FROM APP_USER WHERE ID=$1", userID).Scan(
		&usr.ID, &usr.Tickers, &usr.JoinDate)
	if err == sql.ErrNoRows {
		newUser := user.New(userID)
		err = StoreNewUser(newUser, db)
		return newUser, err
	}
	return usr, err
}

// StoreNewUser Stores a newly created user in the supplied database
func StoreNewUser(usr user.User, db *sql.DB) error {
	stmt, err := db.Prepare(
		"INSERT INTO APP_USER(ID, TICKERS, JOIN_DATE) VALUES ($1, $2, $3)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(usr.ID, pq.Array(usr.Tickers), usr.JoinDate)
	return err
}

// parseUserFromBody Parses a user struct from a request body
func parseUserFromBody(req *http.Request) (user.User, error) {
	var usr user.User
	err := util.DecodeJSON(req.Body, &usr)
	return usr, err
}

// TickerRequest Holds ticker and user info
type TickerRequest struct {
	UserID string `json:"id"`
	Ticker string `json:"ticker"`
}

// HandleTickerRequest Handles requests regarding a users ticker
func (env *Env) HandleTickerRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost && req.Method != http.MethodDelete {
		util.SendErrStatus(
			res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	var tickerRequest TickerRequest
	err := util.DecodeJSON(req.Body, &tickerRequest)
	if err != nil {
		util.SendErrStatus(
			res, errors.New("Could not parse request"), http.StatusBadRequest)
		return
	}
	if !checkTickerIsValid(tickerRequest.Ticker, env.db) {
		util.SendErrStatus(res, errors.New("Invalid ticker"), http.StatusBadRequest)
		return
	}
	env.handleTickerUpdate(res, tickerRequest, req.Method)
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
	res http.ResponseWriter, tickerRequest TickerRequest, method string) {
	tickers, err := getUserTickers(tickerRequest.UserID, env.db)
	if err != nil {
		util.SendErrRes(res, errors.New("Could not alter tickers"))
		return
	}
	err = alterTickers(&tickers, tickerRequest.Ticker, method)
	if err != nil {
		util.SendErrRes(res, errors.New("Could not alter tickers"))
		return
	}
	err = updateUserTickers(tickers, tickerRequest.UserID, env.db)
	if err != nil {
		util.SendErrRes(res, errors.New("Could not alter tickers"))
		return
	}
	util.SendOK(res)
}

// alterTickers Changes a slice of tickers based on the method passed
func alterTickers(tickers *stock.Tickers, ticker, method string) error {
	switch method {
	case http.MethodPost:
		return tickers.Add(ticker)
	case http.MethodDelete:
		return tickers.Remove(ticker)
	default:
		return errors.New("Method not allowed")
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
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(pq.Array(tickers), userID)
	if err != nil {
		return err
	}
	return nil
}
