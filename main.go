package main

import (
	"fmt"
	"github.com/davecremins/ToDo-Manager/content"
	"github.com/davecremins/ToDo-Manager/dates"
	"log"
	"os"
	"strings"
)

const filename, searchStr = "TODOs.txt", "TODOs"

func main() {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	contentContainingStr := content.FindSearchStr(file, size, searchStr)

	log.Println("Content found.")
	fmt.Println("")
	fmt.Println(contentContainingStr)

	dateStr, err := dates.FindDate(contentContainingStr)
	if err != nil {
		panic("Failed to find date in content")
	}

	datetime, err := dates.ConvertToTime(dateStr)
	if err != nil {
		panic("Failed to convert date to time format")
	}

	datetime = dates.AddDay(datetime)
	newDateStr := dates.ExtractShortDate(datetime)
	newContent := strings.ReplaceAll(contentContainingStr, dateStr, newDateStr)

	log.Println("Content updated with new date")
	fmt.Println("")
	fmt.Println(newContent)

	file.Seek(0, 2)
	_, err = file.Write([]byte("\n\n"))
	if err != nil {
		panic("Falied to write newlines to file")
	}

	_, err = file.Write([]byte(newContent))
	if err != nil {
		panic("Falied to write new content to file")
	}

}
