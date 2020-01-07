package actions

import (
	"flag"
	"fmt"
	. "github.com/davecremins/ToDo-Manager/config"
	"github.com/davecremins/ToDo-Manager/content"
	"github.com/davecremins/ToDo-Manager/dates"
	"github.com/davecremins/ToDo-Manager/manager"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type ConfigFunc func(*ToDoConfig)

var actionMap map[string]ConfigFunc

func Process(args []string, defaultConfig *ToDoConfig) {
	if len(args) < 2 {
		log.Fatal("expected subcommands to perform an action")
	}

	actionMap = make(map[string]ConfigFunc)

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	filename := initCmd.String("filename", defaultConfig.Filename, "Name of file to initialise")
	actionMap["init"] = func(config *ToDoConfig) {
		initCmd.Parse(args[2:])
		config.Filename = *filename
		log.Println("Config over-written for init action")
		initAction(config)
	}

	newCmd := flag.NewFlagSet("newday", flag.ExitOnError)
	searchStr := newCmd.String("search", defaultConfig.SearchStr, "Search string to look for")
	daysToAdd := newCmd.Int("days", defaultConfig.DaysToAdd, "Total amount of days to increment by")
	actionMap["newday"] = func(config *ToDoConfig) {
		newCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.DaysToAdd = *daysToAdd
		log.Println("Config over-written for newday action")
		newDayAction(config)

	}

	action, ok := actionMap[args[1]]
	if !ok {
		log.Fatal(args[1] + " subcommand is not supported right now :(")
	}
	action(defaultConfig)
}

func initAction(config *ToDoConfig) {
	initContent := content.GetInitContentWithPlaceholders()
	formattedDate := dates.ExtractShortDate(time.Now())
	filledInitContent := fmt.Sprintf(initContent, formattedDate, formattedDate)
	err := ioutil.WriteFile(config.Filename, []byte(filledInitContent), 0644)
	if err != nil {
		log.Fatal("Error ocurred writing content for init action: ", err)
	}
	log.Println(config.Filename + " created successfully with default TODOs and Completed")
}

func newDayAction(config *ToDoConfig) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	content := manager.CopyPreviousContent(config, file)
	newContent := manager.ChangeDate(config, content)
	manager.WriteContent(file, newContent)
}
