package main

import (
	"log"
)

func checkErrFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
