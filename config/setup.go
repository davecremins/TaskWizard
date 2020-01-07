package config

import (
	"log"
	"os"
)

const (
	DIRECTORY_NAME = ".ToDo-Manager"
	CONFIG_NAME    = "ToDo-Manager.yaml"
)

var userDirForConfig string

func init() {
	home, err := os.UserHomeDir()
	checkError(err)
	userDirForConfig = home + "/" + DIRECTORY_NAME
	createConfig()
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func fullPathToConfigFile() string {
	return userDirForConfig + "/" + CONFIG_NAME
}

func createConfig() {
	if _, err := os.Stat(userDirForConfig); os.IsNotExist(err) {
		log.Println("Required directory missing")
		err := os.Mkdir(userDirForConfig, 0755)
		checkError(err)
		log.Println(userDirForConfig + " directory created successfully")
		defaultConfig := NewDefault()
		err = SaveConfig(*defaultConfig, fullPathToConfigFile())
		checkError(err)
		log.Println(CONFIG_NAME + " configuration created successfully")
	} else {
		log.Println("Required directory already exists")
	}
}

func LoadConfig() *ToDoConfig {
	config := GetConfig(fullPathToConfigFile())
	return config
}
