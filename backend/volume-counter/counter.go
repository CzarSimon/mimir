package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func volumeCount(config Config) {
	log.Println("Running volume count")
	db := connectPostgres(config.DB)
	defer db.Close()
	tickers, err := getAllTickers(db)
	checkErrFatal(err)
	err = countHourlyVolumes(tickers, db)
	checkErrFatal(err)
	err = sendResult(tickers, config.Server)
	checkErr(err)
}

func totalCount(config DBConfig) {
	log.Println("Performing total count")
	db := connectPostgres(config)
	defer db.Close()
	volumes, err := countTotalVolume(db)
	checkErrFatal(err)
	err = updateTotalVolumes(volumes, db)
	if err == nil {
		log.Println("Total count done")
	} else {
		log.Println(err.Error())
	}
}

//Volume type
type Volume struct {
	Ticker string
	Count  int64
}

func getVolume(rows *sql.Rows, err error) ([]Volume, error) {
	volumes := make([]Volume, 0)
	defer rows.Close()
	if err != nil {
		return volumes, err
	}
	var row Volume
	for rows.Next() {
		err := rows.Scan(&row.Ticker, &row.Count)
		if err != nil {
			return volumes, err
		}
		volumes = append(volumes, row)
	}
	return volumes, nil
}

func countTotalVolume(db *sql.DB) ([]Volume, error) {
	query := "SELECT ticker, COUNT(*) FROM stocktweets GROUP BY ticker"
	return getVolume(db.Query(query))
}

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

func getHour() string {
	now := time.Now().UTC().Truncate(time.Hour)
	return now.Format(time.RFC3339)
}

func updateTotalVolumes(volumes []Volume, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	stmt, err := tx.Prepare("UPDATE stocks SET total_count=$1 WHERE ticker=$2")
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, volume := range volumes {
		_, err = stmt.Exec(volume.Count, volume.Ticker)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func sendResult(volumes HourVolumes, server ServerConfig) error {
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
	return errors.New("Non 200 response")
}

func buildURL(server ServerConfig) string {
	return fmt.Sprintf("http://%s:%s/api/app/twitter-data/volumes", server.IP, server.Port)
}

//HourVolume conatains this hours count an the minute the count occured
type HourVolume struct {
	Volume int `json:"volume"`
	Minute int `json:"minute"`
}

//HourVolumes is a ticker -> HourVolume map
type HourVolumes map[string]*HourVolume

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

func getMinute() int {
	return time.Now().UTC().Minute()
}
