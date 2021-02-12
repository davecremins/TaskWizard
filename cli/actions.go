package actions

import (
	"flag"
	"bufio"
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

func getDataStore(dataStore string) (*os.File, int64) {
	jsonFile, err := os.Open(dataStore)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	stats, _ := jsonFile.Stat()
	size := stats.Size()
	return jsonFile, size
}

func persistToDataStore(dataStore string, data *t.Data) {
	file, _ := os.OpenFile(dataStore, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	encoder := json.NewEncoder(file)
	encoder.Encode(data)
	file.Close()
}

func decode(file *os.File, data *t.Data) {
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		log.Panicf("Decode issue: %s", err)
	}
}

func showTasks(config *TaskConfig) Action {
	showTaskCmd := flag.NewFlagSet("list", flag.ExitOnError)
	action := func(args []string) {
		jsonFile, size := getDataStore(config.DataStore)
		defer jsonFile.Close()
		if size == 0 {
			log.Println("Data file is empty, no tasks to show")
			return
		}
		data := new(t.Data)
		decode(jsonFile, data)
		display.Show(data)
	}
	addFlagSetDefault(showTaskCmd.Usage)
	return action
}

func newTask(config *TaskConfig) Action {
	newTaskCmd := flag.NewFlagSet("add", flag.ExitOnError)
	task := newTaskCmd.String("desc", "New task placeholder", "Description of new task")
	action := func(args []string) {
		jsonFile, size := getDataStore(config.DataStore)
		defer jsonFile.Close()
		data := new(t.Data)
		if size == 0 {
			log.Println("Data file is empty, no decode required")
		} else {
			decode(jsonFile, data)
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
		jsonFile, size := getDataStore(config.DataStore)
		defer jsonFile.Close()
		if size == 0 {
			log.Println("Data file is empty, no tasks to complete")
			return
		}

		data := new(t.Data)
		decode(jsonFile, data)

		display.Show(data)
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
		jsonFile, size := getDataStore(config.DataStore)
		defer jsonFile.Close()
		if size == 0 {
			log.Println("Data file is empty, no tasks to move")
			return
		}

		data := new(t.Data)
		decode(jsonFile, data)

		display.Show(data)
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
		jsonFile, size := getDataStore(config.DataStore)
		defer jsonFile.Close()
		if size == 0 {
			log.Println("Data file is empty, no tasks to merge")
			return
		}

		data := new(t.Data)
		decode(jsonFile, data)

		display.Show(data)
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

func importTasks(config *TaskConfig) Action {
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	importFilePath := importCmd.String("file", "", "Path to file containing tasks to import")
	action := func(args []string) {
		importCmd.Parse(args[2:])
		if *importFilePath == "" {
			log.Println("Import action needs a path to the file containing tasks")
			return
		}
		jsonFile, size := getDataStore(config.DataStore)
		defer jsonFile.Close()
		data := new(t.Data)
		if size == 0 {
			log.Println("Data file is empty, no tasks to complete")
		} else {
			decode(jsonFile, data)
		}

		// TODO: Import line by line and create new task
		file, err := os.Open(*importFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		imported := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			data.AddNewTask(t.Task{Item: scanner.Text(), DateCreated: time.Now()})
			imported++
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		persistToDataStore(config.DataStore, data)
		msg := fmt.Sprintf("%d tasks imported successfully", imported)
		log.Println(msg)
	}
	addFlagSetDefault(importCmd.Usage)
	return action
}
