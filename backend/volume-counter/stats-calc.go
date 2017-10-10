package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/CzarSimon/util"
)

// HoursInDay Number of hours in a day
const HoursInDay = 24

// StatsCalc Calculates mean and stdev volumes for tweet about a ticker
func StatsCalc(config Config) {
	fmt.Println(config.Server.ToURL(config.Routes.StatsResult))
	db := util.ConnectPG(config.DB)
	defer db.Close()
	tickers, err := getTrackedTickers(db)
	if err != nil {
		util.LogErr(err)
		return
	}
	log.Println(tickers)
	statistics, err := CalcStatsForTickers(tickers, config.StableDate, db)
	if err != nil {
		util.LogErr(err)
		return
	}
	err = sendStatsResult(statistics, config)
	if err != nil {
		util.LogErr(err)
	}
}

// TickerStatistcs Map of statiscs per with ticker as key
type TickerStatistcs map[string]Statistics

// Statistics Mean and Stdev timeseries
type Statistics struct {
	Mean  DayTypeTimeseries `json:"mean"`
	Stdev DayTypeTimeseries `json:"stdev"`
}

// NewStatistics Creates a new, empty Statistics struct
func NewStatistics() Statistics {
	return Statistics{
		Mean:  NewDayTypeTimeseries(),
		Stdev: NewDayTypeTimeseries(),
	}
}

// CalcMean Calculates mean volume per hour per day type
func (statistics *Statistics) CalcMean(busdayVolumes, weekendVolumes map[int][]int) {
	statistics.Mean = DayTypeTimeseries{
		BusinessDays: CalculateMeanPerHour(busdayVolumes),
		WeekendDays:  CalculateMeanPerHour(weekendVolumes),
	}
}

// CalculateMeanPerHour Calculates volume mean per hour
func CalculateMeanPerHour(volumeMap map[int][]int) StatsTimeseries {
	means := make(StatsTimeseries, HoursInDay)
	var volumes []int
	for i := range means {
		volumes, _ = volumeMap[i]
		means[i] = CalculateMean(volumes)
	}
	return means
}

// CalculateMean Caluculates a mean value based on a slice of ints
func CalculateMean(data []int) float64 {
	var sum float64
	for _, value := range data {
		sum += float64(value)
	}
	return (sum / float64(len(data)))
}

// CalcStdev Calculates standard deviation of volume per hour per day type
func (statistics *Statistics) CalcStdev(busdayVolumes, weekendVolumes map[int][]int) {
	statistics.Stdev = DayTypeTimeseries{
		BusinessDays: CalculateStdevPerHour(busdayVolumes),
		WeekendDays:  CalculateStdevPerHour(weekendVolumes),
	}
}

// CalculateStdevPerHour Calculates volume standard deviation per hour
func CalculateStdevPerHour(volumeMap map[int][]int) StatsTimeseries {
	stdev := make(StatsTimeseries, HoursInDay)
	var volumes []int
	for i := range stdev {
		volumes, _ = volumeMap[i]
		stdev[i] = CalculateStdev(volumes)
	}
	return stdev
}

// CalculateStdev Caluculates a standard deviation value based on a slice of ints
func CalculateStdev(data []int) float64 {
	mean := CalculateMean(data)
	var squareSum float64
	for _, value := range data {
		squareSum += math.Pow((float64(value) - mean), 2)
	}
	return math.Sqrt(squareSum / (float64(len(data) - 1)))
}

// DayTypeTimeseries Holds volume statiscs separated by day type
type DayTypeTimeseries struct {
	BusinessDays StatsTimeseries `json:"busdays"`
	WeekendDays  StatsTimeseries `json:"weekend_days"`
}

// NewDayTypeTimeseries Creates a new, empty DayTypeTimeseries struct
func NewDayTypeTimeseries() DayTypeTimeseries {
	return DayTypeTimeseries{
		BusinessDays: make(StatsTimeseries, HoursInDay),
		WeekendDays:  make(StatsTimeseries, HoursInDay),
	}
}

// StatsTimeseries Holds volume statiscs in a timeseries
type StatsTimeseries []float64

