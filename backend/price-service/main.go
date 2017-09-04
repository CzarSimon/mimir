package main

import (
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

func main() {
	config := GetConfig()
	gocron.Every(1).Day().At(config.Timing).Do(GetAndStorePrices, config)
	<-gocron.Start()
}
