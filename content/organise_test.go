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
