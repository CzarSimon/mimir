package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"

	r "gopkg.in/gorethink/gorethink.v2"
)

type Env struct {
	pg      *sql.DB
	rdb     *r.Session
	auth    authConfig
	devMode bool
}

func parseFlags() bool {
	var devMode bool
	flag.BoolVar(&devMode, "dev", false, "sets development mode")
	flag.Parse()
	return devMode
}

func setupEnvironment(conf Config) *Env {
	env := &Env{
		pg:      connectPostgres(conf.pg),
		rdb:     connectRethink(conf.rdb),
		auth:    conf.auth,
		devMode: conf.server.devMode,
	}
	return env
}

func main() {
	devMode := parseFlags()
	config := getConfig(devMode)

	/* ---- Environment setup ---- */
	env := setupEnvironment(config)
	defer env.pg.Close()
	defer env.rdb.Close()

	//Route handler
	mux := setupRoutes(env, config)
	server := &http.Server{
		Addr:         ":" + config.server.port,
		Handler:      mux,
		TLSConfig:    getTlsConfig(),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	/* ---- Starting Server ---- */
	log.Println("Starting server on port " + config.server.port)
	certPath := config.cert.folder + "/server.rsa.crt"
	keyPath := config.cert.folder + "/server.rsa.key"
	if !devMode {
		err := server.ListenAndServeTLS(certPath, keyPath)
		checkErr(err)
	} else {
		log.Println("Starting in dev mode")
		err := server.ListenAndServe()
		checkErr(err)
	}
}
