package content

import (
	"bufio"
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
