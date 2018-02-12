package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/CzarSimon/util"
	"github.com/goware/urlx"
)

const (
	ForbiddenURLsConfigFile = "resources/forbidden-urls.json"
	WWWPrefix               = "www."
)

// ForbiddenURLs set of forbidden urls which not to send to ranking.
type ForbiddenURLs map[string]bool

// IsForbidden checks if a URL is forbidden.
func (URLs ForbiddenURLs) IsForbidden(URL string) bool {
	url, err := urlx.Parse(URL)
	if err != nil {
		log.Println(err)
		return false
	}
	_, forbidden := URLs[url.Host]
	return forbidden
}

// FilterURLs filters out forbidden URLs from a list of candidates.
func (URLs ForbiddenURLs) FilterURLs(candidates []string) []string {
	filteredURLs := make([]string, 0, len(candidates))
	for _, URL := range candidates {
		if !URLs.IsForbidden(URL) {
			filteredURLs = append(filteredURLs, URL)
		}
	}
	return filteredURLs
}

// GetForbiddenURLs gets forbidden urls from configuration and structures into a set.
func GetForbiddenURLs() ForbiddenURLs {
	URLs, err := readForbiddenConfig()
	util.CheckErrFatal(err)
	forbiddenURLs := make(ForbiddenURLs)
	for _, URL := range URLs {
		forbiddenURLs[URL] = true
		forbiddenURLs[WWWPrefix+URL] = true
	}
	return forbiddenURLs
}

// readForbiddenConfig reads forbidden urls list from config file.
func readForbiddenConfig() ([]string, error) {
	forbiddenList := make([]string, 0)
	content, err := ioutil.ReadFile(ForbiddenURLsConfigFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &forbiddenList)
	if err != nil {
		return nil, err
	}
	return forbiddenList, nil
}
