package main

import "time"

type Meeting struct {
	id          int
	startTime   time.Time
	endTime     time.Time
	paritipants []string
	topic       string
}

func NewMeeting(id int, start, end time.Time) *Meeting {
	return &Meeting{
		id:        id,
		startTime: start,
		endTime:   end,
	}
}
