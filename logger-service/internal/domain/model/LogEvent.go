package model

import "time"

type LogEvent struct {
	AppName string
	Level   int
	Time    time.Time
	Text    string
}
