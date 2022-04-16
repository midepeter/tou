package main 

import (
	"time"
)

type Timer struct {
	current string
}

func NewTimer() *Timer {
	return &Timer{
		current: time.Now().Format("15:04:03\n"),
	}
}