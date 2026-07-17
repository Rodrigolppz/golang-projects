package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	IsComplete  bool      `json:"is_complete"`
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

	found := false

	if len(tasks) == 0 {
		fmt.Println("This list is empty")
	} else {
		for i := range tasks {
			if ID == tasks[i].ID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println("This task does not exist.")
	}

	return tasks

}

func listalltask(tasks []Task) []Task {

	if len(tasks) == 0 {
		fmt.Println("This list is empty")
	} else {
		fmt.Println(tasks)
	}

	return tasks

}

func listcompletetasks(tasks []Task) []Task {

	if len(tasks) == 0 {
		fmt.Println("This list is empty")
	} else {
		for i := range tasks {
			if tasks[i].IsComplete == true {
				fmt.Println(tasks[i])
			}
		}
	}

	return tasks
}

func main() {
	tasks := []Task{}

	tasks = AddTask(tasks, "Estudar Go")
	tasks = AddTask(tasks, "Arrumar a mesa")
	tasks = AddTask(tasks, "Ler documentação")

	fmt.Println("Depois de adicionar:")
	listalltask(tasks)

	tasks = completetask(tasks, 2)

	fmt.Println("Depois de completar a task 2:")
	listalltask(tasks)

	tasks = deletetask(tasks, 1)

	fmt.Println("Depois de deletar a task 1:")
	listalltask(tasks)
}
