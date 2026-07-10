package main

import (
	"time"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

func AddTask(tasks []Task, description string) []Task {

	newID := 0

	if len(tasks) == 0 {

		newID += 1

	} else {

		newID = tasks[len(tasks)-1].ID + 1
	}

	newtask := Task{
		ID:          newID,
		Description: description,
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, newtask)

}

func main() {

}
