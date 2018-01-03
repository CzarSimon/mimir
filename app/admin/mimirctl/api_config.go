package main

import (
  "fmt"
  "ioutil"
  "json"

  endpoint "github.com/CzarSimon/go-endpoint"
  "github.com/CzarSimon/util"
)

const CONFIG_PATH = "~/.mimirctl/config.json"

type Auth struct {
  AccessKey string `json:"accessKey"`
}

type ApiConfig struct {
  Auth Auth                `json:"auth"`
  API  endpoint.ServerAddr `json:"api"`
}

func getApiConfig() (ApiConfig, error) {
  rawCofig := readConfigFile()
  var config ApiConfig
  err := json.Unmarshall(rawConfig, &ApiConfig)
  return config, err
}

func readConfigFile() []byte {
  bytes, err := ioutil.ReadFile(CONFIG_PATH)
  if err != nil {
    fmt.Fatalf(
      "Could not read config in: %s have you run mimirctl configure?\n Error: %s",
      CONFIG_PATH, err.Error())
  }
  return bytes
}

func (config ApiConfig) Save() {
  bytes, err := json.MarshallIndent(config, "", "\t")
  util.CheckErrFatal(err)
  err := ioutil.WriteFile(CONFIG_PATH, bytes, 0600)
  util.CheckErrFatal(err)
}
