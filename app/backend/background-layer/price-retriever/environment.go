package main

import (
	"log"

	"github.com/CzarSimon/mimir/app/lib/go/api"
)

// Env Holds environmet information
type Env struct {
	Tickers []string
	API     api.PriceAPI
	Config  Config
}

// SetupEnv Creates a new evironmnet
func SetupEnv(config Config) *Env {
	tickers, err := GetTickers(config.TickerDB)
	if err != nil {
		log.Fatalln(err)
	}
	return &Env{
		Tickers: tickers,
		API:     api.NewIexAPI(config.Timezone),
		Config:  config,
	}
}

// UpdateTickers Queries for new tickers
func (env *Env) UpdateTickers() {
	newTickers, err := GetTickers(env.Config.TickerDB)
	if err != nil {
		log.Println(err)
		return
	}
	env.Tickers = newTickers
	logTickers(env)
}
