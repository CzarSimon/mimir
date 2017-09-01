package main

import (
	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

func main() {
	config := getConfig()
	VolumeCount(config)
	gocron.Every(1).Minute().Do(VolumeCount, config)
	gocron.Every(1).Day().At(config.Timing.TotalCount).Do(TotalCount, config.DB)
	<-gocron.Start()
}
