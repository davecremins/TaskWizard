package manager

import (
	"os"
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
