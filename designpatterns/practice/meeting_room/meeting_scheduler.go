package main

import "fmt"

// ability to set meetings among meeting rooms available

type MeetingScheduler struct {
	strategy     SchedulerStrategy
	meetingRooms []*MeetingRoom
}

func NewScheduler(strategy SchedulerStrategy, meetingRooms []*MeetingRoom) *MeetingScheduler {
	return &MeetingScheduler{
		strategy:     strategy,
		meetingRooms: meetingRooms,
	}
}

func (s *MeetingScheduler) schedule(meeting *Meeting) error {
	room, err := s.strategy.scheduleMeeting(meeting.start, meeting.end, s.meetingRooms)
	if err != nil {
		return err
	}
	room.meetings = append(room.meetings, meeting)
	fmt.Println("Meeting Scheduled")
	return nil
}

func (s *MeetingScheduler) addRoom(room *MeetingRoom) {
	s.meetingRooms = append(s.meetingRooms, room)
}
