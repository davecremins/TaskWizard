package display

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
)

const indentLevel = 8

var (
	indentStr    string
	headerFormat table.Formatter
	columnFormat table.Formatter
)

func init() {
	for i := 0; i < indentLevel; i++ {
		indentStr += " "
	}
	headerFormat = color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFormat = color.New(color.FgYellow).SprintfFunc()
}

func AcceptInput(prompt string) string {
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s%s", indentStr, prompt)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func Present(tbl table.Table) {
	fmt.Println()
	tbl.Print()
	fmt.Println()
}

func PrepareTable(headers ...interface{}) table.Table {
	tbl := table.New(headers...)
	tbl.WithHeaderFormatter(headerFormat).WithFirstColumnFormatter(columnFormat)
	return tbl
}

func PopulateTable(tbl table.Table, rowData ...interface{}) {
	tbl.AddRow(rowData...)
}
