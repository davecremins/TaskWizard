package main

import (
	"encoding/json"
	"fmt"
	t "github.com/davecremins/ToDo-Manager/todos"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"time"
)

func main1() {
	data := new(t.Data)
	data.AddNewToDo(t.ToDo{Item: "Test this code", DateCreated: time.Now()})
	data.AddNewToDo(t.ToDo{Item: "Continue building application", DateCreated: time.Now()})
	data.AddNewToDo(t.ToDo{Item: "Continue building application", DateCreated: time.Now()})
	data.AddCompletedItem(t.Done{Item: "Start simplified ToDo App", DateCompleted: time.Now()})

	file, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	encoder := json.NewEncoder(file)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	file.Close()

	jsonFile, _ := os.Open("data.json")
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	if err = decoder.Decode(data); err != nil {
		fmt.Println(err)
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("No.", "ToDo", "Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for i, todo := range data.ToDos {
		tbl.AddRow(i+1, todo.Item, todo.DateCreated)
	}

	tbl.Print()

	fmt.Println("Adding new TODO to in-memory structure")

	data.AddNewToDo(t.ToDo{Item: "Continue building application on another night", DateCreated: time.Now()})

	file, _ = os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	encoder = json.NewEncoder(file)
	encoder.Encode(data)
	file.Close()

	tbl1 := table.New("No.", "ToDo", "Added")
	tbl1.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for i, todo := range data.ToDos {
		tbl1.AddRow(i+1, todo.Item, todo.DateCreated)
	}

	tbl1.Print()
}
