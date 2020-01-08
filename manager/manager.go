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

func CopyPreviousContent(config *ToDoConfig, file *os.File) string {
	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, config.SearchStr)

	log.Println("Content found.")
	fmt.Println("")
	fmt.Println(contentContainingStr)

	return contentContainingStr
}

func AddNewItem(config *ToDoConfig, file *os.File, newItem string) {
	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	position := content.FindSearchStrLocation(file, size, "Completed")
	log.Println("Position found: ", position)
	log.Println("Increasing size by 1 to account for newline")
	position += 1

	writingPos := size - int64(position)
	file.Seek(writingPos, 0)
	_, err := file.Write([]byte(newItem))
	if err != nil {
		panic("Falied to write new item to file")
	}

	_, err = file.Write([]byte("\n"))
	if err != nil {
		panic("Falied to write new content to file")
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
