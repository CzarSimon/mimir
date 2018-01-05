package main

import (
	"database/sql"
	"log"

	"github.com/CzarSimon/util"
)

// TotalCount Counts the total number or tweets per ticker and stores the result
func TotalCount(dbConfig util.PGConfig) {
	log.Println("Performing total count")
	db := util.ConnectPG(dbConfig)
	defer db.Close()
	volumes, err := countTotalVolume(db)
	if err != nil {
		util.LogErr(err)
		return
	}
	err = updateTotalVolumes(volumes, db)
	if err == nil {
		log.Println("Total count done")
	} else {
		util.LogErr(err)
	}
}

// countTotalVolume Counts the total volume of tweet about all stocks in the database
func countTotalVolume(db *sql.DB) ([]Volume, error) {
	query := "SELECT ticker, COUNT(*) FROM stocktweets GROUP BY ticker"
	return GetVolume(db.Query(query))
}

// GetVolume Parses a row set from a query into a list of volumes
func GetVolume(rows *sql.Rows, err error) ([]Volume, error) {
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

// updateTotalVolumes Stores the updated count of a stocks total volume
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
