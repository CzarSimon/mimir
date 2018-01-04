package main

import (
	"fmt"

	"github.com/CzarSimon/mimir/app/admin/mimirctl/api"
	"github.com/CzarSimon/util"
)

func getConfig() api.Config {
	config, err := api.GetConfig()
	util.CheckErrFatal(err)
	return config
}

func test() {
	config := getConfig()
	fmt.Println(config)
}
