package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func writeFile(filename string, content string) {
	file, err := os.Create(filename)
	check(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	check(err)
	writer.Flush()
}

func getDomain(inputURL string) string {
	parsedURL, err := url.Parse(inputURL)
	check(err)
	return parsedURL.Host
}

func generateHashFrom(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
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
	fmt.Printf("Fetched url: %s\n", input.url)
}

func main() {
	urls := loadFile()
	ch := make(chan FetchOutput)
	var wg sync.WaitGroup

	alreadyFetchedDomains := make(map[string]bool)

	for _, url := range urls {
		domain := getDomain(url)
		if alreadyFetchedDomains[domain] {
			fmt.Printf("Skipped url: %s\n", url)
			continue
		}
		alreadyFetchedDomains[domain] = true

		wg.Add(1)
		input := FetchInput{url}
		go fetch(&wg, ch, input)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	dirName := time.Now().Format("20060102")
	os.MkdirAll(dirName, os.ModePerm)
	filenamePrefix := fmt.Sprintf("./%s/", dirName)
	for output := range ch {
		outputFilename := filenamePrefix + generateHashFrom(output.input.url) + ".html"
		writeFile(outputFilename, output.content)
	}
}
