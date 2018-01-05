package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"testing"
	"time"
)

const TestDayLayout = "2006-01-02"

func TestSeparateTimestamps(t *testing.T) {
	timestamps := getSampleTimestamps()
	bt, wt := SeparateTimesamps(timestamps)
	noBusdays := len(bt)
	if noBusdays != 20 {
		t.Error(
			"Expeted busdays:", 6,
			"Found:", noBusdays,
		)
	}
	noWeekdays := len(wt)
	if noWeekdays != 6 {
		t.Error(
			"Expeted weekend days:", 6,
			"Found:", noWeekdays,
		)
	}
}

func TestIsBusinessDay(t *testing.T) {
	sunday, _ := time.Parse(TestDayLayout, "2017-12-03")
	if isBusinessDay(sunday) {
		t.Error(sunday, "is a sunday, not a business day")
	}
	tuseday, _ := time.Parse(TestDayLayout, "2017-12-05")
	if !isBusinessDay(tuseday) {
		t.Error(tuseday, "is a tuseday, not a weekend day")
	}
}

func TestNewDayID(t *testing.T) {
	dayStrs := []string{
		"2017-05-07",
		"2017-04-04",
		"1999-11-30",
		"2002-03-17",
		"2022-05-01",
	}
	expected := []string{
		"2017-5-7",
		"2017-4-4",
		"1999-11-30",
		"2002-3-17",
		"2022-5-1",
	}
	for i, dayStr := range dayStrs {
		date, _ := time.Parse(TestDayLayout, dayStr)
		dayID := NewDayID(date)
		if dayID != expected[i] {
			t.Error("DayID creation failed, Expected: ", expected[i], " Found: ", dayID)
		}
	}
}

func TestReduceVolume(t *testing.T) {
	expectedVolume := getSampleVolume()
	sort.Ints(expectedVolume)
	foundVolume := ReduceVolume(getSampleTimestamps())
	sort.Ints(foundVolume)
	for i, found := range foundVolume {
		if found != expectedVolume[i] {
			t.Error(
				"ReduceVolume failed. \nExpected:", expectedVolume,
				"\nFound:   ", foundVolume,
			)
			break
		}
	}
}

func TestNewHourMap(t *testing.T) {
	hourMap := NewHourMap()
	mapSize := len(hourMap)
	if mapSize != 24 {
		t.Error("NewHourMap failed. Unexpected size:", mapSize, "expected", 24)
	}
	for i, timestamps := range hourMap {
		if len(timestamps) != 0 {
			t.Error("NewHourMap failed. Expeted empty slices, found:", len(timestamps), "for index", i)
		}
	}
}

func TestCalcMean(t *testing.T) {
	data := getSampleVolume()
	expectedMean := 3.25
	foundMean := CalculateMean(data)
	if expectedMean != foundMean {
		t.Error("CalculateMean failed: Expected:", expectedMean, "found:", foundMean)
	}
}

func TestCalcStdev(t *testing.T) {
	data := getSampleVolume()
	expectedStdevLow := 1.388
	expectedStdevHigh := 1.389
	foundStdev := CalculateStdev(data)
	if expectedStdevLow > foundStdev || expectedStdevHigh < foundStdev {
		t.Error(
			"CalculateStdev failed: Expected between:[",
			expectedStdevLow, ",", expectedStdevHigh, "] found:", foundStdev,
		)
	}
}

func TestMapTimestamps(t *testing.T) {
	testData := getSampleTimestamps()
	index15 := []int{3, 14, 19, 23}
	expected15 := make([]time.Time, 0)
	expected14 := make([]time.Time, 0)
	for i, timestamp := range testData {
		if valueIsIn(i, index15) {
			expected15 = append(expected15, timestamp)
		} else {
			expected14 = append(expected14, timestamp)
		}
	}
	found := MapTimestamps(testData)
	err := testTimstampMapContent(expected14, found[14])
	if err != nil {
		t.Error(err.Error())
	}
	err = testTimstampMapContent(expected15, found[15])
	if err != nil {
		t.Error(err.Error())
	}
}

func testTimstampMapContent(expected []time.Time, found []time.Time) error {
	if len(found) != len(expected) {
		return errors.New("Unexpected length, expexted: " + string(len(expected)) + " found " + string(len(found)))
	}
	for i, foundTimestamp := range found {
		if foundTimestamp != expected[i] {
			fmt.Println(expected[i])
			fmt.Println(foundTimestamp)
			return errors.New("Unexpected time")
		}
	}
	return nil
}

func valueIsIn(candidate int, values []int) bool {
	for _, value := range values {
		if candidate == value {
			return true
		}
	}
	return false
}

func TestReduceVolumeByDay(t *testing.T) {
	hourMap := getSampleHourMap()
	expectedVolumes := getSampleVolume()
	sort.Ints(expectedVolumes)
	for _, foundVolume := range ReduceVolumeByDay(hourMap) {
		sort.Ints(foundVolume)
		if len(foundVolume) != len(expectedVolumes) {
			t.Error("Unexpected map size, expected", len(expectedVolumes), "found", len(foundVolume))
		}
		for i, expectedVolume := range expectedVolumes {
			if foundVolume[i] != expectedVolume {
				t.Error("Unexpected volume, expected:", expectedVolume, "found", foundVolume[i])
			}
		}
	}
}

