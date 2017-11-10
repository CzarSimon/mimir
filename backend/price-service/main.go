package main

import (
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

func debug(env *Env) {
	//GetAndStoreClosePrices(config)
	env.GetAndStoreLatestPrices()
}

func main() {
	config := GetConfig()
	config.LogTiming()
	env := NewEnv(config)
	debug(env)
	gocron.Every(1).Day().At(config.Timing.ClosePriceTime).Do(env.GetAndStoreClosePrices)
	gocron.Every(1).Minute().Do(env.GetAndStoreLatestPrices)
	gocron.Every(30).Minutes().Do(env.UpdateTickers)
	<-gocron.Start()
}
