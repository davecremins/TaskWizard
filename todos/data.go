package todos

import (
	"time"
)

type ToDo struct {
	Item        string
	DateCreated time.Time
}

type Done struct {
	Item          string
	DateCompleted time.Time
	Comment       string
}

type Data struct {
	ToDos     []ToDo `json:"todos"`
	Completed []Done `json:"completed"`
}

func (d *Data) AddNewToDo(todo ToDo) {
	d.ToDos = append(d.ToDos, todo)
}

func (d *Data) AddCompletedItem(item Done) {
	d.Completed = append(d.Completed, item)
}
