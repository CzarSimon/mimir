package spam

import (
	"fmt"
)

// Supported spam labels
const (
	SPAM_LABEL     = "SPAM"
	NON_SPAM_LABEL = "NON-SPAM"
)

// Candidate Struct that can be classified as spam
type Candidate struct {
	Text  string `json:"text"`
	Label string `json:"label"`
}

// NewCandidate Creates a new unlabeled spam Candidate
func NewCandidate(text string) Candidate {
	return Candidate{
		Text: text,
	}
}

// IsSpam Checks whether a spam candidate is classified as spam
func (candidate Candidate) IsSpam() bool {
	return candidate.Label == SPAM_LABEL
}

// String Returns a string representation of a spam Candidate
func (candidate Candidate) String() string {
	return fmt.Sprintf("Text=%s Label=%s", candidate.Text, candidate.Label)
}
