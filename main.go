package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type FetchInput struct {
	url string
}

type FetchOutput struct {
	input   FetchInput
	content string
}

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

func fetch(wg *sync.WaitGroup, ch chan<- FetchOutput, input FetchInput) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	req, err := http.NewRequest("GET", input.url, nil)
	check(err)
	client := new(http.Client)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()
	byteArray, err := io.ReadAll(res.Body)
	check(err)
	ch <- FetchOutput{input, string(byteArray)}
}

func main() {
	urls := loadFile()
	ch := make(chan FetchOutput)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		input := FetchInput{url}
		go fetch(&wg, ch, input)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
}
