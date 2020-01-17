package manager

import (
	"fmt"
	. "github.com/davecremins/ToDo-Manager/config"
	"github.com/davecremins/ToDo-Manager/content"
	"github.com/davecremins/ToDo-Manager/dates"
	"log"
	"os"
	"strings"
)

func GetContent(config *ToDoConfig, file *os.File) string {
	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, config.SearchStr)
	return contentContainingStr
}

func AddNewItem(config *ToDoConfig, file *os.File, newItem string) {
	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, "Completed")
	contentSize := len(contentContainingStr)
	log.Println("Position found: ", contentSize)
	// Account for newline
	contentSize += 1

	writingPos := size - int64(contentSize)
	file.Seek(writingPos, 0)
	_, err := file.Write([]byte(newItem))
	if err != nil {
		panic("Falied to write new item to file")
	}

	_, err = file.Write([]byte("\n\n"))
	if err != nil {
		panic("Falied to write new line to file")
	}

	_, err = file.Write([]byte(contentContainingStr))
	if err != nil {
		panic("Falied to write original content to file")
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
	fmt.Println("")
	fmt.Println(newContent)

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
