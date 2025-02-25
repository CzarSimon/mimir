package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/query"
	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

// TwitterData Holds information about ticker voluems and statistics
type TwitterData struct {
	Ticker string     `json:"ticker"`
	Minute int16      `json:"minute"`
	Hour   int        `json:"hour"`
	Volume Volumes    `json:"volume"`
	Mean   Statistics `json:"mean"`
	Stdev  Statistics `json:"stdev"`
}

// NewTwitterData Creates a new TwitterDat struct
func NewTwitterData() TwitterData {
	return TwitterData{
		Hour: time.Now().UTC().Hour(),
	}
}

// Volumes Slice containing recorded tweet volumes for the latest 24 hours
type Volumes []int64

// Scan Reads array from database into volume slice
func (volumes *Volumes) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return errors.New("Could not scan volumes")
	}
	vol, err := util.BytesToIntSlice(bytes, 64)
	if err != nil {
		return errors.New("Could not scan volumes")
	}
	(*volumes) = vol
	return nil
}

// DayTypeStatistics Struct containing statistics separated by daytype
type DayTypeStatistics struct {
	Busdays     Statistics `json:"busdays"`
	WeekendDays Statistics `json:"weekendDays"`
}

// Statistics Slice containing statistics about tweet volumes for all hours of the day
type Statistics []float64

// Scan Reads array from database into statistics slice
func (statistics *Statistics) Scan(source interface{}) error {
	bytes, ok := source.([]byte)
	if !ok {
		return errors.New("Could not scan statistics")
	}
	stats, err := util.BytesToFloatSlice(bytes, 64)
	if err != nil {
		return errors.New("Could not scan statistics")
	}
	(*statistics) = stats
	return nil
}

// HandleGetTwitterDataRequest Handles GET request for twitter data for a supplied set of tickers
func (env *Env) HandleGetTwitterDataRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httputil.MethodNotAllowed
	}
	tickers, err := query.ParseValues(r, TICKER_KEY)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	twitterData, err := getTwitterData(tickers, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	return httputil.SendJSON(w, twitterData)
}

// getTwitterData Retreives twitter data for a supplied set of tickers from the database
func getTwitterData(tickers []string, db *sql.DB) ([]TwitterData, error) {
	twitterData := make([]TwitterData, 0)
	rows, err := db.Query(getTwitterDataQuery(), pq.Array(tickers))
	defer rows.Close()
	if err != nil {
		return twitterData, err
	}
	td := NewTwitterData()
	for rows.Next() {
		err = rows.Scan(&td.Ticker, &td.Minute, &td.Volume, &td.Mean, &td.Stdev)
		if err != nil {
			return twitterData, err
		}
		twitterData = append(twitterData, td)
	}
	return twitterData, nil
}

// getTwitterDataQuery Gets a query to retrive twitter data
func getTwitterDataQuery() string {
	query := "SELECT TICKER, MINUTE, VOLUME, %s FROM TWITTER_DATA WHERE TICKER = ANY($1)"
	if !util.IsWeekend() {
		return fmt.Sprintf(query, "BUSDAY_MEAN, BUSDAY_STDEV")
	}
	return fmt.Sprintf(query, "WEEKEND_MEAN, WEEKEND_STDEV")
}

// Volume Holds info about volumes for a given minute
type Volume struct {
	Minute int64 `json:"minute"`
	Volume int64 `json:"volume"`
}

// HourVolumes Holds volumes for all stocks for a given hour
type HourVolumes struct {
	Hour    int               `json:"hour"`
	Volumes map[string]Volume `json:"volumes"`
}

// HandleNewVolumesRequest Handles posting of new ticker volumes
func (env *Env) HandleNewVolumesRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return httputil.MethodNotAllowed
	}
	var volumes HourVolumes
	err := util.DecodeJSON(r.Body, &volumes)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	err = storeVolumes(volumes, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendOK(w)
	return nil
}

// storeVolumes Stores volumes in the database
func storeVolumes(volumes HourVolumes, db *sql.DB) error {
	tx, err := db.Begin()
	if util.CheckErrAndRollback(err, tx) {
		return err
	}
	for ticker, volume := range volumes.Volumes {
		err = updateTickerVolume(ticker, volumes.Hour, volume, tx)
		if util.CheckErrAndRollback(err, tx) {
			return err
		}
	}
	return tx.Commit()
}

// updateTickerVolume Updates the voluems for a supplied ticker
func updateTickerVolume(ticker string, hour int, volume Volume, tx *sql.Tx) error {
	volumes, err := getTickerVolumes(ticker, tx)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	volumes[hour] = volume.Volume
	stmt, err := tx.Prepare("UPDATE TWITTER_DATA SET VOLUME=$1, MINUTE=$2 WHERE TICKER=$3")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(pq.Array(volumes), volume.Minute, ticker)
	return err
}

// getTickerVolumes Gets the currently stored volumes for a ticker
func getTickerVolumes(ticker string, tx *sql.Tx) (Volumes, error) {
	var volumes Volumes
	err := tx.QueryRow("SELECT VOLUME FROM TWITTER_DATA WHERE TICKER=$1", ticker).Scan(&volumes)
	return volumes, err
}

// VolumeStatistics Holds statistics about tweet volumes
type VolumeStatistics struct {
	Mean  DayTypeStatistics `json:"mean"`
	Stdev DayTypeStatistics `json:"stdev"`
}

// HandleNewStatsRequest Handles postings of new volumes statistics
func (env *Env) HandleNewStatsRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return httputil.MethodNotAllowed
	}
	statistics := make(map[string]VolumeStatistics)
	err := util.DecodeJSON(r.Body, &statistics)
	if err != nil {
		log.Println(err)
		return httputil.BadRequest
	}
	err = storeStatistics(statistics, env.db)
	if err != nil {
		log.Println(err)
		return httputil.InternalServerError
	}
	httputil.SendOK(w)
	return nil
}

// storeStatistics Stores updated volume statistics in database
func storeStatistics(statistics map[string]VolumeStatistics, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for ticker, stats := range statistics {
		err = updateStatistics(ticker, stats, tx)
		if util.CheckErrAndRollback(err, tx) {
			return err
		}
	}
	return tx.Commit()
}

// updateStatistics Stores updated statistics for a given ticker in a database
func updateStatistics(ticker string, statistics VolumeStatistics, tx *sql.Tx) error {
	query := "UPDATE TWITTER_DATA SET BUSDAY_MEAN=$1, WEEKEND_MEAN=$2, BUSDAY_STDEV=$3, WEEKEND_STDEV=$4 WHERE TICKER=$5"
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		pq.Array(statistics.Mean.Busdays),
		pq.Array(statistics.Mean.WeekendDays),
		pq.Array(statistics.Stdev.Busdays),
		pq.Array(statistics.Stdev.WeekendDays),
		ticker)
	return err
}
