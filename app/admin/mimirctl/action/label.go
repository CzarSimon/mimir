package action

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/lib/go/schema/spam"
	"github.com/urfave/cli"
)

// LabelMap maps resource type to a labeleing action.
var LabelMap = ResourceMap{
	SPAM: getSpamCandidates,
}

// Label enable getting of labeling candidates and labeling them
func Label(c *cli.Context) error {
	function := LabelMap.GetFunc(getResource(c))
	return function(c)
}

// getSpamCandidates gets candidates for spam labeling.
func getSpamCandidates(c *cli.Context) error {
	candidates := api.GetSpamCandidates()
	for i, candidate := range candidates {
		fmt.Println("Candidate %d of %d", i+1, len(candidates))
		labelSpam(candidate)
	}
	return nil
}

func labelSpam(candidate spam.Candidate) {
	fmt.Printf("Text: %s \nChose label (press any other key to skip)\n")
	choice := getInput(fmt.Sprintf(
		"1. - %s , 0. - %s", spam.SPAM_LABEL, spam.NON_SPAM_LABEL))
	if choice == "1" {
		candidate.Label = spam.SPAM_LABEL
	} else if choice == "2" {
		candidate.Label = spam.NON_SPAM_LABEL
	} else {
		fmt.Println("Skipping")
		return
	}
	fmt.Println(candidate)
}
