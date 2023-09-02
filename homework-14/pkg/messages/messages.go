package messages

import "time"

type Message struct {
	ID      int
	Time    time.Time
	Content string
}
