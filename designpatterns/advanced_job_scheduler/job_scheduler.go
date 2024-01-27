package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id      int
	execute func()
}

type Scheduler struct {
	totalTasks    int
	taskMutex     sync.Mutex
	taskStopChan  map[*Task]chan struct{}
	stopChanMutex sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		totalTasks:   0,
		taskStopChan: make(map[*Task]chan struct{}),
	}
}

func (s *Scheduler) schedule(task *Task, offset time.Time) {
	s.taskMutex.Lock()
	s.totalTasks++
	task.id = s.totalTasks
	s.taskMutex.Unlock()

	s.stopChanMutex.Lock()
	stopChannel := make(chan struct{})
	s.taskStopChan[task] = stopChannel
	s.stopChanMutex.Unlock()

	interval := offset.Sub(time.Now())
	if interval <= 0 {
		go s.executeTasks(task)
		return
	}

	go func() {
		select {
		case <-time.After(interval):
			go s.executeTasks(task)
		case <-stopChannel:
			fmt.Println("Before execution got signal to stop")
			return
		}
	}()
}

func (s *Scheduler) scheduleAtFixedInterval(task *Task, interval time.Duration) {
	s.taskMutex.Lock()
	s.totalTasks++
	task.id = s.totalTasks
	s.taskMutex.Unlock()

	s.stopChanMutex.Lock()
	stopChannel := make(chan struct{})
	s.taskStopChan[task] = stopChannel
	s.stopChanMutex.Unlock()

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				go s.executeTasks(task)
			case <-stopChannel:
				fmt.Println("fixed interval scheduling stopping now")
				return
			}
		}
	}()
}

func (s *Scheduler) executeTasks(task *Task) {
	fmt.Println("Executing task with id ", task.id)
	task.execute()
}

func main() {
	scheduler := NewScheduler()

	task1 := Task{
		execute: func() {
			fmt.Println("Connecting to Uber")
			time.Sleep(10 * time.Second)
			fmt.Println("Connected to Uber")
		},
	}

	// scheduler.schedule(&task1, time.Now().Add(5*time.Second))
	// fmt.Println("Sleep for 2 secs")
	// time.Sleep(2 * time.Second)
	// stopChan := scheduler.taskStopChan[&task1]
	// close(stopChan)

	scheduler.scheduleAtFixedInterval(&task1, 2*time.Second)

	time.Sleep(15 * time.Second)

	stopChan := scheduler.taskStopChan[&task1]
	close(stopChan)

	time.Sleep(2 * time.Minute)
}
