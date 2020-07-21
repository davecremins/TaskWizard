package content

import "testing"

func TestContentIsOrganised(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	organisedContent := NewOrganisedContent(content)
	todosWant, completedWant := 4, 3
	todosLength, completedLength := len(organisedContent.TODOs), len(organisedContent.Completed)
	if todosLength != todosWant {
		t.Errorf("Incorrect length for TODOs got %d, want %d", todosLength, todosWant)
	}

	if completedLength != completedWant {
		t.Errorf("Incorrect length for Completed got %d, want %d", completedLength, completedWant)
	}

}

func TestContentIsMergedWhenCompletedItemIsRemoved(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	want := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================`

	organisedContent := NewOrganisedContent(content)
	i := 1
	organisedContent.Completed = append(organisedContent.Completed[:i-1], organisedContent.Completed[:i+1]...)
	organisedContent.MergeContent()
	if want != organisedContent.MergedContent {
		t.Errorf("Merged content is incorrect got %s, want %s", organisedContent.MergedContent, want)
		t.Logf("Size of got %d and want %d", len(organisedContent.MergedContent), len(want))
	}
}

func TestContentIsMergedWhenTODOItemIsRemoved(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	want := `TODOs 13/01/2020
================

Completed 13/01/2020
====================
Finished a book
Do Something in the evening
Write some code`

	organisedContent := NewOrganisedContent(content)
	for i, str := range organisedContent.TODOs {
		if i > 1 {
			organisedContent.Completed = append(organisedContent.Completed, str)
		}
	}
	i := 1
	organisedContent.TODOs = append(organisedContent.TODOs[:i-1], organisedContent.TODOs[:i+1]...)
	organisedContent.MergeContent()
	if want != organisedContent.MergedContent {
		t.Errorf("Merged content is incorrect got %s, want %s", organisedContent.MergedContent, want)
		t.Logf("Size of got %d and want %d", len(organisedContent.MergedContent), len(want))
	}
}

func TestMoveTODOToCompleted(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	want := `TODOs 13/01/2020
================
Write some code

Completed 13/01/2020
====================
Finished a book
Do Something in the evening`

	organisedContent := NewOrganisedContent(content)
	organisedContent.CompleteTODO(1)
	organisedContent.MergeContent()

	if want != organisedContent.MergedContent {
		t.Errorf("Merged content is incorrect got %s, want %s", organisedContent.MergedContent, want)
		t.Logf("Size of got %d and want %d", len(organisedContent.MergedContent), len(want))
	}
}

func TestCompleteTODOPanicsWhenOutOfBounds(t *testing.T) {
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
	organisedContent.CompleteTODO(5)
}

func TestMoveTODOItemToNewPositionInTODOItems(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	want := `TODOs 13/01/2020
================
Write some code
Do Something in the evening

Completed 13/01/2020
====================
Finished a book`

	organisedContent := NewOrganisedContent(content)
	organisedContent.MoveTODO(2, 1)
	organisedContent.MergeContent()

	if want != organisedContent.MergedContent {
		t.Errorf("Merged content is incorrect got %s, want %s", organisedContent.MergedContent, want)
		t.Logf("Size of got %d and want %d", len(organisedContent.MergedContent), len(want))
	}
}

func TestMoveTODOItemForwardToNewPositionInTODOItems(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code
Finish laying the fence

Completed 13/01/2020
====================
Finished a book`

	want := `TODOs 13/01/2020
================
Write some code
Finish laying the fence
Do Something in the evening

Completed 13/01/2020
====================
Finished a book`

	organisedContent := NewOrganisedContent(content)
	organisedContent.MoveTODO(1, 3)
	organisedContent.MergeContent()

	if want != organisedContent.MergedContent {
		t.Errorf("Merged content is incorrect got %s, want %s", organisedContent.MergedContent, want)
		t.Logf("Size of got %d and want %d", len(organisedContent.MergedContent), len(want))
	}
}

func TestMoveTODOItemToSamePositionInTODOItems(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code
Finish laying the fence

Completed 13/01/2020
====================
Finished a book`

	want := `TODOs 13/01/2020
================
Do Something in the evening
Write some code
Finish laying the fence

Completed 13/01/2020
====================
Finished a book`

	organisedContent := NewOrganisedContent(content)
	organisedContent.MoveTODO(2, 2)
	organisedContent.MergeContent()

	if want != organisedContent.MergedContent {
		t.Errorf("Merged content is incorrect got %s, want %s", organisedContent.MergedContent, want)
		t.Logf("Size of got %d and want %d", len(organisedContent.MergedContent), len(want))
	}
}

func TestMoveTODOPanicsWhenToDoItemIsOutOfBounds(t *testing.T) {
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
	organisedContent.MoveTODO(3, 5)
}

func TestMoveTODOPanicsWhenNewToDoPositionIsOutOfBounds(t *testing.T) {
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
	organisedContent.MoveTODO(1, 3)
}

func TestMergeTODOItems(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code
Finish laying the fence

Completed 13/01/2020
====================
Finished a book`

	want := `TODOs 13/01/2020
================
Do Something in the evening - Write some code
Finish laying the fence

Completed 13/01/2020
====================
Finished a book`

	organisedContent := NewOrganisedContent(content)
	organisedContent.MergeTODOs(2, 1)
	organisedContent.MergeContent()

	if want != organisedContent.MergedContent {
		t.Errorf("Merged content is incorrect got %s, want %s", organisedContent.MergedContent, want)
		t.Logf("Size of got %d and want %d", len(organisedContent.MergedContent), len(want))
	}
}
