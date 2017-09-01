package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/util"
)

// VolumeCount Counts the tweet volume in the last hour
func VolumeCount(config Config) {
	log.Println("Running volume count")
	db := util.ConnectPG(config.DB)
	defer db.Close()
	tickers, err := getAllTickers(db)
	if err != nil {
		util.LogErr(err)
		return
	}
	err = countHourlyVolumes(tickers, db)
	if err != nil {
		util.LogErr(err)
		return
	}
	err = sendResult(tickers, config.Server)
	util.CheckErr(err)
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
	var count int
	for rows.Next() {
		err = rows.Scan(&ticker, &count)
		if err != nil {
			return err
		}
		volumes[ticker].Volume = count
	}
	return nil
}

// sendResult Sends resulting volumes to revciving server
func sendResult(volumes HourVolumes, server util.ServerConfig) error {
	jsonStr, err := json.Marshal(volumes)
	if err != nil {
		return err
	}
	log.Println(string(jsonStr))
	req, err := http.NewRequest("POST", buildURL(server), bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("Error code %d in response", resp.StatusCode)
}

// buildURL Returns the url to post volume stats to
func buildURL(server util.ServerConfig) string {
	fmt.Println(server.ToURL("api/app/twitter-data/volumes"))
	return server.ToURL("api/app/twitter-data/volumes")
}

// HourVolume conatains this hours count and the minute the count occured
type HourVolume struct {
	Volume int `json:"volume"`
	Minute int `json:"minute"`
}

// HourVolumes is a ticker -> HourVolume map
type HourVolumes map[string]*HourVolume

// getAllTickers Gets the ticker of each tracked stock and returns an HourVolumes
// map with the current minute and no recorded volume per ticker
func getAllTickers(db *sql.DB) (HourVolumes, error) {
	tickers := make(HourVolumes)
	rows, err := db.Query("SELECT ticker FROM stocks")
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

// getHour Returns the current hour
func getHour() string {
	now := time.Now().UTC().Truncate(time.Hour)
	return now.Format(time.RFC3339)
}

// getMinute Returns the current minute of the hour
func getMinute() int {
	return time.Now().UTC().Minute()
}
