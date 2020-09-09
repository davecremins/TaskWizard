package display

import (
	"github.com/davecremins/ToDo-Manager/content"
	"testing"
)

func TestPrintWithIndent(t *testing.T) {
	str := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	organisedContent := content.NewOrganisedContent(str)
	PrintWithIndent(organisedContent, content.ALL)
}
