package crawler

import (
	"crawler-gogo/pkg/util"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DURATION = 1 * time.Second
)

func logInput(input Input) {
	fmt.Printf("Start crawler: %#v\n", input)
}

func logOutput(output Output) {
	fmt.Print("End crawler: ")
	if len(output.Content) > 50 {
		output.Content = output.Content[:50] + "...[truncated]"
	}
	fmt.Printf("%#v\n", output)
}

func waitDuration(duration time.Duration) {
	time.Sleep(duration)
}

func createGetRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	util.CheckError(err)
	return req
}

func retrieveResponse(req *http.Request) *http.Response {
	res, err := new(http.Client).Do(req)
	util.CheckError(err)
	return res
}

func readResponseBody(res *http.Response) string {
	defer res.Body.Close()
	bodyByteArray, err := io.ReadAll(res.Body)
	util.CheckError(err)
	return string(bodyByteArray)
}

func Start(input Input) Output {
	logInput(input)

	waitDuration(DURATION)

	req := createGetRequest(input.Url)
	res := retrieveResponse(req)
	resBody := readResponseBody(res)

	output := Output{
		Input:      input,
		FetchedAt:  time.Now(),
		StatusCode: res.StatusCode,
		Content:    resBody,
	}

	logOutput(output)

	return output
}
