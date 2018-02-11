package main

import (
	"database/sql"
	"errors"

	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

// storeVolumes sends resulting volumes to revciving server.
func storeVolumes(volumes VolumeResult, config Config) error {
	db, err := config.AppDB.Connect()
	util.CheckErrFatal(err)
	defer db.Close()
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
func updateTickerVolume(ticker string, hour int, volume *HourVolume, tx *sql.Tx) error {
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
