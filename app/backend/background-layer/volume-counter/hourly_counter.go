package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/CzarSimon/util"
)

func RunVolumeCount(config Config) {
	log.Println("Running volume count")
	err := countVolume(config)
	if err != nil {
		log.Println(err)
	}
}

// VolumeCount Counts the tweet volume in the last hour
func countVolume(config Config) error {
	db, err := config.TweetDB.Connect()
	util.CheckErrFatal(err)
	defer db.Close()
	tickers, err := getAllTickers(db)
	if err != nil {
		return err
	}
	err = countHourlyVolumes(tickers, db)
	if err != nil {
		return err
	}
	return storeVolumes(NewVolumeResult(tickers), config)
}

// Volume Value pair of the ticker and volume for a specific stock
type Volume struct {
	Ticker string
	Count  int64
}

// countHourlyVolumes Performes the calculation of hourly volumes per ticker
func countHourlyVolumes(volumes HourVolumes, db *sql.DB) error {
	query := "SELECT ticker, COUNT(*) FROM stocktweets WHERE createdAt>$1 GROUP BY ticker"
	rows, err := db.Query(query, getHour())
	defer rows.Close()
	if err != nil {
		return err
	}
	var ticker string
	var count int64
	for rows.Next() {
		err = rows.Scan(&ticker, &count)
		if err != nil {
			return err
		}
		volumes[ticker].Volume = count
	}
	return nil
}

// VolumeResult Struct containgn counted voluemes and the hour in which the calcualtion happened.
type VolumeResult struct {
	Hour    int         `json:"hour"`
	Volumes HourVolumes `json:"volumes"`
}

// NewVolumeResult Creates new VolumeResult struct based on a given volume calcualtion
func NewVolumeResult(volumes HourVolumes) VolumeResult {
	return VolumeResult{
		Hour:    getHourNumber(),
		Volumes: volumes,
	}
}

// HourVolume conatains this hours count and the minute the count occured
type HourVolume struct {
	Volume int64 `json:"volume"`
	Minute int   `json:"minute"`
}

// HourVolumes is a ticker -> HourVolume map
type HourVolumes map[string]*HourVolume

// getAllTickers Gets the ticker of each tracked stock and returns an HourVolumes
// map with the current minute and no recorded volume per ticker
func getAllTickers(db *sql.DB) (HourVolumes, error) {
	tickers := make(HourVolumes)
	rows, err := db.Query("SELECT ticker FROM stocks WHERE is_tracked=TRUE")
	defer rows.Close()
	if err != nil {
		return tickers, err
	}
	minute := getMinute()
	var ticker string
	for rows.Next() {
		err = rows.Scan(&ticker)
		if err != nil {
			return tickers, err
		}
		tickers[ticker] = &HourVolume{
			Volume: 0,
			Minute: minute,
		}
	}
	return tickers, nil
}

// getHour returns the current time rounded down to the nearest hour.
func getHour() string {
	now := time.Now().UTC().Truncate(time.Hour)
	return now.Format(time.RFC3339)
}

// getHourNumber Returns the current hour of the day
func getHourNumber() int {
	return time.Now().UTC().Hour()
}

// getMinute Returns the current minute of the hour
func getMinute() int {
	return time.Now().UTC().Minute()
}
