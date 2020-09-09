package display

import (
	"bufio"
	"fmt"
	"github.com/davecremins/ToDo-Manager/content"
	"os"
	"strconv"
)

const indentLevel = 8

var indentStr string

func init() {
	for i := 0; i < indentLevel; i++ {
		indentStr += " "
	}
}

func PrintWithIndent(organisedContent *content.OrganisedContent) {
	todoHeadings := organisedContent.TODOs[:2]
	for _, heading := range todoHeadings {
		fmt.Println(fmt.Sprintf("%s%s", indentStr, ApplyHeadingColor(heading)))
	}
	todos := organisedContent.TODOs[2:]
	for i, todo := range todos {
		s := fmt.Sprintf("%s%s) %s", indentStr, strconv.Itoa(i+1), todo)
		fmt.Println(ApplyTODOColor(s))
	}

	fmt.Println("")

	completedHeadings := organisedContent.Completed[:2]
	for _, heading := range completedHeadings {
		fmt.Println(fmt.Sprintf("%s%s", indentStr, ApplyHeadingColor(heading)))
	}
	completed := organisedContent.Completed[2:]
	for i, complete := range completed {
		s := fmt.Sprintf("%s%s) %s", indentStr, strconv.Itoa(i+1), complete)
		fmt.Println(ApplyCompleteColor(s))
	}
}

func PresentItems(organisedContent *content.OrganisedContent) {
	beginNumbering, num := 1, 1
	for i, todo := range organisedContent.TODOs {
		if i > beginNumbering {
			s := fmt.Sprintf("%s%s) %s", indentStr, strconv.Itoa(num), todo)
			fmt.Println(ApplyTODOColor(s))
			num++
		} else {
			fmt.Println(ApplyHeadingColor(indentStr + todo))
		}
	}
}

func AcceptInput(prompt string) string {
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s%s", indentStr, prompt)
	scanner.Scan()
	text := scanner.Text()
	return text
}
