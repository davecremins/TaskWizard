package dates

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

const dateRegEx, layout = `([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}`, "2006-01-02"

func FindDate(content string) (string, error) {
	// RegEx date finder
	re := regexp.MustCompile(dateRegEx)
	val := re.Find([]byte(content))
	dateStr := string(val)
	if dateStr == "" {
		return "", errors.New("No date found in provided content")
	}
	return dateStr, nil
}

func ConvertToTime(dateStr string) (time.Time, error) {
	// Split based on date delimiter
	splitVals := strings.Split(dateStr, "/")
	reverse(splitVals)

	// Format date based on time lib requirement
	formattedDate := strings.Join(splitVals, "-")
	return time.Parse(layout, formattedDate)
}

func ExtractShortDate(t time.Time) string {
	// Extract new date
	splitNewDayInfo := strings.Split(t.String(), " ")
	newDate := splitNewDayInfo[0]

	// Format new date to use original delimiter
	splitNewDate := strings.Split(string(newDate), "-")
	reverse(splitNewDate)
	formattedNewDate := strings.Join(splitNewDate, "/")
	return formattedNewDate

}

func AddDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func reverse(values []string) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}
