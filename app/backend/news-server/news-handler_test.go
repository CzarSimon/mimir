package main

import (
	"errors"
	"testing"
	"time"
)

func TestCalcDateAdjustment1W(t *testing.T) {
	expectedDate := time.Now().UTC().AddDate(0, 0, -7)
	foundDate := calcDateAdjustment("1W")
	if AssertDate(expectedDate, foundDate) != nil {
		t.Error(
			"Unexpected date for period 1W expected =",
			expectedDate, "found =", foundDate,
		)
	}
	expectedDate = time.Now().UTC()
	foundDate = calcDateAdjustment("TODAY")
	if AssertDate(expectedDate, foundDate) != nil {
		t.Error(
			"Unexpected date for period TODAY expected =",
			expectedDate, "found =", foundDate,
		)
	}
	expectedDate = time.Now().UTC().AddDate(0, -1, 0)
	foundDate = calcDateAdjustment("1M")
	if AssertDate(expectedDate, foundDate) != nil {
		t.Error(
			"Unexpected date for period 1W expected =",
			expectedDate, "found =", foundDate,
		)
	}
	expectedDate = time.Now().UTC().AddDate(0, -3, 0)
	foundDate = calcDateAdjustment("3M")
	if AssertDate(expectedDate, foundDate) != nil {
		t.Error(
			"Unexpected date for period 1W expected =",
			expectedDate, "found =", foundDate,
		)
	}
	expectedDate = time.Now().UTC()
	foundDate = calcDateAdjustment("1N")
	if err := AssertDate(expectedDate, foundDate); err != nil {
		t.Error(
			"Unexpected date for invalid expected =",
			expectedDate, "found =", foundDate,
		)
	}
}

func AssertDate(expectedDate, foundDate time.Time) error {
	if expectedDate.Year() != foundDate.Year() {
		return errors.New("Wrong year")
	}
	if expectedDate.Month() != foundDate.Month() {
		return errors.New("Wrong month")
	}
	if expectedDate.Day() != foundDate.Day() {
		return errors.New("Wrong day")
	}
	return nil
}
