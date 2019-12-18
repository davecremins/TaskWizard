package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/davecremins/ToDo-Manager/dates"
)

const filename, searchStr, bufferSize, start = "TODOs.txt", "TODOs", 1024, 0

func main() {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	stats, _ := file.Stat()
	size := stats.Size()
	log.Println("Size of file:", size)

	buffer := make([]byte, bufferSize)
	readPosition := size - bufferSize
	var builder string

	for {
		file.Seek(readPosition, start)
		file.Read(buffer)
		builder = string(buffer) + builder

		if strings.LastIndex(string(buffer), searchStr) >= 0 {
			break
		}

		readPosition -= bufferSize
	}

	pos := strings.LastIndex(builder, searchStr)
	length := len(builder)
	content := builder[pos:length]

	fmt.Println("Content found.")
	fmt.Println("")

	fmt.Println(content)

	dateStr, err := dates.FindDate(content)
	if err != nil {
		panic("Failed to find date in content")
	}
	datetime, err := dates.ConvertToTime(dateStr)
	if err != nil {
		panic("Failed to convert date to time format")
	}
	datetime = dates.AddDay(datetime)
	newDateStr := dates.ExtractShortDate(datetime)
	newContent := strings.ReplaceAll(content, dateStr, newDateStr)

	fmt.Println("Content updated with new date")
	fmt.Println("")
	fmt.Println(newContent)
}
