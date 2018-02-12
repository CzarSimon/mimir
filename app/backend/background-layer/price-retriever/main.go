package main

import (
	"github.com/CzarSimon/go-file-heartbeat/heartbeat"
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

// startHeartbeat starts heartbeat emission.
func startHeartbeat(env *Env) {
	conf := env.Config.Heartbeat
	env.GetAndStoreLatestPrices()
	heartbeat.RunFileHeartbeat(conf.File, conf.Interval)
}

func main() {
	config := GetConfig()
	env := SetupEnv(config)
	go startHeartbeat(env)

	gocron.Every(1).Minute().Do(env.GetAndStoreLatestPrices)
	gocron.Every(30).Minutes().Do(env.UpdateTickers)
	<-gocron.Start()
}
