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
		JoinDate: time.Now().UTC(),
	}
}
