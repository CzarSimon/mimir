package main

import (
	"log"

	"github.com/CzarSimon/util"
)

func main() {
	log.Println("Retriving close prices")
	config := GetConfig()
	env := SetupEnv(config)
	defer env.PriceDB.Close()

	logTickers(env)
	err := env.GetAndStoreClosePrices()
	util.CheckErr(err)
}
