package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id     int
	doWork func()
}

type Scheduler struct {
	jobChannel chan *Job
	wg         sync.WaitGroup
}

func NewScheduler(workers int) *Scheduler {
	scheduler := Scheduler{
		jobChannel: make(chan *Job),
	}
	for i := 0; i < workers; i++ {
		go scheduler.executeTasks()
	}
	return &scheduler
}

func (s *Scheduler) executeTasks() {
	for {
		select {
		case job, ok := <-s.jobChannel:
			if !ok {
				fmt.Println("Job channel has been closed, stopping this goroutine")
				return
			}
			fmt.Println("Executing job with id ", job.id)
			job.doWork()
			fmt.Println("Job ", job.id, " has been completed")
			s.wg.Done()
		}
	}
}

func (s *Scheduler) acceptTasks(job *Job) {
	s.wg.Add(1)
	s.jobChannel <- job
}

func main() {
	job1 := Job{
		id: 1,
		doWork: func() {
			fmt.Println("Job 1 in progress")
			time.Sleep(5 * time.Second)
		},
	}

	scheduler := NewScheduler(2)
	scheduler.acceptTasks(&job1)

	job2 := Job{
		id: 2,
		doWork: func() {
			fmt.Println("Job 2 in progress")
			time.Sleep(1 * time.Second)
		},
	}
	scheduler.acceptTasks(&job2)
	scheduler.wg.Wait()
}
