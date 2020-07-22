package manager

import (
	. "github.com/davecremins/ToDo-Manager/config"
	"github.com/davecremins/ToDo-Manager/dates"
	"log"
	"os"
	"strings"
)

func WriteUpdatedContent(file *os.File, originalContentSize int, newContent string) {
	stats, _ := file.Stat()
	size := stats.Size()

	writingPos := size - int64(originalContentSize)
	file.Seek(writingPos, 0)
	_, err := file.Write([]byte(newContent))
	if err != nil {
		panic("Falied to write updated content to file")
	}
}

func ChangeDate(config *ToDoConfig, content string) string {
	dateStr, err := dates.FindDate(content)
	if err != nil {
		panic("Failed to find date in content")
	}

	datetime, err := dates.ConvertToTime(dateStr)
	if err != nil {
		panic("Failed to convert date to time format")
	}

	datetime = dates.AddDays(datetime, config.DaysToAdd)
	newDateStr := dates.ExtractShortDate(datetime)
	newContent := strings.ReplaceAll(content, dateStr, newDateStr)
	log.Println("Content updated with new date")

	return newContent
}

func WriteContent(file *os.File, content string) {
	file.Seek(0, 2)
	_, err := file.Write([]byte("\n\n"))
	if err != nil {
		panic("Falied to write newlines to file")
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		panic("Falied to write new content to file")
	}
}
