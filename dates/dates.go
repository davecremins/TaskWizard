package dates

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	re := regexp.MustCompile(`([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}`)
	val := re.Find([]byte(`TODOs 13/12/2019 asdfasdfadsf`))
	splitVals := strings.Split(string(val), "/")
	reverse(splitVals)
	fmt.Println(splitVals)
	formattedDate := strings.Join(splitVals, "-")
	fmt.Println(formattedDate)
	layout := "2006-01-02"
	t, _ := time.Parse(layout, formattedDate)
	fmt.Println(t)
	fmt.Println("Adding a day")
	nextDay := t.AddDate(0, 0, 1)
	fmt.Println(nextDay)
	splitNewDayInfo := strings.Split(nextDay.String(), " ")
	newDate := splitNewDayInfo[0]
	splitNewDate := strings.Split(string(newDate), "-")
	reverse(splitNewDate)
	fmt.Println(splitNewDate)
	formattedNewDate := strings.Join(splitNewDate, "/")
	fmt.Println(formattedNewDate)
}

func reverse(values []string) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}
