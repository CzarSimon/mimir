package main

import (
	"github.com/CzarSimon/util"
	_ "github.com/lib/pq"
)

func main() {
	config := GetConfig()
	db, err := config.DB.Connect()
	util.CheckErrFatal(err)
	defer db.Close()

	err = CalculateAndStoreTotalVolumes(db)
	util.CheckErrFatal(err)
}
