package display

import (
	"testing"
	"github.com/davecremins/ToDo-Manager/content"
)

func TestPrintWithIndent(t *testing.T) {
	str := "I watch too much Netflix\n"
	organisedContent := content.NewOrganisedContent(str)
	PrintWithIndent(organisedContent)
}
