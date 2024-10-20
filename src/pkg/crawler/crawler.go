package crawler

import (
	"fmt"
	"time"
)

const (
	DURATION = 1 * time.Second
)

func waitDuration(duration time.Duration) {
	time.Sleep(duration)
}

func Start(input Input) {
	fmt.Printf("Start crawler: %#v\n", input)
	waitDuration(DURATION)
}
