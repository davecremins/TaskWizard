package utilities

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Filename  string `yaml:"filename"`
	SearchStr string `yaml:"searchstr"`
}

func GetConfig(filename string) *Config {
	var c *Config
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshaling error: %s", err)
	}

	return c
}
