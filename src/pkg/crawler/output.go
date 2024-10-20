package crawler

import "time"

type Output struct {
	Input      Input
	FetchedAt  time.Time
	StatusCode int
	Content    string
}
