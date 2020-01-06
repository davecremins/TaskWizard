package actions

import (
	"flag"
	"fmt"
	"github.com/davecremins/ToDo-Manager/manager"
	. "github.com/davecremins/ToDo-Manager/utilities"
	"log"
	"os"
)

func Process(args []string, defaultConfig *Config) {
	if len(args) < 2 {
		log.Fatal("expected subcommands to perform an action")
	}

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	filename := initCmd.String("filename", defaultConfig.Filename, "Name of file to initialise")

	newCmd := flag.NewFlagSet("newday", flag.ExitOnError)
	searchStr := newCmd.String("search", defaultConfig.SearchStr, "Search string to look for")
	daysToAdd := newCmd.Int("days", defaultConfig.DaysToAdd, "Total amount of days to increment by")

	switch args[1] {

	case "init":
		initCmd.Parse(args[2:])
		fmt.Println("subcommand 'init'")
		fmt.Println("  filename:", *filename)
		fmt.Println("  tail:", initCmd.Args())
		defaultConfig.Filename = *filename
		log.Println("Config over-written for init action")
	case "newday":
		newCmd.Parse(args[2:])
		fmt.Println("subcommand 'newday'")
		fmt.Println("  searchStr:", *searchStr)
		fmt.Println("  days:", *daysToAdd)
		fmt.Println("  tail:", newCmd.Args())
		defaultConfig.SearchStr = *searchStr
		defaultConfig.DaysToAdd = *daysToAdd
		log.Println("Config over-written for newday action")
		newDayAction(defaultConfig)
	default:
		log.Fatal(args[1] + " subcommand is not supported right now :(")
	}
}

func newDayAction(config *Config) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	content := manager.CopyPreviousContent(config, file)
	newContent := manager.ChangeDate(config, content)
	manager.WriteContent(file, newContent)
}
