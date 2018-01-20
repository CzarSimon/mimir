package user

import (
	"time"

	"github.com/CzarSimon/mimir/app/lib/go/schema/stock"
)

// User Holds user information
type User struct {
	ID       string        `json:"id"`
	Email    string        `json:"email,omitempty"`
	Tickers  stock.Tickers `json:"tickers"`
	JoinDate time.Time     `json:"joinDate"`
}

// NewUser Creates new user based on a userID
func New(ID string) User {
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

// getCurrentTimestamp gets a current timestamp in the correct timezone.
func getCurrentTimestamp() time.Time {
	return time.Now().UTC()
}
