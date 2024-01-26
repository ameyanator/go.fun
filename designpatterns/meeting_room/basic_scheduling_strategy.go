package main

import (
	"errors"
	"time"
)

type BasicSchedulingStrategy struct {
}

func (b *BasicSchedulingStrategy) scheduleMeeting(startTime, endTime time.Time, meetingRooms []*MeetingRoom) (*MeetingRoom, error) {
	for _, room := range meetingRooms {
		intersection := false
		for _, meeting := range room.meetings {
			if (meeting.startTime.Before(startTime) && startTime.Before(meeting.endTime)) || (meeting.startTime.Before(endTime) && endTime.Before(meeting.endTime)) || (startTime.Before(meeting.startTime) && meeting.startTime.Before(endTime)) || (startTime.Before(meeting.endTime) && meeting.endTime.Before(endTime)) {
				intersection = true
				break
			}
		}
		if intersection {
			continue
		}
		return room, nil
	}
	return nil, errors.New("There are no meeting rooms, which can accomodate this meeting, please try to reschedule!")
}
