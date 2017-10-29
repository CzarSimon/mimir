package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

// User Holds user information
type User struct {
	ID       string    `json:"id"`
	Email    string    `json:"email,omitempty"`
	Tickers  Tickers   `json:"tickers"`
	JoinDate time.Time `json:"joinDate"`
}

// NewUser Creates new user based on a userID
func NewUser(userID string) User {
	return User{
		ID:       userID,
		Tickers:  GetInitalTickers(),
		JoinDate: time.Now().UTC(),
	}
}

// HandleUserRequest Handles retrival of user or creation of new user
func (env *Env) HandleUserRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		util.SendErrStatus(res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	userID, err := parseUserIDFromQuery(req)
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
func GetUser(userID string, db *sql.DB) (User, error) {
	var user User
	err := db.QueryRow("SELECT ID, TICKERS, JOIN_DATE FROM APP_USER WHERE ID=$1", userID).Scan(
		&user.ID, &user.Tickers, &user.JoinDate)
	if err == sql.ErrNoRows {
		newUser := NewUser(userID)
		err = StoreNewUser(newUser, db)
		return newUser, err
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

// StoreNewUser Stores a newly created user in the supplied database
func StoreNewUser(newUser User, db *sql.DB) error {
	stmt, err := db.Prepare(
		"INSERT INTO APP_USER(ID, TICKERS, JOIN_DATE) VALUES ($1, $2, $3)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newUser.ID, pq.Array(newUser.Tickers), newUser.JoinDate)
	if err != nil {
		return err
	}
	return nil
}

// parseUserFromBody Parses a user struct from a request body
func parseUserFromBody(req *http.Request) (User, error) {
	var user User
	err := util.DecodeJSON(req.Body, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// parseUserIDFromQuery Parses a user id from a given request query
func parseUserIDFromQuery(req *http.Request) (string, error) {
	return util.ParseValueFromQuery(req, "id", "No user id supplied")
}

// Tickers Slice of tickers that can be queried and inserted into a postgres database
type Tickers []string

// Scan Scans a slice of tickers
func (tickers *Tickers) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	(*tickers) = util.BytesToStrSlice(bytes)
	return nil
}

// AddTicker Adds new ticker if it is not yet added
func (tickers *Tickers) AddTicker(ticker string) error {
	for _, currentTicker := range *tickers {
		if currentTicker == ticker {
			return errors.New("Ticker already added")
		}
	}
	*tickers = append(*tickers, ticker)
	return nil
}

// RemoveTicker Removes ticket is it is present
func (tickers *Tickers) RemoveTicker(ticker string) error {
	filteredTickers := make([]string, 0)
	tickerPresent := false
	for _, currentTicker := range *tickers {
		if currentTicker != ticker {
			filteredTickers = append(filteredTickers, currentTicker)
		} else {
			tickerPresent = true
		}
	}
	if !tickerPresent {
		return errors.New("Ticker not present")
	}
	*tickers = filteredTickers
	return nil
}

// GetInitalTickers Returns a new users inital tickers
func GetInitalTickers() []string {
	return []string{"AAPL", "FB", "TSLA", "TWTR", "AMZN"}
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
func alterTickers(tickers *Tickers, ticker, method string) error {
	switch method {
	case http.MethodPost:
		return tickers.AddTicker(ticker)
	case http.MethodDelete:
		return tickers.RemoveTicker(ticker)
	default:
		return errors.New("Method not allowed")
	}
}

// getUserTickers Retrives the tickers of a given user
func getUserTickers(userID string, db *sql.DB) (Tickers, error) {
	tickers := make(Tickers, 0)
	err := db.QueryRow("SELECT TICKERS FROM APP_USER WHERE ID=$1", userID).Scan(&tickers)
	if err != nil {
		return tickers, err
	}
	return tickers, nil
}

// updateUserTickers Stores the update of a user tickers in the database
func updateUserTickers(tickers Tickers, userID string, db *sql.DB) error {
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
