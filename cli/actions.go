package actions

import (
	"flag"
	"github.com/davecremins/ToDo-Manager/manager"
	. "github.com/davecremins/ToDo-Manager/utilities"
	"log"
	"os"
)

var actionMap map[string]func(*Config)

func Process(args []string, defaultConfig *Config) {
	if len(args) < 2 {
		log.Fatal("expected subcommands to perform an action")
	}

	actionMap = make(map[string]func(*Config))

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	filename := initCmd.String("filename", defaultConfig.Filename, "Name of file to initialise")
	actionMap["init"] = func(config *Config) {
		initCmd.Parse(args[2:])
		defaultConfig.Filename = *filename
		log.Println("Config over-written for init action")
		log.Println("No implementation provided")
	}

	newCmd := flag.NewFlagSet("newday", flag.ExitOnError)
	searchStr := newCmd.String("search", defaultConfig.SearchStr, "Search string to look for")
	daysToAdd := newCmd.Int("days", defaultConfig.DaysToAdd, "Total amount of days to increment by")
	actionMap["newday"] = func(config *Config) {
		newCmd.Parse(args[2:])
		defaultConfig.SearchStr = *searchStr
		defaultConfig.DaysToAdd = *daysToAdd
		log.Println("Config over-written for newday action")
		newDayAction(defaultConfig)

	}

	action, ok := actionMap[args[1]]
	if !ok {
		log.Fatal(args[1] + " subcommand is not supported right now :(")
	}
	action(defaultConfig)
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
