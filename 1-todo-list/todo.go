package main

import (
	"fmt"
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
		IsComplete:  false,
	}

	tasks = append(tasks, newtask)

	return tasks

}

func completetask(tasks []Task, ID int) []Task {

	var taskID int

	if len(tasks) == 0 {

		fmt.Println("Task list is empty")

	} else {

		for i, v := range tasks {
			if taskID == ID {
				tasks = append(tasks)

			}
		}

	}

}

func main() {

}
