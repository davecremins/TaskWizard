package main

import (
	"github.com/davecremins/ToDo-Manager/manager"
	. "github.com/davecremins/ToDo-Manager/utilities"
	"log"
	"os"
)

func main() {
	config := GetConfig("config.yaml")
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	manager.CopyPreviousContent(config, file)
}
