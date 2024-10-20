package main

import (
	"crawler-gogo/pkg/crawler"
	"crawler-gogo/pkg/util"
	"fmt"
)

func appendOutputIntoFile(output crawler.Output) {
	util.AppendInto("./output.log", fmt.Sprintf("%#v\n", output))
}

func main() {
	fmt.Println("Hello, world!")
	output := crawler.Start(crawler.Input{Url: "https://example.com/"})
	appendOutputIntoFile(output)
}
