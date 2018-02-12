package main

import "testing"

var testURLs = []string{
	"owler.us",
	"www.owler.us",
	"owler.com",
	"www.owler.com",
}

var allowedCandidates = []string{
	"example.com",
	"https://example.com",
	"www.example.com",
	"https://example.com/some/path",
}

var forbiddenCandidates = []string{
	"owler.com",
	"https://owler.com",
	"www.owler.com",
	"https://owler.com/some/path",
}

var allCandidates = append(allowedCandidates, forbiddenCandidates...)

func TestGetForbiddenURLs(t *testing.T) {
	forbidden := GetForbiddenURLs()
	for i, URL := range testURLs {
		_, isPresent := forbidden[URL]
		if !isPresent {
			t.Errorf("%d - Error in GetForbiddenURLs: URL: [%s] not found", i, URL)
		}
	}
}

func TestIsForbidden(t *testing.T) {
	forbidden := GetForbiddenURLs()
	for i, URL := range forbiddenCandidates {
		if !forbidden.IsForbidden(URL) {
			t.Errorf("%d - ForbiddenURLs.IsForbidden wrong. Expected URL:[%s] to be forrbidden", i, URL)
		}
	}
	for i, URL := range allowedCandidates {
		if forbidden.IsForbidden(URL) {
			t.Errorf("%d - ForbiddenURLs.IsForbidden wrong. Expected URL:[%s] to not be forrbidden", i, URL)
		}
	}
}

func TestFilterURLs(t *testing.T) {
	forbidden := GetForbiddenURLs()
	filtered := forbidden.FilterURLs(allCandidates)
	expectedLength := len(allowedCandidates)
	foundLength := len(filtered)
	if expectedLength != foundLength {
		t.Fatalf("forbiddenURLs.FilterURLs found length wrong, Expected=%d Found=%d",
			expectedLength, foundLength)
	}
	for i, foundURL := range filtered {
		expectedURL := allowedCandidates[i]
		if expectedURL != foundURL {
			t.Errorf("%d - forbiddenURLs.FilterURLs wrong URL, Expected=%s Found=%s",
				expectedURL, foundURL)
		}
	}
}
