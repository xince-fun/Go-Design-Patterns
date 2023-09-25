package command

import "time"

type Command interface {
	Info() string
}

type TimePassrd struct {
	start time.Time
}

func (t *TimePassrd) Info() string {
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}
