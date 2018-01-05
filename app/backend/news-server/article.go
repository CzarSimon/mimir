package main

import (
	"errors"

	"github.com/CzarSimon/util"
)

//Article holds values for an article
type Article struct {
	Title             string            `json:"title"`
	Summary           string            `json:"summary"`
	URL               string            `json:"url"`
	Timestamp         string            `json:"timestamp"`
	Keywords          Keywords          `json:"keywords"`
	TwitterReferences TwitterReferences `json:"twitterReferences"`
}

//TwitterReferences is the list of twitter users that has refered an article
type TwitterReferences []int64

//Scan implements scanning a postgres array to TwitterReferences
func (tr *TwitterReferences) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	intSlice, err := util.BytesToIntSlice(bytes, 64)
	if util.IsErr(err) {
		return err
	}
	(*tr) = intSlice
	return nil
}

//Keywords is the list of keywords computed for an article
type Keywords []string

//Scan implements scanning a postgres arrya into a string slice
func (keywords *Keywords) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	strSlice := util.BytesToStrSlice(bytes)
	(*keywords) = strSlice
	return nil
}
