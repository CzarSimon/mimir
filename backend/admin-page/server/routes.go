package main

import (
  "net/http"
)

func setupRoutes(env *Env, config Config) *http.ServeMux {
  mux := http.NewServeMux()
  mux.Handle("/", http.FileServer(http.Dir(config.server.staticFolder)))
  mux.HandleFunc("/api/login", env.login)
  mux.HandleFunc("/api/tracked-stocks", env.sendStockInfo)
  mux.HandleFunc("/api/untrack-stock", env.untrackStock)
  mux.HandleFunc("/api/untracked-tickers", env.sendTickers)
  mux.HandleFunc("/api/track-ticker", env.trackTicker)
  mux.HandleFunc("/api/update-stock-info", env.updateStockInfo)
  return mux
}
