package display

import (
	"bufio"
	"fmt"
	t "github.com/davecremins/TaskWizard/tasks"
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

func Show(data *t.Data) {
	fmt.Println()
	tbl := table.New("No.", "Task", "Added")
	tbl.WithHeaderFormatter(headerFormat).WithFirstColumnFormatter(columnFormat)
	for i, task := range data.Tasks {
		tbl.AddRow(i+1, task.Item, task.FormatDate())
	}
	tbl.Print()
	fmt.Println()
}
