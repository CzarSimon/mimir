package main

import (
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

func debug(config Config) {
	tickers, _ := GetTickers(config.TickerDB)
	logTickers(tickers)
	GetAndStorePrices(config)
}

func main() {
	config := GetConfig()
	config.LogTiming()
	debug(config)
	gocron.Every(1).Day().At(config.Timing).Do(GetAndStorePrices, config)
	<-gocron.Start()
}
