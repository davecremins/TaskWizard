package content

import (
	"bufio"
	"strings"
)

type OrganisedContent struct {
	Content   string
	TODOs     []string
	Completed []string
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
		panic("error scanning content")
	}

	return organisedContent
}
