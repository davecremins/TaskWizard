package display

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

const level = 8

func PrintWithIndent(content string) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(indent + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
