package main

import (
	. "github.com/davecremins/ToDo-Manager/cli"
	. "github.com/davecremins/ToDo-Manager/config"
	"os"
)

func main() {
	config := LoadConfig()
	Process(os.Args, config)
}
