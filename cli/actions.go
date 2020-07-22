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

type ConfigFunc func([]string)

var FlagDefaults []func()

func addFlagSetDefault(f func()) {
	FlagDefaults = append(FlagDefaults, f)
}

func printDefaults() {
	fmt.Println("")
	for _, f := range FlagDefaults {
		f()
		fmt.Println("")
	}
}

func initActionMakeup(config *ToDoConfig) ConfigFunc {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	filename := initCmd.String("filename", config.Filename, "Name of file to initialise")
	action := func(args []string) {
		initCmd.Parse(args[2:])
		config.Filename = *filename
		log.Println("Config over-written for init action")
		initAction(config)
	}
	addFlagSetDefault(initCmd.Usage)
	return action
}

func newDayActionMakeup(config *ToDoConfig) ConfigFunc {
	newDayCmd := flag.NewFlagSet("newday", flag.ExitOnError)
	searchStr := newDayCmd.String("search", config.SearchStr, "Search string to look for")
	daysToAdd := newDayCmd.Int("days", config.DaysToAdd, "Total amount of days to increment by")
	filename := newDayCmd.String("filename", config.Filename, "Name of file to add new day to")
	action := func(args []string) {
		newDayCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.DaysToAdd = *daysToAdd
		config.Filename = *filename
		log.Println("Config over-written for newday action")
		newDayAction(config)
	}
	addFlagSetDefault(newDayCmd.Usage)
	return action
}

func newTodoActionMakeup(config *ToDoConfig) ConfigFunc {
	newTodoCmd := flag.NewFlagSet("newtodo", flag.ExitOnError)
	searchStr := newTodoCmd.String("search", config.SearchStr, "Search string to look for")
	filename := newTodoCmd.String("filename", config.Filename, "Name of file to add new todo")
	todo := newTodoCmd.String("desc", "New todo item placeholder", "Description of new todo")
	action := func(args []string) {
		newTodoCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.Filename = *filename
		log.Println("Config over-written for newtodo action")
		newTodoAction(config, *todo)
	}
	addFlagSetDefault(newTodoCmd.Usage)
	return action
}

func todaysTodosActionMakeup(config *ToDoConfig) ConfigFunc {
	todaysTodosCmd := flag.NewFlagSet("today", flag.ExitOnError)
	searchStr := todaysTodosCmd.String("search", config.SearchStr, "Search string to look for")
	filename := todaysTodosCmd.String("filename", config.Filename, "Name of file to search in")
	action := func(args []string) {
		todaysTodosCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.Filename = *filename
		log.Println("Config over-written for today action")
		todaysTodosAction(config)
	}
	addFlagSetDefault(todaysTodosCmd.Usage)
	return action

}

func completeTodoActionMakeup(config *ToDoConfig) ConfigFunc {
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	includeEdit := completeCmd.Bool("edit", false, "Option to edit todo item before completion")
	action := func(args []string) {
		completeCmd.Parse(args[2:])
		log.Println("Config not over-written for complete action")
		log.Println("Specific complete cmd flag included", *includeEdit)
		completeTodoAction(config, *includeEdit)
	}
	addFlagSetDefault(completeCmd.Usage)
	return action
}

func moveTodoActionMakeup(config *ToDoConfig) ConfigFunc {
	moveCmd := flag.NewFlagSet("move", flag.ExitOnError)
	action := func(args []string) {
		log.Println("Config not over-written for move action")
		moveTodoAction(config)
	}
	addFlagSetDefault(moveCmd.Usage)
	return action
}

func mergeTodoActionMakeup(config *ToDoConfig) ConfigFunc {
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	action := func(args []string) {
		log.Println("Config not over-written for move action")
		mergeTodoAction(config)
	}
	addFlagSetDefault(mergeCmd.Usage)
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

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, config.SearchStr)
	dateStr, err := dates.FindDate(contentContainingStr)
	if err != nil {
		panic("Failed to find date in content")
	}

	datetime, err := dates.ConvertToTime(dateStr)
	if err != nil {
		panic("Failed to convert date to time format")
	}

	datetime = dates.AddDays(datetime, config.DaysToAdd)
	newDateStr := dates.ExtractShortDate(datetime)
	newContent := strings.ReplaceAll(contentContainingStr, dateStr, newDateStr)
	log.Println("Content updated with new date")

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

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, "Completed")
	contentSize := len(contentContainingStr)
	log.Println("Position found:", contentSize)
	// Account for newline
	contentSize += 1

	writingPos := size - int64(contentSize)
	file.Seek(writingPos, 0)
	_, err = file.Write([]byte(todo))
	if err != nil {
		panic("Falied to write new item to file")
	}

	_, err = file.Write([]byte("\n\n"))
	if err != nil {
		panic("Falied to write new line to file")
	}

	_, err = file.Write([]byte(contentContainingStr))
	if err != nil {
		panic("Falied to write original content to file")
	}
	log.Println("New todo item added successfully")
}

func todaysTodosAction(config *ToDoConfig) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, config.SearchStr)
	organisedContent := content.NewOrganisedContent(contentContainingStr)

	fmt.Println("")
	display.PrintWithIndent(organisedContent)
	fmt.Println("")
}

func completeTodoAction(config *ToDoConfig, includeEdit bool) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, config.SearchStr)
	fmt.Println("")
	organisedContent := content.NewOrganisedContent(contentContainingStr)

	display.PresentItems(organisedContent)
	response := display.AcceptInput("Enter TODO number for completion: ")
	i, err := strconv.Atoi(response)
	if err != nil {
		panic(err)
	}

	var edit string = ""
	if includeEdit {
		edit = display.AcceptInput("Enter edition: ")
	}

	organisedContent.CompleteTODO(i, edit)
	organisedContent.MergeContent()
	manager.WriteUpdatedContent(file, len(contentContainingStr), organisedContent.MergedContent)
	log.Println("Updated content written to file successfully")
}

func moveTodoAction(config *ToDoConfig) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, config.SearchStr)
	fmt.Println("")
	organisedContent := content.NewOrganisedContent(contentContainingStr)

	display.PresentItems(organisedContent)
	response := display.AcceptInput("Enter TODO number for move followed by number for new position: ")
	entries := strings.Fields(response)
	item, err := strconv.Atoi(entries[0])
	if err != nil {
		panic(err)
	}
	newPosition, err := strconv.Atoi(entries[1])
	if err != nil {
		panic(err)
	}

	organisedContent.MoveTODO(item, newPosition)
	organisedContent.MergeContent()
	manager.WriteUpdatedContent(file, len(contentContainingStr), organisedContent.MergedContent)
	log.Println("Updated content written to file successfully")
}

func mergeTodoAction(config *ToDoConfig) {
	file, err := os.OpenFile(config.Filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, config.SearchStr)
	fmt.Println("")
	organisedContent := content.NewOrganisedContent(contentContainingStr)

	display.PresentItems(organisedContent)
	response := display.AcceptInput("Enter TODO number for merge followed by TODO number to merge with: ")
	entries := strings.Fields(response)
	item, err := strconv.Atoi(entries[0])
	if err != nil {
		panic(err)
	}
	mergeWith, err := strconv.Atoi(entries[1])
	if err != nil {
		panic(err)
	}

	organisedContent.MergeTODOs(item, mergeWith)
	organisedContent.MergeContent()
	manager.WriteUpdatedContent(file, len(contentContainingStr), organisedContent.MergedContent)
	log.Println("Updated content written to file successfully")
}
