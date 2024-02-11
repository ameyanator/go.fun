package main

import "time"

type SchedulerStrategy interface {
	scheduleMeeting(time.Time, time.Time, []*MeetingRoom) (*MeetingRoom, error)
}
