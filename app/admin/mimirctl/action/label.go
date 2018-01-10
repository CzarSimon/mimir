package action

import (
	"bytes"
	"fmt"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
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

// labelSpam interacts with the user to label a spam candidate.
func labelSpam(candidate spam.Candidate) {
	fmt.Printf("Text: %s \nChose label (press any other key to skip)\n")
	choice := getInput(spamChoices.String())
	label, ok := spamChoices[choice]
	if !ok {
		fmt.Println("Skipping")
		return
	}
	candidate.Label = label
	fmt.Println(candidate)
}

// choiceMap map used for linking user input to a string label.
type choiceMap map[string]string

// spamChoices map for linking choices to spam labels.
var spamChoices = choiceMap{
	"1": spam.SPAM_LABEL,
	"0": spam.NON_SPAM_LABEL,
}

// String returns a string representation of a choiceMap.
func (cm choiceMap) String() string {
	var buf bytes.Buffer
	for key, choice := range cm {
		buf.WriteString(fmt.Sprintf("%s. - %s\n", key, choice))
	}
	return buf.String()
}
