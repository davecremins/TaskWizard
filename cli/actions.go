package actions

import (
	"bufio"
	"flag"
	"fmt"
	. "github.com/davecremins/ToDo-Manager/config"
	"github.com/davecremins/ToDo-Manager/content"
	"github.com/davecremins/ToDo-Manager/dates"
	"github.com/davecremins/ToDo-Manager/display"
	"github.com/davecremins/ToDo-Manager/manager"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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

func todaysTodosActionMakeup(args []string, config *ToDoConfig) ConfigFunc {
	todaysTodosCmd := flag.NewFlagSet("today", flag.ExitOnError)
	searchStr := todaysTodosCmd.String("search", config.SearchStr, "Search string to look for")
	filename := todaysTodosCmd.String("filename", config.Filename, "Name of file to search in")
	action := func(config *ToDoConfig) {
		todaysTodosCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.Filename = *filename
		log.Println("Config over-written for today action")
		todaysTodosAction(config)
	}
	return action

}

func completeTodoActionMakeup(args []string, config *ToDoConfig) ConfigFunc {
	flag.NewFlagSet("complete", flag.ExitOnError)
	action := func(config *ToDoConfig) {
		log.Println("Config not over-written for complete action")
		completeTodoAction(config)
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

	content := manager.GetContent(config, file)
	newContent := manager.ChangeDate(config, content)
	scanner := bufio.NewScanner(strings.NewReader(newContent))
	scanner.Split(bufio.ScanLines)
	strFound := false
	readAfterFound := 0
	var take []string
	for scanner.Scan() {
		output := scanner.Text()
		strFound = strings.Contains(output, "Completed")
		take = append(take, output)
		if strFound || readAfterFound > 0 {
			readAfterFound++
		}

		if readAfterFound == 2 {
			log.Println("Previous completed todos removed successfully")
			break
		}
	}

	manager.WriteContent(file, strings.Join(take, "\n"))
	log.Println("New day todos copied successfully")
}

func newTodoAction(config *ToDoConfig, todo string) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	manager.AddNewItem(config, file, todo)
	log.Println("New todo item added successfully")
}

func todaysTodosAction(config *ToDoConfig) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	contents := manager.GetContent(config, file)

	fmt.Println("")
	display.PrintWithIndent(contents)
	fmt.Println("")
}

func completeTodoAction(config *ToDoConfig) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	contents := manager.GetContent(config, file)

	fmt.Println("")
	display.PresentItems(contents)
	option := display.AcceptInput()
	i, err := strconv.Atoi(option)
	if err != nil {
		panic(err)
	}

	organisedContent := content.NewOrganisedContent(contents)
	organisedContent.CompleteTODO(i)
	organisedContent.MergeContent()
	manager.WriteUpdatedContent(file, len(contents), organisedContent.MergedContent)
	log.Println("Updated content written to file successfully")
}
