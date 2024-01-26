package main

import "time"

type MeetingSchedulerStrategy interface {
	scheduleMeeting(startTime time.Time, endTime time.Time, meetingRooms []*MeetingRoom) (*MeetingRoom, error)
}
