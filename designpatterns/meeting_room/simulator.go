package main

import (
	"fmt"
	"time"
)

func main() {
	strategy1 := &BasicSchedulingStrategy{}
	scheduler := New(nil, strategy1)
	loc, _ := time.LoadLocation("Asia/Kolkata")

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	startTime, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 3:00pm (IST)", loc)
	endTime, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 3:30pm (IST)", loc)
	meeting := NewMeeting(1, startTime, endTime)

	_, err := scheduler.scheduleMeeting(meeting)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Meeting Scheduled")
	}

	meetingRoom1 := NewMeetingRoom(1)

	scheduler.addMeetingRoom(meetingRoom1)

	_, err = scheduler.scheduleMeeting(meeting)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Meeting Scheduled")
	}

	newEndTime, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 3:45pm (IST)", loc)
	meeting2 := NewMeeting(2, startTime, newEndTime)
	_, err = scheduler.scheduleMeeting(meeting2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Meeting Scheduled")
	}

	meetingRoom2 := NewMeetingRoom(2)
	scheduler.addMeetingRoom(meetingRoom2)
	
	_, err = scheduler.scheduleMeeting(meeting2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Meeting Scheduled")
	}
}
