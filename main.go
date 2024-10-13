package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INPUT_FILE_NAME = "url_list.txt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFile() []string {
	file, err := os.Open(INPUT_FILE_NAME)
	check(err)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	check(scanner.Err())
	return lines
}

func main() {
	urls := loadFile()
	for _, url := range urls {
		fmt.Println(url)
	}
}
