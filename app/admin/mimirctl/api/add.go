package api

import (
	"fmt"
	"net/http"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
)

// AddLabeledSpam instructs the admin api to store a labeled spam candidate.
func AddLabeledSpam(candidate spam.Candidate) {
	resp, err := makePostRequest(createRoute("spam"), candidate)
	if err != nil {
		checkErr(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		checkErr(fmt.Errorf("Error status. %d", resp.StatusCode))
	}
	fmt.Println("Spam labeled")
}
