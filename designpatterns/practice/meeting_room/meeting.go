package main

import "time"

type Meeting struct {
	id        int
	start     time.Time
	end       time.Time
	attendees []string
}

func NewMeeting(id int, start, end time.Time, attendees []string) *Meeting {
	return &Meeting{
		id:        id,
		start:     start,
		end:       end,
		attendees: attendees,
	}
}
