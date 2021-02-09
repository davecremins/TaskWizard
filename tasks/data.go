package tasks

import (
	"fmt"
	"time"
)

type Task struct {
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
	Tasks     []Task `json:"tasks"`
	Completed []Done `json:"completed"`
}

func (d *Data) AddNewTask(task Task) {
	d.Tasks = append(d.Tasks, task)
}

func (d *Data) AddCompletedItem(item Done) {
	d.Completed = append(d.Completed, item)
}

func (d *Data) CompleteTask(taskNum int, taskComment string) {
	taskCount := len(d.Tasks)
	if taskCount == 0 {
		panic("No task to complete")
	}
	if taskNum > taskCount || taskNum <= 0 {
		panic("Task number is out of bounds")
	}

	// Account for slice index
	realIndex := taskNum - 1
	task := d.Tasks[realIndex]
	completedTask := Done{Item: task.Item, DateCreated: task.DateCreated, DateCompleted: time.Now(), Comment: taskComment}
	d.AddCompletedItem(completedTask)
	d.Tasks = append(d.Tasks[:realIndex], d.Tasks[realIndex+1:]...)
}

func (d *Data) MoveTask(taskNum, newPosition int) {
	taskcount := len(d.Tasks)
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
	task := d.Tasks[realIndex]
	copy(d.Tasks[realIndex:], d.Tasks[realIndex+1:])
	d.Tasks = d.Tasks[:len(d.Tasks)-1]
	realIndexForNewPosition := newPosition - 1
	emptyTask := Task{}
	d.Tasks = append(d.Tasks, emptyTask)
	copy(d.Tasks[realIndexForNewPosition+1:], d.Tasks[realIndexForNewPosition:])
	d.Tasks[realIndexForNewPosition] = task
}

func (d *Data) MergeTasks(taskNum, mergeWith int) {
	taskcount := len(d.Tasks)
	if taskcount == 0 {
		panic("no tasks to merge")
	}

	if taskNum == mergeWith {
		return
	}

	if taskNum > taskcount || taskNum <= 0 {
		panic("task number is out of bounds")
	}

	if mergeWith > taskcount || mergeWith <= 0 {
		panic("new position for task is out of bounds")
	}

	realIndex := taskNum - 1
	taskToMerge := d.Tasks[realIndex]

	realIndexForTaskToMergeWith := mergeWith - 1
	mergingTask := d.Tasks[realIndexForTaskToMergeWith]
	mergedTask := Task{Item: fmt.Sprintf("%s%s%s", mergingTask.Item, " - ", taskToMerge.Item), DateCreated: mergingTask.DateCreated}
	d.Tasks[realIndexForTaskToMergeWith] = mergedTask
	d.Tasks = append(d.Tasks[:realIndex], d.Tasks[realIndex+1:]...)
}
