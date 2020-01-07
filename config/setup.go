package config

import (
	"log"
	"os"
)

const (
	DIRECTORY_NAME = ".ToDo-Manager"
	CONFIG_NAME    = "ToDo-Manager.yaml"
)

func init() {
	createConfig()
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func userDirectoryForToDoConfig() string {
	home, err := os.UserHomeDir()
	checkError(err)
	return home + "/" + DIRECTORY_NAME
}

func createConfig() {
	userDirectoryForToDo := userDirectoryForToDoConfig()
	if _, err := os.Stat(userDirectoryForToDo); os.IsNotExist(err) {
		log.Println("Required directory missing")
		err := os.Mkdir(userDirectoryForToDo, 0755)
		checkError(err)
		log.Println(userDirectoryForToDo + " directory created successfully")
		defaultConfig := NewDefault()
		err = SaveConfig(*defaultConfig, userDirectoryForToDo+"/"+CONFIG_NAME)
		checkError(err)
		log.Println(CONFIG_NAME + " configuration created successfully")
	} else {
		log.Println("Required directory already exists")
	}
}

func LoadConfig() *ToDoConfig {
	userDirectoryForToDo := userDirectoryForToDoConfig()
	config := GetConfig(userDirectoryForToDo + "/" + CONFIG_NAME)
	return config
}