func TestCalcNumberOfDays(t *testing.T) {
	startDay, _ := time.Parse(TestDayLayout, "2017-09-01")
	wrongEnd, _ := time.Parse(TestDayLayout, "2017-08-31")
	endDay, _ := time.Parse(TestDayLayout, "2017-10-10")
	expectedDays := NumberOfDays{
		Business: 28,
		Weekend:  12,
	}
	_, err := CalcNumberOfDays(startDay, wrongEnd)
	if err == nil {
		t.Error(
			"Error should not be nil, startDate greater than end date for",
			"start =", startDay, "end =", wrongEnd,
		)
	}
	foundDays, err := CalcNumberOfDays(startDay, endDay)
	if err != nil {
		t.Error("Unexpected error for start =", startDay, "end =", wrongEnd)
	}
	if foundDays.Business != expectedDays.Business {
		t.Error("Unexpeted no. of busdays =", foundDays.Business, "expected =", expectedDays.Business)
	}
	if foundDays.Weekend != expectedDays.Weekend {
		t.Error("Unexpeted no. of weekend =", foundDays.Weekend, "expected =", expectedDays.Weekend)
	}
}

func TestFillVolumes(t *testing.T) {
	volumes := []int{1, 4, 3, 7, 2}
	fullLength := 8
	expectedVolumes := []int{1, 4, 3, 7, 2, 0, 0, 0}
	foundVolumes := fillVolumes(volumes, fullLength)
	if len(foundVolumes) != fullLength {
		t.Error("Unexpected filled length, expected", fullLength, "found", len(foundVolumes))
	}
	for i, expectedVolume := range expectedVolumes {
		if foundVolumes[i] != expectedVolume {
			t.Error("Unexpected volume, expected:", expectedVolume, "found", foundVolumes[i])
		}
	}
}

func TestFillVolumeMap(t *testing.T) {
	expectedLength := 12
	foundVolumeMap := FillVolumeMap(getSampleVolumeMap(), expectedLength)
	for key, foundVolumes := range foundVolumeMap {
		if len(foundVolumes) != expectedLength {
			t.Error(
				"Unexpected length off filled volumes for key = ", key,
				", expected = ", expectedLength, " found = ", len(foundVolumes),
			)
		}
	}
}

func TestCalculateMeanPerHour(t *testing.T) {
	foundMeans := CalculateMeanPerHour(getSampleVolumeMap())
	expectedMean := 3.25
	for key, foundMean := range foundMeans {
		if foundMean != expectedMean {
			t.Error("Unexpected mean for key = ", key, "found =", foundMean, "expected =", expectedMean)
		}
	}
}

func TestCalculateStdevPerHour(t *testing.T) {
	foundStdev := CalculateStdevPerHour(getSampleVolumeMap())
	expectedStdevLow := 1.388
	expectedStdevHigh := 1.389
	for key, foundStdev := range foundStdev {
		if expectedStdevLow > foundStdev || expectedStdevHigh < foundStdev {
			t.Error(
				"Unexpected stdev for key = ", key,
				"found =", foundStdev, "expected bewtween",
				expectedStdevLow, "and", expectedStdevHigh,
			)
		}
	}
}

func TestGetLatestDate(t *testing.T) {
	date1, _ := time.Parse(TestDayLayout, "2010-05-05")
	date2, _ := time.Parse(TestDayLayout, "2010-05-06")
	found1 := getLatestDate(date1, date2)
	if !date2.Equal(found1) {
		t.Error("Wrong date picked, expected = ", date2, "picked =", found1)
	}
	found2 := getLatestDate(date2, date1)
	if !date2.Equal(found2) {
		t.Error("Wrong date picked, expected = ", date2, "picked =", found1)
	}
}

/* --- Test data generators --- */

func getSampleVolume() []int {
	volumes := []int{1, 3, 3, 4, 5, 5, 2, 3}
	return volumes
}

func getSampleTimestamps() []time.Time {
	dates := []string{
		"2017-10-01T14:00:01+00:00",
		"2017-10-02T14:00:00+00:00",
		"2017-10-02T14:01:00+00:00",
		"2017-10-02T15:00:01+00:00",
		"2017-10-03T14:00:01+00:00",
		"2017-10-03T14:00:01+00:00",
		"2017-10-03T14:00:01+00:00",
		"2017-10-04T14:00:01+00:00",
		"2017-10-04T14:00:01+00:00",
		"2017-10-04T14:00:01+00:00",
		"2017-10-04T14:00:01+00:00",
		"2017-10-05T14:00:01+00:00",
		"2017-10-05T14:00:01+00:00",
		"2017-10-05T14:00:01+00:00",
		"2017-10-05T15:00:01+00:00",
		"2017-10-05T14:00:01+00:00",
		"2017-10-06T14:00:01+00:00",
		"2017-10-06T14:00:01+00:00",
		"2017-10-06T14:00:01+00:00",
		"2017-10-06T15:00:01+00:00",
		"2017-10-06T14:00:01+00:00",
		"2017-10-07T14:00:01+00:00",
		"2017-10-07T14:00:01+00:00",
		"2017-10-08T15:00:01+00:00",
		"2017-10-08T14:00:01+00:00",
		"2017-10-08T14:00:01+00:00",
	}
	timestamps := make([]time.Time, len(dates))
	for idx, date := range dates {
		timestamp, err := time.Parse(time.RFC3339, date)
		if err != nil {
			log.Fatal(err)
		}
		timestamps[idx] = timestamp
	}
	return timestamps
}

func getSampleHourMap() map[int][]time.Time {
	hourMap := make(map[int][]time.Time)
	for i := 0; i < HoursInDay; i++ {
		hourMap[i] = getSampleTimestamps()
	}
	return hourMap
}

func getSampleVolumeMap() map[int][]int {
	hourMap := make(map[int][]int)
	for i := 0; i < HoursInDay; i++ {
		hourMap[i] = getSampleVolume()
	}
	return hourMap
}
