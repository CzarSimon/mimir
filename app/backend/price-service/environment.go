package main

import (
	"log"
	"sync"

	"github.com/CzarSimon/util"
)

// Env Holds environmet information
type Env struct {
	Tickers    []string
	TickerLock sync.Mutex
	API        PriceAPI
	Config     Config
}

// NewEnv Creates a new evironmnet
func NewEnv(config Config) *Env {
	tickers, err := GetTickers(config.TickerDB)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &Env{
		Tickers:    tickers,
		TickerLock: sync.Mutex{},
		API:        selectPriceAPI(config.Timezone),
		Config:     config,
	}
}

// getTickers Thread safe retrival of tickers
func (env *Env) getTickers() []string {
	env.TickerLock.Lock()
	defer env.TickerLock.Unlock()
	return env.Tickers
}

// UpdateTickers Queries for new tickers
func (env *Env) UpdateTickers() {
	env.TickerLock.Lock()
	newTickers, err := GetTickers(env.Config.TickerDB)
	if err == nil {
		env.Tickers = newTickers
		log.Printf("Updated tickers. %d tickers now tracked\n", len(env.Tickers))
	} else {
		util.LogErr(err)
	}
	env.TickerLock.Unlock()
}
