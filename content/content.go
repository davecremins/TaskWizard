package content

import (
	"os"
	"strings"
)

const bufferSize = 1024

func FindSearchStr(contents *os.File, size int64, searchStr string) string {
	buffer := make([]byte, bufferSize)
	readPosition := size - bufferSize
	var builder string

	for {
		contents.Seek(readPosition, 0)
		contents.Read(buffer)
		builder = string(buffer) + builder

		if strings.LastIndex(string(buffer), searchStr) >= 0 {
			break
		}

		readPosition -= bufferSize
	}

	pos := strings.LastIndex(builder, searchStr)
	length := len(builder)
	return builder[pos:length]
}
