package main

import (
	"brankas_test/config"
	"log"
	"net/http"
)

func main() {

	// Constants for the configuration
	const (
		configpath string = "env/config.yaml"
	)

	// Load configuration parameters
	var configdata config.Config
	configdata, err := config.Loadconfig(configpath)
	if err != nil {
		log.Fatalln("Loadconfig() - Unable to open the configuration file. | ", err)
	}

	// Start the server
	log.Fatalln(http.ListenAndServe(":"+configdata.Webserver.Port, nil))
}
