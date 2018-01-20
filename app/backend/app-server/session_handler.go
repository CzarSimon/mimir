package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/mimir/app/lib/go/schema/user"
)

// Session Holds user session info
type Session struct {
	UserID       string    `json:"userId"`
	SessionStart time.Time `json:"sessionStart"`
}

// NewSession Creates a new user session
func NewSession(usr user.User) Session {
	return Session{
		UserID:       usr.ID,
		SessionStart: time.Now().UTC(),
	}
}

// HandleSessionRequest Records a session start of the user specified in the request
func (env *Env) HandleSessionRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return httputil.MethodNotAllowed
	}
	usr, err := parseUserFromBody(r)
	if err != nil {
		return httputil.BadRequest
	}
	err = storeUserSession(NewSession(usr), env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendOK(w)
	return nil
}

// storeUserSession Stores a new user sesssion
func storeUserSession(session Session, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO SESSION(USER_ID, SESSION_START) VALUES($1, $2)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.UserID, session.SessionStart)
	return err
}
