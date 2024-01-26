package main

type MeetingRoom struct {
	id       int
	occupied bool
	meetings []*Meeting
}

func NewMeetingRoom(id int) *MeetingRoom {
	return &MeetingRoom{
		id:       id,
		occupied: false,
	}
}
