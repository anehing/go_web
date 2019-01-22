package main

import (
	"os"
	"encoding/json"
	"log"
)

type Configuration struct {
	Address        string
	ReadTimeout    int32
	WriteTimeout   int32
	Static       string
}

var config Configuration


func init() {
	loadConfig()
}

func loadConfig(){
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

