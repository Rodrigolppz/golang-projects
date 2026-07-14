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

	found := false

	if len(tasks) == 0 {

		fmt.Println("Task list is empty")

	} else {
		for i := range tasks {
			if ID == tasks[i].ID {
				tasks[i].IsComplete = true
				found = true
			}
		}
	}

	if found == false {
		fmt.Println("This task does not exist")
	}

	return tasks
}

func deletetask(tasks []Task, ID int) []Task {

	if len(tasks) == 0 {
		fmt.Println("This list is empty")
	} else {
		for i := range tasks {
			if ID == tasks[i].ID {

			}
		}
	}

}

func main() {

}
