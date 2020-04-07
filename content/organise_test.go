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
