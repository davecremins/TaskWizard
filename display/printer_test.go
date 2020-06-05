package display

import (
	"github.com/davecremins/ToDo-Manager/content"
	"testing"
)

func TestPrintWithIndent(t *testing.T) {
	str := "I watch too much Netflix\n"
	organisedContent := content.NewOrganisedContent(str)
	PrintWithIndent(organisedContent)
}
