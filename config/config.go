package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config : A shared object for the service that
// provides all the essential server data
type Config struct {
	Webserver struct {
		Port string `yaml:"port"`
		Auth string `yaml:"auth"`
	} `yaml:"webserver"`

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
		Username string `yaml:"username"`
	} `yaml:"database"`

	Filetype struct {
		Contenttype []string `yaml:"contenttype"`
	} `yaml:"filetype"`
}

var Configdata Config

func Loadconfig(path string) (Config, error) {

	configfile, err := os.Open(path)
	if err != nil {
		return Configdata, err
	}
	// Close the file after its use
	defer configfile.Close()

	decoder := yaml.NewDecoder(configfile)
	err = decoder.Decode(&Configdata)

	// Print the log after successfull loading of the config variables
	log.Println("Loadconfig() - ", Configdata)
	return Configdata, err
}
