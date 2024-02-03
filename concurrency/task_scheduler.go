package main

import "sync"

type Task struct {
	id      int
	message string
}

type TaskScheduler struct {
	tasks          []*Task
	mu             sync.Mutex
	tasksAvailable *sync.Cond
	completedTasks int
	totalTasks     int
	finished       bool
}

func NewScheduler(totalTasks int) *TaskScheduler {
	s := &TaskScheduler{totalTasks: totalTasks}
	s.tasksAvailable = sync.NewCond(&s.mu)

	return s
}

func (s *TaskScheduler) SubmitTask(task *Task) {
	s.mu.Lock()
	s.tasks = append(s.tasks, task)
	s.tasksAvailable.Signal()
	s.mu.Unlock()
}

func (s *TaskScheduler) GetTask() *Task {
	s.mu.Lock()
	for len(s.tasks) == 0 && !s.finished {
		s.tasksAvailable.Wait()
	}
	
}
