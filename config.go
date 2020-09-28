package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// The type of setting we support in our json file
type appSettings struct {
	Host          string `json:"host"`
	Filename      string `json:"filename"`
	User          string `json:"user"`
	Pass          string `json:"pass"`
	NoOfRequest   int    `json:"no_of_request"`
	NoOfConcurret int    `json:"no_of_conccurent"`
	ShowDone      bool   `json:"show_done"`
}

// The filename of our settings file
const filenameSettings = "settings.json"

// Global variable to our settings
var settings appSettings

// All init functions runs before our main func
func init() {
	// Get the path and name to where our program is
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// Get the directory from that
	dir, _ := filepath.Split(ex)
	// Read our settings file
	dat, err := ioutil.ReadFile(dir + filenameSettings)
	if err != nil {
		// Wow, we did not have a settings.json file; so we create an empty one
		data, _ := json.Marshal(settings)
		ioutil.WriteFile(dir+filenameSettings, data, 0664)
		log.Fatal("settings.json missing, " + err.Error())
	}

	// Read the settings from our json file into our struct
	if err := json.Unmarshal(dat, &settings); err != nil {
		log.Fatal(err)
	}
}
