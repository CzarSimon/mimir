package main

import (
	"database/sql"
	"fmt"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
	"github.com/lib/pq"
)

// storeStatsResult stores the result of a statistics calculation in the recieveing database.
func storeStatsResult(statistics TickerStatistcs, dbConfig endpoint.SQLConfig) error {
	db, err := dbConfig.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
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
func updateStatistics(ticker string, stats Statistics, tx *sql.Tx) error {
	query := "UPDATE TWITTER_DATA SET BUSDAY_MEAN=$1, WEEKEND_MEAN=$2, BUSDAY_STDEV=$3, WEEKEND_STDEV=$4 WHERE TICKER=$5"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Print(ticker)
	fmt.Println(stats.Mean)
	_, err = stmt.Exec(
		pq.Array(stats.Mean.BusinessDays),
		pq.Array(stats.Mean.WeekendDays),
		pq.Array(stats.Stdev.BusinessDays),
		pq.Array(stats.Stdev.WeekendDays),
		ticker)
	return err
}
