package todos

import (
	"fmt"
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

func (d *Data) MoveTask(taskNum, newPosition int) {
	taskcount := len(d.ToDos)
	if taskcount == 0 {
		panic("no tasks to move")
	}

	if taskNum == newPosition {
		return
	}

	if taskNum > taskcount || taskNum <= 0 {
		panic("task number is out of bounds")
	}

	if newPosition > taskcount || newPosition <= 0 {
		panic("new position for task is out of bounds")
	}

	// Account for slice index
	realIndex := taskNum - 1
	task := d.ToDos[realIndex]
	copy(d.ToDos[realIndex:], d.ToDos[realIndex+1:])
	d.ToDos = d.ToDos[:len(d.ToDos)-1]
	realIndexForNewPosition := newPosition - 1
	emptyTask := ToDo{}
	d.ToDos = append(d.ToDos, emptyTask)
	copy(d.ToDos[realIndexForNewPosition+1:], d.ToDos[realIndexForNewPosition:])
	d.ToDos[realIndexForNewPosition] = task
}

func (d *Data) MergeTasks(taskNum, mergingTaskNum int) {
	taskcount := len(d.ToDos)
	if taskcount == 0 {
		panic("no tasks to merge")
	}

	if taskNum == mergingTaskNum {
		return
	}

	if taskNum > taskcount || taskNum <= 0 {
		panic("task number is out of bounds")
	}

	if mergingTaskNum > taskcount || mergingTaskNum <= 0 {
		panic("new position for task is out of bounds")
	}

	// Account for headings
	realIndex := taskNum - 1
	taskToMerge := d.ToDos[realIndex]

	realIndexForTaskToMergeWith := mergingTaskNum - 1
	mergingTask := d.ToDos[realIndexForTaskToMergeWith]
	mergedTask := ToDo{Item: fmt.Sprintf("%s%s%s", mergingTask.Item, " - ", taskToMerge.Item), DateCreated: mergingTask.DateCreated}
	d.ToDos[realIndexForTaskToMergeWith] = mergedTask
	d.ToDos = append(d.ToDos[:realIndex], d.ToDos[realIndexForTaskToMergeWith+1:]...)
}
