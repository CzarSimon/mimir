package main

import (
	"database/sql"
	"fmt"
	"log"
)

// Volume ticker and volume for a specific stock.
type Volume struct {
	Ticker string
	Count  int64
}

// String returns a string representation of a volume.
func (volume Volume) String() string {
	return fmt.Sprintf("Ticker=%s Count=%d", volume.Ticker, volume.Count)
}

// CalculateAndStoreTotalVolumes calculates total tweet volumes for all
// tracked tickers and stores the result.
func CalculateAndStoreTotalVolumes(db *sql.DB) error {
	volumes, err := getVolumes(db)
	if err != nil {
		return err
	}
	return storeVolumes(volumes, db)
}

// getVolumes gets total volumes for all tracked tickers.
func getVolumes(db *sql.DB) ([]Volume, error) {
	rows, err := db.Query("SELECT ticker, COUNT(*) FROM stocktweets GROUP BY ticker")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return createVolumeList(rows)
}

// createVolumeList creates a list of volumes.
func createVolumeList(rows *sql.Rows) ([]Volume, error) {
	volumes := make([]Volume, 0)
	var volume Volume
	for rows.Next() {
		err := rows.Scan(&volume.Ticker, &volume.Count)
		if err != nil {
			return nil, err
		}
		volumes = append(volumes, volume)
	}
	return volumes, nil
}

// storeVolumes stores the calculated volumes in the database.
func storeVolumes(volumes []Volume, db *sql.DB) error {
	var err error
	for _, volume := range volumes {
		log.Println(volume)
		if e := storeVolume(volume, db); e != nil {
			err = e
		}
	}
	return err
}

// storeVolume stores a calculated ticker volume in the database.
func storeVolume(volume Volume, db *sql.DB) error {
	stmt, err := db.Prepare("UPDATE stocks SET total_count=$1 WHERE ticker=$2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(volume.Count, volume.Ticker)
	return err
}
