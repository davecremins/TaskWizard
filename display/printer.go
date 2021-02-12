package display

import (
	"bufio"
	"fmt"
	"os"
	t "github.com/davecremins/TaskWizard/tasks"
	"github.com/fatih/color"
	"github.com/rodaine/table"
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

func Show(data *t.Data) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("No.", "Task", "Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for i, task := range data.Tasks {
		tbl.AddRow(i+1, task.Item, task.DateCreated)
	}
	tbl.Print()
}
