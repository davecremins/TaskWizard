package actions

import (
	"flag"
	"fmt"
	. "github.com/davecremins/TaskWizard/config"
	"github.com/davecremins/TaskWizard/display"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	t "github.com/davecremins/TaskWizard/tasks"
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

func getDataStore(dataStore string) *os.File {
	jsonFile, err := os.Open(dataStore)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	return jsonFile
}

func persistToDataStore(dataStore string, data *t.Data) {
	file, _ := os.OpenFile(dataStore, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	encoder := json.NewEncoder(file)
	encoder.Encode(data)
	file.Close()
}

func showTasks(config *TaskConfig) Action {
	showTaskCmd := flag.NewFlagSet("list", flag.ExitOnError)
	action := func(args []string) {
		jsonFile := getDataStore(config.DataStore)
		defer jsonFile.Close()
		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to show")
			return
		}

		decoder := json.NewDecoder(jsonFile)

		data := new(t.Data)
		if err := decoder.Decode(data); err != nil {
			log.Panicf("Decode issue: %s", err)
		}

		// TODO: Move this into display package
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("No.", "Task", "Added")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for i, task := range data.Tasks {
			tbl.AddRow(i+1, task.Item, task.DateCreated)
		}
		tbl.Print()
	}
	addFlagSetDefault(showTaskCmd.Usage)
	return action
}

func newTask(config *TaskConfig) Action {
	newTaskCmd := flag.NewFlagSet("add", flag.ExitOnError)
	task := newTaskCmd.String("desc", "New task placeholder", "Description of new task")
	action := func(args []string) {
		jsonFile := getDataStore(config.DataStore)
		defer jsonFile.Close()
		stats, _ := jsonFile.Stat()
		size := stats.Size()
		data := new(t.Data)
		if size == 0 {
			log.Println("Data file is empty, no decode required")
		} else {
			decoder := json.NewDecoder(jsonFile)
			if err := decoder.Decode(data); err != nil {
				log.Panicf("Decode issue: %s", err)
			}
		}

		newTaskCmd.Parse(args[2:])
		newTask := t.Task{Item: *task, DateCreated: time.Now()}
		data.AddNewTask(newTask)
		persistToDataStore(config.DataStore, data)
		log.Println("New task added successfully")

	}
	addFlagSetDefault(newTaskCmd.Usage)
	return action
}

func completeTask(config *TaskConfig) Action {
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	includeComment := completeCmd.Bool("comment", false, "Option to include comment for task before completion")
	action := func(args []string) {
		completeCmd.Parse(args[2:])
		jsonFile := getDataStore(config.DataStore)
		defer jsonFile.Close()
		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to complete")
			return
		}

		decoder := json.NewDecoder(jsonFile)

		data := new(t.Data)
		if err := decoder.Decode(data); err != nil {
			log.Panicf("Decode issue: %s", err)
		}

		// TODO: Move this into display package
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("No.", "Task", "Created")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for i, task := range data.Tasks {
			tbl.AddRow(i+1, task.Item, task.DateCreated)
		}
		tbl.Print()

		fmt.Println("")
		fmt.Println("")

		response := display.AcceptInput("Enter task number for completion: ")
		i, err := strconv.Atoi(response)
		if err != nil {
			panic(err)
		}

		var comment string = ""
		if *includeComment {
			comment = display.AcceptInput("Enter comment: ")
		}

		data.CompleteTask(i, comment)
		persistToDataStore(config.DataStore, data)
		log.Println("Task completed successfully")
	}
	addFlagSetDefault(completeCmd.Usage)
	return action
}

func moveTask(config *TaskConfig) Action {
	moveCmd := flag.NewFlagSet("move", flag.ExitOnError)
	action := func(args []string) {
		jsonFile := getDataStore(config.DataStore)
		defer jsonFile.Close()
		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to move")
			return
		}

		decoder := json.NewDecoder(jsonFile)

		data := new(t.Data)
		if err := decoder.Decode(data); err != nil {
			log.Panicf("Decode issue: %s", err)
		}

		// TODO: Move this into display package
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("No.", "Task", "Created")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for i, task := range data.Tasks {
			tbl.AddRow(i+1, task.Item, task.DateCreated)
		}
		tbl.Print()

		fmt.Println("")
		fmt.Println("")

		response := display.AcceptInput("Enter task number that you want to move followed by number for new position: ")
		entries := strings.Fields(response)
		taskNum, err := strconv.Atoi(entries[0])
		if err != nil {
			panic(err)
		}
		newPosition, err := strconv.Atoi(entries[1])
		if err != nil {
			panic(err)
		}

		data.MoveTask(taskNum, newPosition)
		persistToDataStore(config.DataStore, data)
		log.Println("Task moved successfully")
	}
	addFlagSetDefault(moveCmd.Usage)
	return action

}

func mergeTasks(config *TaskConfig) Action {
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	action := func(args []string) {
		jsonFile := getDataStore(config.DataStore)
		defer jsonFile.Close()

		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to merge")
			return
		}

		decoder := json.NewDecoder(jsonFile)

		data := new(t.Data)
		if err := decoder.Decode(data); err != nil {
			log.Panicf("Decode issue: %s", err)
		}

		// TODO: Move this into display package
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("No.", "Task", "Created")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for i, task := range data.Tasks {
			tbl.AddRow(i+1, task.Item, task.DateCreated)
		}
		tbl.Print()

		fmt.Println("")
		fmt.Println("")

		response := display.AcceptInput("Enter task number for merge followed by task number to merge with: ")
		entries := strings.Fields(response)
		item, err := strconv.Atoi(entries[0])
		if err != nil {
			panic(err)
		}
		mergeWith, err := strconv.Atoi(entries[1])
		if err != nil {
			panic(err)
		}

		data.MergeTasks(item, mergeWith)
		persistToDataStore(config.DataStore, data)
		log.Println("Task merged successfully")
	}
	addFlagSetDefault(mergeCmd.Usage)
	return action

}
