package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type appSettings struct {
	Host          string `json:"host"`
	Filename      string `json:"filename"`
	User          string `json:"user"`
	Pass          string `json:"pass"`
	NoOfRequest   int    `json:"no_of_request"`
	NoOfConcurret int    `json:"no_of_conccurent"`
}

const filenameSettings = "settings.json"

var settings appSettings

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := filepath.Split(ex)
	dat, err := ioutil.ReadFile(dir + filenameSettings)
	if err != nil {
		data, _ := json.Marshal(settings)
		ioutil.WriteFile(dir+filenameSettings, data, 0664)
		log.Fatal("settings.json missing, " + err.Error())
	}

	if err := json.Unmarshal(dat, &settings); err != nil {
		log.Fatal(err)
	}
}
