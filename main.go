package main

import (
//	"github.com/davecremins/ToDo-Manager/manager"
	. "github.com/davecremins/ToDo-Manager/utilities"
	. "github.com/davecremins/ToDo-Manager/cli"
	"log"
	"os"
)

func main() {
	config := GetConfig("config.yaml")
	err := Process(os.Args, config)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	/*file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	content := manager.CopyPreviousContent(config, file)
	newContent := manager.ChangeDate(config, content)
	manager.WriteContent(file, newContent)
	*/
}
