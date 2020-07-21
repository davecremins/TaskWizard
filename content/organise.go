package content

import (
	"bufio"
	"fmt"
	"strings"
)

type OrganisedContent struct {
	Content       string
	TODOs         []string
	Completed     []string
	MergedContent string
}

func NewOrganisedContent(content string) *OrganisedContent {
	organisedContent := new(OrganisedContent)
	organisedContent.Content = content

	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)
	emptyLineDetected := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			emptyLineDetected = true
			continue
		}

		if !emptyLineDetected {
			organisedContent.TODOs = append(organisedContent.TODOs, line)
		} else {
			organisedContent.Completed = append(organisedContent.Completed, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return organisedContent
}

func (c *OrganisedContent) MergeContent() {
	var sb strings.Builder
	for _, str := range c.TODOs {
		sb.WriteString(str)
		sb.WriteString("\n")
	}

	sb.WriteString("\n")

	size := len(c.Completed)
	for i, str := range c.Completed {
		sb.WriteString(str)
		if i < size-1 {
			sb.WriteString("\n")
		}
	}

	c.MergedContent = sb.String()
}

func (c *OrganisedContent) CompleteTODO(todoPos int, edit string) {
	lengthIgnoringHeadings := len(c.TODOs) - 2
	if lengthIgnoringHeadings == 0 {
		panic("No TODO to complete")
	}
	if todoPos > lengthIgnoringHeadings || todoPos <= 0 {
		panic("TODO item number is out of bounds")
	}

	// Account for headings
	realIndex := todoPos + 1
	todo := c.TODOs[realIndex]
	if len(edit) > 0 {
		todo = fmt.Sprintf("%s%s%s", todo, " - ", edit)
	}
	c.Completed = append(c.Completed, todo)
	c.TODOs = append(c.TODOs[:realIndex], c.TODOs[realIndex+1:]...)
}

func (c *OrganisedContent) MoveTODO(todoPos, newPosition int) {
	lengthIgnoringHeadings := len(c.TODOs) - 2
	if lengthIgnoringHeadings == 0 {
		panic("No TODOs to move")
	}

	if todoPos == newPosition {
		return
	}

	if todoPos > lengthIgnoringHeadings || todoPos <= 0 {
		panic("TODO item number is out of bounds")
	}

	if newPosition > lengthIgnoringHeadings || newPosition <= 0 {
		panic("New position for TODO item is out of bounds")
	}

	// Account for headings
	realIndex := todoPos + 1
	todo := c.TODOs[realIndex]
	copy(c.TODOs[realIndex:], c.TODOs[realIndex+1:])
	c.TODOs[len(c.TODOs)-1] = ""
	c.TODOs = c.TODOs[:len(c.TODOs)-1]
	realIndexForNewPosition := newPosition + 1
	c.TODOs = append(c.TODOs, "")
	copy(c.TODOs[realIndexForNewPosition+1:], c.TODOs[realIndexForNewPosition:])
	c.TODOs[realIndexForNewPosition] = todo
}

func (c *OrganisedContent) MergeTODOs(todoForMerge, todoToMergeWith int) {
	lengthIgnoringHeadings := len(c.TODOs) - 2
	if lengthIgnoringHeadings == 0 {
		panic("No TODOs to merge")
	}

	if todoForMerge == todoToMergeWith {
		return
	}

	if todoForMerge > lengthIgnoringHeadings || todoForMerge <= 0 {
		panic("TODO item number is out of bounds")
	}

	if todoToMergeWith > lengthIgnoringHeadings || todoToMergeWith <= 0 {
		panic("New position for TODO item is out of bounds")
	}

	// Account for headings
	realIndexForTodoForMerge := todoForMerge + 1
	todoToMergeStr := c.TODOs[realIndexForTodoForMerge]

	realIndexForTodoToMergeWith := todoToMergeWith + 1
	todoToMergeWithStr := c.TODOs[realIndexForTodoToMergeWith]
	c.TODOs[realIndexForTodoToMergeWith] = fmt.Sprintf("%s%s%s", todoToMergeWithStr, " - ", todoToMergeStr)
	c.TODOs = append(c.TODOs[:realIndexForTodoForMerge], c.TODOs[realIndexForTodoForMerge+1:]...)
}
