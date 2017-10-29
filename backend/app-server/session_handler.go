package main

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/CzarSimon/util"
)

// Session Holds user session info
type Session struct {
	UserID       string    `json:"userId"`
	SessionStart time.Time `json:"sessionStart"`
}

// NewSession Creates a new user session
func NewSession(user User) Session {
	return Session{
		UserID:       user.ID,
		SessionStart: time.Now().UTC(),
	}
}

// HandleSessionRequest Records a session start of the user specified in the request
func (env *Env) HandleSessionRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		util.SendErrStatus(
			res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	user, err := parseUserFromBody(req)
	if err != nil {
		util.SendErrStatus(res, err, http.StatusBadRequest)
		return
	}
	err = storeUserSession(NewSession(user), env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// storeUserSession Stores a new user sesssion
func storeUserSession(session Session, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO SESSION(USER_ID, SESSION_START) VALUES($1, $2)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.UserID, session.SessionStart)
	if err != nil {
		return err
	}
	return nil
}
