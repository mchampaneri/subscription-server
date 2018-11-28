package main

import (
	"encoding/json"
	"os"

	"github.com/asdine/storm"
	"github.com/fatih/color"
)

// Config is globla varialbe
// to store all configuration needed
// to be preloaded on server bootup
var Config struct {
	Port string `json:"port"`
	Db   string `json:"db"`
}

// Global DatabaseAccessError
var DatabaseAccessError error

// Global DatabaseHandle Singelton make it
// accessible through the package
var DatabaseHandle *storm.DB

// Bootup loads all configuration
// and makes server running
func Bootup() bool {

	// Loading configrutations for current
	if !loadConfig() {
		return false
	}

	// Preparing database for server
	if !preparedDB() {
		return false
	}

	return true
}

// Load config reads configuration file
// and fils " Config " [ global variable defined above ]
func loadConfig() bool {
	// Reading file - Here file has to be in same dir as our executable is ..
	configFileName := "./config.json"
	configFileHandle, configPathError := os.OpenFile(configFileName, os.O_RDONLY, os.ModePerm)
	if configPathError != nil {
		color.Red("Failed to load config file :%s", configPathError.Error())
		return false
	}
	// Readig file with Json Decoder
	configFileJSONReader := json.NewDecoder(configFileHandle)
	configDecondingError := configFileJSONReader.Decode(&Config)
	if configDecondingError != nil {
		color.Red("Failed to decodinbg  config file json : %s", configDecondingError.Error())
		return false
	}
	// Everything was fine
	// Ready to go for next step
	color.Yellow("Configurations loaded sucessfully ")
	return true
}

// PreapareDB opens database for server and
// assings it to a global singleton to make
// accessble throgh complete pacakge
func preparedDB() bool {
	// Creating database for subscription server
	DatabaseHandle, DatabaseAccessError = storm.Open(Config.Db)
	defer DatabaseHandle.Close()
	if DatabaseAccessError != nil {
		color.Red("Could not open database : %s", DatabaseAccessError.Error())
		return false
	}
	color.Yellow("Database prepared succesfully ")
	return true
}
