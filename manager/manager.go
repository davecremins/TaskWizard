package manager

import (
	"os"
)

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
