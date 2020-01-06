package main

import (
	. "github.com/davecremins/ToDo-Manager/cli"
	. "github.com/davecremins/ToDo-Manager/utilities"
	"os"
)

func main() {
	config := GetConfig("config.yaml")
	Process(os.Args, config)
}
