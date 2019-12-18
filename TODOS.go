package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const filename, searchStr, bufferSize = "TODOs.txt", "TODOs", 1024
const start, current, end = 0, 1, 2

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

	fmt.Println("")
	fmt.Println("")

	fmt.Println(content)
}
