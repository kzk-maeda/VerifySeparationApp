package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var Config Configuration
var logger *log.Logger

func init() {
	loadConfig()
	file, err := os.OpenFile("goweb_v1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("Cannot get configuration from file", err)
	}
}

// version
func Version() string {
	return "v1"
}
