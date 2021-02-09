package display

import (
	"bufio"
	"fmt"
	"os"
)

const indentLevel = 8

var indentStr string

func init() {
	for i := 0; i < indentLevel; i++ {
		indentStr += " "
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
