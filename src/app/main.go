package main

import (
	"crawler-gogo/pkg/crawler"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	crawler.Start(crawler.Input{Url: "https://example.com/"})
}
