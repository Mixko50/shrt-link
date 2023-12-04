package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var C = new(config)

func InitConfig() {
	// Load YAML configuration
	fmt.Println("Loading configuration from config.yaml")
	yml, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		panic("Unable to read configuration file")
	}
	err = yaml.Unmarshal(yml, C)
	if err != nil {
		panic("Unable to read configuration file")
	}
}
