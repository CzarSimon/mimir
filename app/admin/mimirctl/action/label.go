package action

import (
	"fmt"

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
	fmt.Println("Getting spam candidates")
	return nil
}