// getTrackedTickers Queries for all tracked tickers in the database
func getTrackedTickers(db *sql.DB) ([]string, error) {
	tickers := make([]string, 0)
	rows, err := db.Query("SELECT TICKER FROM STOCKS WHERE IS_TRACKED=TRUE")
	defer rows.Close()
	if err != nil {
		return tickers, err
	}
	var ticker string
	for rows.Next() {
		err = rows.Scan(&ticker)
		if err != nil {
			return tickers, err
		}
		tickers = append(tickers, ticker)
	}
	return tickers, nil
}

// getStoredAtDate Gets the date at which a stock was first stored and tracked
func getStoredAtDate(ticker string, db *sql.DB) (time.Time, error) {
	var storedAt time.Time
	query := "SELECT storedat FROM STOCKS WHERE TICKER=$1"
	err := db.QueryRow(query, ticker).Scan(&storedAt)
	if err != nil {
		return storedAt, err
	}
	logTickerStoredAt(ticker, storedAt)
	return storedAt, nil
}

// logTickerStoredAt Logs stored at date for ticker
func logTickerStoredAt(ticker string, date time.Time) {
	log.Printf("%s stored at = %d-%d-%d\n", ticker, date.Year(), date.Month(), date.Day())
}

// getLatestDate Returns the latest of two supplied dates
func getLatestDate(date1, date2 time.Time) time.Time {
	if date1.After(date2) {
		return date1
	}
	return date2
}

// CalcStatsForTickers Calculates mean and stdev for a supplied list of tickers
func CalcStatsForTickers(tickers []string, stableDate time.Time, db *sql.DB) (TickerStatistcs, error) {
	statistics := make(TickerStatistcs)
	for _, ticker := range tickers {
		timestamps, err := GetTickerTimestamps(ticker, stableDate, db)
		if err != nil {
			util.LogErr(err)
			continue
		}
		firstTickerDate, err := getStoredAtDate(ticker, db)
		if err != nil {
			util.LogErr(err)
			continue
		}
		numberOfDays, err := CalcNumberOfDays(getLatestDate(firstTickerDate, stableDate), getToday())
		if err != nil {
			util.LogErr(err)
			continue
		}
		tickerStats := CalcStatsForTicker(timestamps, numberOfDays)
		if err != nil {
			util.LogErr(err)
			continue
		}
		statistics[ticker] = tickerStats
	}
	return statistics, nil
}

// CalcStatsForTicker Calculates mean and stdev for a supplied ticker
func CalcStatsForTicker(timestamps []time.Time, numberOfDays NumberOfDays) Statistics {
	statistics := NewStatistics()
	busdayTimestamps, weekendTimestamps := SeparateTimesamps(timestamps)
	busdayVolumes := FillVolumeMap(
		ReduceTimestamps(busdayTimestamps), numberOfDays.Business)
	weekendVolumes := FillVolumeMap(
		ReduceTimestamps(weekendTimestamps), numberOfDays.Weekend)
	statistics.CalcMean(busdayVolumes, weekendVolumes)
	statistics.CalcStdev(busdayVolumes, weekendVolumes)
	return statistics
}

// FillVolumeMap Pads volume with zeros map if on some days no volume was recorded
func FillVolumeMap(volumeMap map[int][]int, fullLength int) map[int][]int {
	for key, volumes := range volumeMap {
		volumeMap[key] = fillVolumes(volumes, fullLength)
	}
	return volumeMap
}

// fillVolumes Fills unrecorded volumes with zeros
func fillVolumes(volumes []int, fullLength int) []int {
	fillVolumes := make([]int, fullLength-len(volumes))
	return append(volumes, fillVolumes...)
}

// NewHourMap Creates an empty hour -> timestamp map
func NewHourMap() map[int][]time.Time {
	hourMap := make(map[int][]time.Time)
	for i := 0; i < HoursInDay; i++ {
		hourMap[i] = make([]time.Time, 0)
	}
	return hourMap
}

// MapTimestamps Maps in timestamps to an hour map
func MapTimestamps(timestamps []time.Time) map[int][]time.Time {
	hourMap := NewHourMap()
	var hour int
	for _, timestamp := range timestamps {
		hour = timestamp.Hour()
		hourMap[hour] = append(hourMap[hour], timestamp)
	}
	return hourMap
}

