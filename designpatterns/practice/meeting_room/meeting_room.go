package main

type MeetingRoom struct {
	id       int
	meetings []*Meeting
}

func NewMeetingRoom(id int) *MeetingRoom {
	return &MeetingRoom{
		id: id,
	}
}
