package main

import (
	"github.com/CzarSimon/go-file-heartbeat/heartbeat"
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

// startHeartbeat starts heartbeat emission.
func startHeartbeat(config Config) {
	RunVolumeCount(config)
	conf := config.Heartbeat
	heartbeat.RunFileHeartbeat(conf.File, conf.Interval)
}

func main() {
	config := getConfig()
	go startHeartbeat(config)
	gocron.Every(1).Minute().Do(RunVolumeCount, config)
	<-gocron.Start()
}
