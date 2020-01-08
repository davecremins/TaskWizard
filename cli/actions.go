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

func initActionMakeup(args []string, config *ToDoConfig) ConfigFunc {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	filename := initCmd.String("filename", config.Filename, "Name of file to initialise")
	action := func(config *ToDoConfig) {
		initCmd.Parse(args[2:])
		config.Filename = *filename
		log.Println("Config over-written for init action")
		initAction(config)
	}
	return action
}

func newDayActionMakeup(args []string, config *ToDoConfig) ConfigFunc {
	newDayCmd := flag.NewFlagSet("newday", flag.ExitOnError)
	searchStr := newDayCmd.String("search", config.SearchStr, "Search string to look for")
	daysToAdd := newDayCmd.Int("days", config.DaysToAdd, "Total amount of days to increment by")
	filename := newDayCmd.String("filename", config.Filename, "Name of file to add new day to")
	action := func(config *ToDoConfig) {
		newDayCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.DaysToAdd = *daysToAdd
		config.Filename = *filename
		log.Println("Config over-written for newday action")
		newDayAction(config)
	}
	return action
}

func newTodoActionMakeup(args []string, config *ToDoConfig) ConfigFunc {
	newTodoCmd := flag.NewFlagSet("newtodo", flag.ExitOnError)
	searchStr := newTodoCmd.String("search", config.SearchStr, "Search string to look for")
	filename := newTodoCmd.String("filename", config.Filename, "Name of file to add new todo")
	todo := newTodoCmd.String("desc", "New todo item placeholder", "Description of new todo")
	action := func(config *ToDoConfig) {
		newTodoCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.Filename = *filename
		log.Println("Config over-written for newtodo action")
		newTodoAction(config, *todo)
	}
	return action
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

func newTodoAction(config *ToDoConfig, todo string) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	log.Println("Adding new item")
	manager.AddNewItem(config, file, todo)
	log.Println("Added new item")
}
