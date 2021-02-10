package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type TaskConfig struct {
	Filename string `yaml:"filename"`
	DataStore string
}

func GetConfig(filename string) *TaskConfig {
	var c *TaskConfig
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

func SaveConfig(config TaskConfig, filename string) error {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Marshaling error: %s", err)
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func NewDefault() *TaskConfig {
	return &TaskConfig{
		Filename: "taskWizardData.json",
	}
}
