package configuration

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func InitConfig() Config {
	C := &Config{}

	// Load YAML configuration
	fmt.Println("Loading configuration from configuration.yaml")
	yml, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		fmt.Println(err)
		panic("Unable to read configuration file")
	}
	err = yaml.Unmarshal(yml, C)
	if err != nil {
		panic("Unable to read configuration file")
	}

	return *C
}
