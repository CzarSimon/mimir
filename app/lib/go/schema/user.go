package schema

import (
	"time"

	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
)

// User holds user information.
type User struct {
	ID       string        `json:"id"`
	Email    string        `json:"email,omitempty"`
	Tickers  stock.Tickers `json:"tickers"`
	JoinDate time.Time     `json:"joinDate"`
}

// NewUser creates new user based on a userID.
func NewUser(ID string) User {
	return User{
		ID:       ID,
		Tickers:  stock.InitalTickers,
		JoinDate: getCurrentTimestamp(),
	}
}

// Session Holds user session info
type Session struct {
	UserID       string    `json:"userId"`
	SessionStart time.Time `json:"sessionStart"`
}

// NewSession Creates a new user session
func NewSession(user User) Session {
	return Session{
		UserID:       user.ID,
		SessionStart: getCurrentTimestamp(),
	}
}

// getCurrentTimestamp get the current timestamp in the general mimir timezone.
func getCurrentTimestamp() time.Time {
	return time.Now().UTC()
}
