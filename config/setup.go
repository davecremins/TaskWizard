package config

import (
	"log"
	"os"
)

const (
	DIRECTORY_NAME = ".TaskWizard"
	CONFIG_NAME    = "TaskWizard.yaml"
	LOG_NAME       = "TaskWizard.log"
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

func fullPathToLogFile() string {
	return userDirForConfig + "/" + LOG_NAME
}

func setupLogger() {
	f, err := os.OpenFile(fullPathToLogFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkError(err)
	defer f.Close()
	log.SetOutput(f)
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
		f, err := os.Create(userDirForConfig + "/" +  defaultConfig.Filename)
		checkError(err)
		f.Close()
	}
}

func LoadConfig() *TaskConfig {
	config := GetConfig(fullPathToConfigFile())
	config.DataStore = userDirForConfig + "/" + config.Filename
	return config
}
