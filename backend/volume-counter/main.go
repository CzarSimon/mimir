package main

import "github.com/jasonlvhit/gocron"

func main() {
	config := getConfig()
	gocron.Every(1).Minute().Do(volumeCount, config)
	gocron.Every(1).Day().At(config.Timing.TotalCount).Do(totalCount, config.DB)
	<-gocron.Start()
}
