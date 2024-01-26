package main

type MeetingScheduler struct {
	meetingRooms []*MeetingRoom
	strategy     MeetingSchedulerStrategy
}

func New(meetingRooms []*MeetingRoom, strategy MeetingSchedulerStrategy) *MeetingScheduler {
	return &MeetingScheduler{
		meetingRooms: meetingRooms,
		strategy:     strategy,
	}
}

func (m *MeetingScheduler) addMeetingRoom(meetingRoom *MeetingRoom) {
	m.meetingRooms = append(m.meetingRooms, meetingRoom)
}

func (m *MeetingScheduler) scheduleMeeting(meeting *Meeting) (bool, error) {
	room, err := m.strategy.scheduleMeeting(meeting.startTime, meeting.endTime, m.meetingRooms)
	if err != nil {
		return false, err
	}
	room.occupied = true
	room.meetings = append(room.meetings, meeting)
	return true, nil
}