// ReduceVolumeByDay Reduces recoded timestamps for every hour and returns as a
// map of hour -> volumes
func ReduceVolumeByDay(hourMap map[int][]time.Time) map[int][]int {
	volumeMap := make(map[int][]int)
	for hour, timestamps := range hourMap {
		volumeMap[hour] = ReduceVolume(timestamps)
	}
	return volumeMap
}

// ReduceVolume Counts the number of timestamps per day and returns as volumes
func ReduceVolume(timestamps []time.Time) []int {
	dayMap := make(map[string]int)
	var dayID string
	var volume int
	var present bool
	for _, timestamp := range timestamps {
		dayID = NewDayID(timestamp)
		volume, present = dayMap[dayID]
		if present {
			dayMap[dayID] = volume + 1
		} else {
			dayMap[dayID] = 1
		}
	}
	volumes := make([]int, 0)
	for _, volume := range dayMap {
		volumes = append(volumes, volume)
	}
	return volumes
}

// NewDayID Turns a timestamp into a day idetifier
func NewDayID(timestamp time.Time) string {
	return fmt.Sprintf("%d-%d-%d", timestamp.Year(), timestamp.Month(), timestamp.Day())
}

// ReduceTimestamps Maps timestamps to the hour in which they occured
// and sums up the volume per day
func ReduceTimestamps(timestamps []time.Time) map[int][]int {
	hourMap := MapTimestamps(timestamps)
	return ReduceVolumeByDay(hourMap)
}

// SeparateTimesamps Separates a list of timestamps into one list containing the
// business day timestamps and one for weekend days
func SeparateTimesamps(timestamps []time.Time) ([]time.Time, []time.Time) {
	busdayTimestamps := make([]time.Time, 0)
	weekdayTimestamps := make([]time.Time, 0)
	for _, timestamp := range timestamps {
		if isBusinessDay(timestamp) {
			busdayTimestamps = append(busdayTimestamps, timestamp)
		} else {
			weekdayTimestamps = append(weekdayTimestamps, timestamp)
		}
	}
	return busdayTimestamps, weekdayTimestamps
}

// GetTickerTimestamps Gets the timestamps of recoreded tweets about a supplied ticker
func GetTickerTimestamps(ticker string, stableDate time.Time, db *sql.DB) ([]time.Time, error) {
	timestamps := make([]time.Time, 0)
	query := "SELECT createdAt FROM STOCKTWEETS WHERE TICKER=$1 AND createdAt>=$2 AND createdAt IS NOT NULL"
	rows, err := db.Query(query, ticker, stableDate)
	defer rows.Close()
	if err != nil {
		return timestamps, err
	}
	var timestamp time.Time
	for rows.Next() {
		err := rows.Scan(&timestamp)
		if err != nil {
			return timestamps, err
		}
		timestamps = append(timestamps, timestamp)
	}
	return timestamps, nil
}

// sendStatsResult Sends result of stats calc
func sendStatsResult(statistics TickerStatistcs, config Config) error {
	payload, err := json.Marshal(&statistics)
	if err != nil {
		return err
	}
	return Send(payload, config.Server.ToURL(config.Routes.StatsResult))
}

// NumberOfDays Holds the value for number of busniess and weekend days
type NumberOfDays struct {
	Business int
	Weekend  int
}

// CalcNumberOfDays Calculates the number of busniess and weekend days between two dates
func CalcNumberOfDays(start, end time.Time) (NumberOfDays, error) {
	noDays := NumberOfDays{}
	if start.After(end) {
		return noDays, errors.New("Start date greater than end date")
	}
	date := start
	for !date.After(end) {
		if isBusinessDay(date) {
			noDays.Business++
		} else {
			noDays.Weekend++
		}
		date = date.AddDate(0, 0, 1)
	}
	return noDays, nil
}

// isBusinessDay Checks if the supplied date is a business day or not
func isBusinessDay(date time.Time) bool {
	weekday := date.Weekday()
	return !(weekday == time.Saturday || weekday == time.Sunday)
}

// getToday Returns todays date
func getToday() time.Time {
	return time.Now().UTC()
}
