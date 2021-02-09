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

		tbl := table.New("No.", "Task", "Added")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for i, task := range data.Tasks {
			tbl.AddRow(i+1, task.Item, task.DateCreated)
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
		newTask := t.Task{Item: *task, DateCreated: time.Now()}
		data.AddNewTask(newTask)

		file, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		encoder := json.NewEncoder(file)
		encoder.Encode(data)
		file.Close()

		log.Println("New task added successfully")

	}
	addFlagSetDefault(newTaskCmd.Usage)
	return action
}

func completeTask(config *ToDoConfig) Action {
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	includeComment := completeCmd.Bool("comment", false, "Option to include comment for task before completion")
	action := func(args []string) {
		completeCmd.Parse(args[2:])
		jsonFile, err := os.Open("data.json")
		defer jsonFile.Close()
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to complete")
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

		file, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		encoder := json.NewEncoder(file)
		encoder.Encode(data)
		file.Close()

		log.Println("Task completed successfully")
	}
	addFlagSetDefault(completeCmd.Usage)
	return action
}

func moveTask(config *ToDoConfig) Action {
	moveCmd := flag.NewFlagSet("move", flag.ExitOnError)
	action := func(args []string) {
		jsonFile, err := os.Open("data.json")
		defer jsonFile.Close()
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to move")
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

		file, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		encoder := json.NewEncoder(file)
		encoder.Encode(data)
		file.Close()

		log.Println("Task moved successfully")
	}
	addFlagSetDefault(moveCmd.Usage)
	return action

}

func mergeTasks(config *ToDoConfig) Action {
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	action := func(args []string) {

		jsonFile, err := os.Open("data.json")
		defer jsonFile.Close()
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		stats, _ := jsonFile.Stat()
		size := stats.Size()
		if size == 0 {
			log.Println("Data file is empty, no tasks to merge")
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
		file, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		encoder := json.NewEncoder(file)
		encoder.Encode(data)
		file.Close()

		log.Println("Task merged successfully")
	}
	addFlagSetDefault(mergeCmd.Usage)
	return action

}
