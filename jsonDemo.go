package main

import (
	t "github.com/davecremins/ToDo-Manager/todos"
	"time"
	"encoding/json"
	"os"
	"fmt"
)

func main() {
	data := new(t.Data)
	data.AddNewToDo(t.ToDo{Item: "Test this code", DateCreated: time.Now()})
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

	data.AddNewToDo(t.ToDo{Item: "Continue building application on another night", DateCreated: time.Now()})
	fmt.Println("Size of ToDos now: ", len(data.ToDos))

	file, _ = os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	encoder = json.NewEncoder(file)
	encoder.Encode(data)
	file.Close()
}
