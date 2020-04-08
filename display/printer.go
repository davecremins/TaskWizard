package display

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const level = 8

func createIndentStr() string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}
	return indent
}

func PrintWithIndent(content string) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)
	indent := createIndentStr()
	for scanner.Scan() {
		fmt.Println(indent + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func PresentItems(content string) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)
	indent := createIndentStr()
	i, beginNumbering, num := 0, 1, 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		if i > beginNumbering {
			fmt.Println(ApplyTODOColor(indent + strconv.Itoa(num) + ") " + line))
			num++
		} else {
			fmt.Println(ApplyHeadingColor(indent + line))
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func AcceptInput() string {
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	indent := createIndentStr()
	fmt.Print(indent + "Enter TODO number for completion: ")
	scanner.Scan()
	text := scanner.Text()
	return text
}
