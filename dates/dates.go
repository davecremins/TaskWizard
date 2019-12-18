package dates

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const dateRegEx = `([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}`

func FindDate(content string) string {
	// RegEx date finder
	re := regexp.MustCompile(dateRegEx)
	val := re.Find([]byte(content))

	// Split based on date delimiter
	splitVals := strings.Split(string(val), "/")
	reverse(splitVals)

	// Format date based on time lib requirement
	formattedDate := strings.Join(splitVals, "-")
	layout := "2006-01-02"
	t, _ := time.Parse(layout, formattedDate)

	// Add day
	nextDay := t.AddDate(0, 0, 1)

	// Extract new date
	splitNewDayInfo := strings.Split(nextDay.String(), " ")
	newDate := splitNewDayInfo[0]

	// Format new date to use original delimiter
	splitNewDate := strings.Split(string(newDate), "-")
	reverse(splitNewDate)
	formattedNewDate := strings.Join(splitNewDate, "/")
	return formattedNewDate
}

func reverse(values []string) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}
