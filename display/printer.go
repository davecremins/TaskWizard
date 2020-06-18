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
	i, str := 0, ""
	for _, todo := range organisedContent.TODOs {
		if i <= 1 {
			str = ApplyHeadingColor(todo)
		} else {
			str = ApplyTODOColor(todo)
		}
		i++
		fmt.Println(fmt.Sprintf("%s%s", indentStr, str))
	}

	i = 0
	fmt.Println("")

	for _, completed := range organisedContent.Completed {
		if i <= 1 {
			str = ApplyHeadingColor(completed)
		} else {
			str = ApplyCompleteColor(completed)
		}
		i++
		fmt.Println(fmt.Sprintf("%s%s", indentStr, str))
	}
}

func PresentItems(organisedContent *content.OrganisedContent) {
	i, beginNumbering, num := 0, 1, 1
	for _, todo := range organisedContent.TODOs {
		if i > beginNumbering {
			s := fmt.Sprintf("%s%s) %s", indentStr, strconv.Itoa(num), todo)
			fmt.Println(ApplyTODOColor(s))
			num++
		} else {
			fmt.Println(ApplyHeadingColor(indentStr + todo))
		}
		i++
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
