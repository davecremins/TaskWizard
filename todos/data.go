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
	DateCreated   time.Time
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

func (d *Data) CompleteTask(taskNum int, taskComment string) {
	taskCount := len(d.ToDos)
	if taskCount == 0 {
		panic("No task to complete")
	}
	if taskNum > taskCount || taskNum <= 0 {
		panic("Task number is out of bounds")
	}

	// Account for slice index
	realIndex := taskNum - 1
	task := d.ToDos[realIndex]
	completedTask := Done{Item: task.Item, DateCreated: task.DateCreated, DateCompleted: time.Now(), Comment: taskComment}
	d.AddCompletedItem(completedTask)
	d.ToDos = append(d.ToDos[:realIndex], d.ToDos[realIndex+1:]...)
}
