package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ToDoConfig struct {
	Filename          string `yaml:"filename"`
	SearchStr         string `yaml:"searchstr"`
	DaysToAdd         int    `yaml:"daysToAdd"`
	UseTodayForNewDay bool   `yaml:"useTodayForNewDay"`
}

func GetConfig(filename string) *ToDoConfig {
	var c *ToDoConfig
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

func SaveConfig(config ToDoConfig, filename string) error {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Marshaling error: %s", err)
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func NewDefault() *ToDoConfig {
	return &ToDoConfig{
		Filename:          "TODOs.txt",
		SearchStr:         "TODOs",
		DaysToAdd:         1,
		UseTodayForNewDay: false,
	}
}
