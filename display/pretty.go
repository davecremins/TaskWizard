package display

import (
	"github.com/fatih/color"
)

type ColoredString func(...interface{}) string

var white ColoredString
var yellow ColoredString
var green ColoredString

func init() {

	white = color.New(color.FgWhite).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
}

func ApplyHeadingColor(s string) string {
	return white(s)
}

func ApplyTODOColor(s string) string {
	return yellow(s)
}

func ApplyCompleteColor(s string) string {
	return green(s)
}
