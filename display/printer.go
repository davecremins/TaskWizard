package display

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const indentLevel = 8
var indentStr string

func init() {
	for i := 0; i < indentLevel; i++ {
		indentStr += " "
	}
}

func PrintWithIndent(content string) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(indentStr + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func PresentItems(content string) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)
	i, beginNumbering, num := 0, 1, 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		if i > beginNumbering {
			s := fmt.Sprintf("%s%s) %s", indentStr, strconv.Itoa(num), line)
			fmt.Println(ApplyTODOColor(s))
			num++
		} else {
			fmt.Println(ApplyHeadingColor(indentStr + line))
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
	fmt.Printf("%sEnter TODO number for completion: ", indentStr)
	scanner.Scan()
	text := scanner.Text()
	return text
}
