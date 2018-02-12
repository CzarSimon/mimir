package main

import "log"

func main() {
	config := GetConfig()
	err := StatsCalc(config)
	if err != nil {
		log.Fatal(err)
	}
}
