package tasks

import (
	"testing"
	"time"
)

func TestMergeTasks(t *testing.T) {
	data := new(Data)
	data.AddNewTask(Task{Item: "Test this code", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Continue building application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application more", DateCreated: time.Now()})

	data.MergeTasks(4, 2)

	want := 3
	if want != len(data.Tasks) {
		t.Errorf("Merged tasks size is incorrect got %d, want %d", len(data.Tasks), 3)
		t.Log(data.Tasks)
	}
	t.Log(data.Tasks)
}

func TestMergeTasksPanicsWhenPositionIsOutOfBounds(t *testing.T) {
	data := new(Data)
	data.AddNewTask(Task{Item: "Refactor application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application more", DateCreated: time.Now()})

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	data.MergeTasks(7, 2)
}

func TestMergeTasksPanicsWhenDestinationPositionIsOutOfBounds(t *testing.T) {
	data := new(Data)
	data.AddNewTask(Task{Item: "Test this code", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Continue building application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application more", DateCreated: time.Now()})

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()

	data.MergeTasks(2, 9)
}

func TestMergeTasksPanicsWhenDestinationPositionIsOutOfBoundsNegative(t *testing.T) {
	data := new(Data)
	data.AddNewTask(Task{Item: "Test this code", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Continue building application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application more", DateCreated: time.Now()})

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	data.MergeTasks(2, -4)
}

func TestMergeTasksPanicsWhenPositionIsOutOfBoundsNegative(t *testing.T) {
	data := new(Data)
	data.AddNewTask(Task{Item: "Test this code", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Continue building application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application", DateCreated: time.Now()})
	data.AddNewTask(Task{Item: "Refactor application more", DateCreated: time.Now()})

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	data.MergeTasks(-3, 2)
}

func TestMergeTasksPanicsWhenNoDataExists(t *testing.T) {
	data := new(Data)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	data.MergeTasks(-3, 2)
}

func TestMergeTaskToItself(t *testing.T) {
	data := new(Data)
	data.AddNewTask(Task{Item: "Test this code", DateCreated: time.Now()})
	data.MergeTasks(1, 1)

	want := "Test this code"
	if data.Tasks[0].Item != want {
		t.Errorf("Side effect when merging to self, got %s, want %s", data.Tasks[0].Item, want)
		t.Log(data.Tasks)
	}
}
