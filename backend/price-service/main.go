package main

import (
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

func debug(config Config) {
	//GetAndStoreClosePrices(config)
	GetAndStoreLatestPrices(config)
}

func main() {
	config := GetConfig()
	config.LogTiming()
	debug(config)
	gocron.Every(1).Day().At(
		config.Timing.ClosePriceTime).Do(GetAndStoreClosePrices, config)
	gocron.Every(1).Minute().Do(GetAndStoreLatestPrices, config)
	<-gocron.Start()
}
