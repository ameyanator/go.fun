package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	startTime, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 3:00pm (IST)", loc)
	endTime, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 3:30pm (IST)", loc)

	scheduler := NewScheduler(&BasicStrategy{}, nil)

	meeting1 := NewMeeting(1, startTime, endTime, []string{"Ameya", "Sinha"})
	err := scheduler.schedule(meeting1)
	fmt.Println(err)

	room1 := NewMeetingRoom(1)

	scheduler.addRoom(room1)
	err = scheduler.schedule(meeting1)

	startTime2, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 3:15pm (IST)", loc)
	endTime2, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 3:45pm (IST)", loc)
	meeting2 := NewMeeting(2, startTime2, endTime2, []string{"Ameya", "Sinha"})
	err = scheduler.schedule(meeting2)
	fmt.Println(err)

	room2 := NewMeetingRoom(2)
	scheduler.addRoom(room2)

	err = scheduler.schedule(meeting2)

	startTime3, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 4:15pm (IST)", loc)
	endTime3, _ := time.ParseInLocation(longForm, "Jan 25, 2024 at 4:45pm (IST)", loc)
	meeting3 := NewMeeting(2, startTime3, endTime3, []string{"Ameya", "Sinha"})
	err = scheduler.schedule(meeting3)
	fmt.Println(err)
}
