package display

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
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

func PresentItems(content string) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)

	i, beginNumbering, num := 0, 1, 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		if i > beginNumbering {
			fmt.Println(indent + strconv.Itoa(num) + ") " + line)
			num++
		} else {
			fmt.Println(indent + line)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
