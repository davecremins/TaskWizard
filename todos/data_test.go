package todos

import (
	"testing"
	"time"
)

func TestMergeTasks(t *testing.T) {
	data := new(Data)
	data.AddNewToDo(ToDo{Item: "Test this code", DateCreated: time.Now()})
	data.AddNewToDo(ToDo{Item: "Continue building application", DateCreated: time.Now()})
	data.AddNewToDo(ToDo{Item: "Refactor application", DateCreated: time.Now()})
	data.AddNewToDo(ToDo{Item: "Refactor application more", DateCreated: time.Now()})

	data.MergeTasks(4, 2)

	want := 3
	if want != len(data.ToDos) {
		t.Errorf("Merged tasks size is incorrect got %d, want %d", len(data.ToDos), 3)
		t.Log(data.ToDos)
	}
	t.Log(data.ToDos)
}

/*func TestMergeTODOsPanicsWhenToDoPositionIsOutOfBounds(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	organisedContent := NewOrganisedContent(content)
	organisedContent.MergeTODOs(7, 1)
}

func TestMergeTODOsPanicsWhenToDoDestinationPositionIsOutOfBounds(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	organisedContent := NewOrganisedContent(content)
	organisedContent.MergeTODOs(1, 3)
}

func TestMergeTODOsPanicsWhenToDoPositionIsOutOfBoundsNegative(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	organisedContent := NewOrganisedContent(content)
	organisedContent.MergeTODOs(-1, 3)
}

func TestMergeTODOsPanicsWhenToDoDestinationPositionIsOutOfBoundsNegative(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic with bad input")
		}
		t.Log(r)
	}()
	organisedContent := NewOrganisedContent(content)
	organisedContent.MergeTODOs(1, -3)
}*/
