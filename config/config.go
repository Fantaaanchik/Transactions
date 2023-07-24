package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type App struct {
	ServerName string `json:"server_name"`
	PortRun    string `json:"port_run"`
	DB         string `json:"db"`
}

var Configure App

func ReadConfig(F string) {
	byteValue, err := ioutil.ReadFile(F)
	if err != nil {
		log.Fatalf("Error read json File, %v", err)
		return
	}
	err = json.Unmarshal(byteValue, &Configure)
	if err != nil {
		log.Fatal("Error unmarshal json File, err : ", err.Error())
		return
	}
}
