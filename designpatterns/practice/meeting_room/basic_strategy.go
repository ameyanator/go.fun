package main

import (
	"errors"
	"time"
)

type BasicStrategy struct {
}

func (b *BasicStrategy) scheduleMeeting(start, end time.Time, rooms []*MeetingRoom) (*MeetingRoom, error) {
	for _, room := range rooms {
		possible := true
		for _, meeting := range room.meetings {
			if (start.Before(meeting.start) && meeting.start.Before(end)) || (start.Before(meeting.end) && meeting.end.Before(end)) || (meeting.start.Before(start) && start.Before(meeting.end)) || (meeting.start.Before(end) && end.Before(meeting.end)) {
				possible = false
				break
			}
		}
		if possible {
			return room, nil
		}
	}
	return nil, errors.New("There are no meeting rooms which can at specified time, pls try during a different slot")
}
