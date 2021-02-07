package actions

import (
	"flag"
	"fmt"
	. "github.com/davecremins/ToDo-Manager/config"
	"github.com/davecremins/ToDo-Manager/content"
	"github.com/davecremins/ToDo-Manager/display"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	t "github.com/davecremins/ToDo-Manager/todos"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Action func([]string)

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

func showTasks(config *ToDoConfig) Action {
	action := func(args []string) {
		jsonFile, err := os.Open("data.json")
		defer jsonFile.Close()
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to show")
			return
		}

		decoder := json.NewDecoder(jsonFile)

		data := new(t.Data)
		if err = decoder.Decode(data); err != nil {
			log.Panicf("Decode issue: %s", err)
		}

		// TODO: Move this into display package
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("No.", "ToDo", "Added")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for i, todo := range data.ToDos {
			tbl.AddRow(i+1, todo.Item, todo.DateCreated)
		}
		tbl.Print()
	}
	return action
}

func newTask(config *ToDoConfig) Action {
	newTaskCmd := flag.NewFlagSet("add", flag.ExitOnError)
	task := newTaskCmd.String("desc", "New task placeholder", "Description of new task")
	action := func(args []string) {
		jsonFile, err := os.Open("data.json")
		defer jsonFile.Close()
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		data := new(t.Data)

		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no decode required")
		} else {
			decoder := json.NewDecoder(jsonFile)
			if err = decoder.Decode(data); err != nil {
				log.Panicf("Decode issue: %s", err)
			}
		}

		newTaskCmd.Parse(args[2:])
		newTask := t.ToDo{Item: *task, DateCreated: time.Now()}
		data.AddNewToDo(newTask)

		file, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		encoder := json.NewEncoder(file)
		encoder.Encode(data)
		file.Close()

		log.Println("New task added successfully")

	}
	addFlagSetDefault(newTaskCmd.Usage)
	return action
}

// API to create a new task
func newTodoActionMakeup(config *ToDoConfig) Action {
	newTaskCmd := flag.NewFlagSet("newtodo", flag.ExitOnError)
	searchStr := newTaskCmd.String("search", config.SearchStr, "Search string to look for")
	filename := newTaskCmd.String("filename", config.Filename, "Name of file to add new todo")
	todo := newTaskCmd.String("desc", "New todo item placeholder", "Description of new todo")
	action := func(args []string) {
		newTaskCmd.Parse(args[2:])
		config.SearchStr = *searchStr
		config.Filename = *filename
		log.Println("Config over-written for newtodo action")
		newTodoAction(config, *todo)
	}
	addFlagSetDefault(newTaskCmd.Usage)
	return action
}

// API to list tasks
func todaysTodosActionMakeup(config *ToDoConfig) Action {
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

// API to complete task
func completeTodoActionMakeup(config *ToDoConfig) Action {
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

// API to reorder task
func moveTodoActionMakeup(config *ToDoConfig) Action {
	moveCmd := flag.NewFlagSet("move", flag.ExitOnError)
	action := func(args []string) {
		log.Println("Config not over-written for move action")
		moveTodoAction(config)
	}
	addFlagSetDefault(moveCmd.Usage)
	return action
}

// API to merge tasks
func mergeTodoActionMakeup(config *ToDoConfig) Action {
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	action := func(args []string) {
		log.Println("Config not over-written for move action")
		mergeTodoAction(config)
	}
	addFlagSetDefault(mergeCmd.Usage)
	return action
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
	display.PrintWithIndent(organisedContent, content.ALL)
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

	display.PrintWithIndent(organisedContent, content.TODOS)
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

	writingPos := size - int64(len(contentContainingStr))
	file.Seek(writingPos, 0)
	_, err = file.Write([]byte(organisedContent.MergedContent))
	if err != nil {
		panic("Falied to write updated content to file")
	}
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

	display.PrintWithIndent(organisedContent, content.TODOS)
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

	writingPos := size - int64(len(contentContainingStr))
	file.Seek(writingPos, 0)
	_, err = file.Write([]byte(organisedContent.MergedContent))
	if err != nil {
		panic("Falied to write updated content to file")
	}
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

	display.PrintWithIndent(organisedContent, content.TODOS)
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

	writingPos := size - int64(len(contentContainingStr))
	file.Seek(writingPos, 0)
	_, err = file.Write([]byte(organisedContent.MergedContent))
	if err != nil {
		panic("Falied to write updated content to file")
	}
	log.Println("Updated content written to file successfully")
}
