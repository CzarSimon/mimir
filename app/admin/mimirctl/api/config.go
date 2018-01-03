package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	endpoint "github.com/CzarSimon/go-endpoint"
	"github.com/CzarSimon/util"
)

const (
	CONFIG_FOLDER = "/.mimirctl/"
	CONFIG_PATH   = CONFIG_FOLDER + "config.json"
)

// Auth Authorization credentials for the admin-api
type Auth struct {
	AccessKey string `json:"accessKey"`
}

// Config for communicating with the admin-api
type Config struct {
	Auth Auth                `json:"auth"`
	API  endpoint.ServerAddr `json:"api"`
}

// GetConfig Retrieves the api configuration from the CONFIG_PATH
func GetConfig() (Config, error) {
	rawConfig := readConfigFile()
	var config Config
	err := json.Unmarshal(rawConfig, &config)
	return config, err
}

// readConfigFile Reads the content of the api configuration file
func readConfigFile() []byte {
	configPath := prependUserDir(CONFIG_PATH)
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf(
			"Could not read config in: %s have you run mimirctl configure?\n Error: %s",
			configPath, err.Error())
		log.Fatal()
	}
	return bytes
}

// Save Saves an config struct to the CONFIG_PATH as
// read/writeable only by the current user
func (config Config) Save() {
	bytes, err := json.MarshalIndent(config, "", "    ")
	util.CheckErrFatal(err)
	createConfigDir()
	err = ioutil.WriteFile(prependUserDir(CONFIG_PATH), bytes, 0600)
	util.CheckErrFatal(err)
}

// String Returns a string representation of the configuration
func (config Config) String() string {
	return fmt.Sprintf("Host=%s Port=%s Protocol=%s",
		config.API.Host, config.API.Port, config.API.Protocol)
}

// prependUserDir Prepends a supplied path with the current users home directory
func prependUserDir(path string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir + path
}

// createConfigDir Creates configuration folder if missing
func createConfigDir() {
	err := os.Mkdir(prependUserDir(CONFIG_FOLDER), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		util.CheckErrFatal(err)
	}
}
