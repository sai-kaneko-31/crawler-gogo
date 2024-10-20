package util

import (
	"bufio"
	"os"
)

func AppendInto(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	CheckError(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	CheckError(err)
	writer.Flush()
}
