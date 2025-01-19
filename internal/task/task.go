package task

import (
	"fmt"
	"time"
	"strconv"
)

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func AddTask(tasks []Task, description string) []Task {
	newTask := Task{
		ID:          int64(len(tasks) + 1),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	return append(tasks, newTask)
}

func DeleteTask(tasks []Task, idStr string) []Task {
	idToDelete, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return tasks
	}

	for i, task := range tasks {
		if task.ID == idToDelete {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}

	fmt.Println("Task not found.")
	return tasks
}

func UpdateTask(tasks []Task, idStr, newDescription string) []Task {
	idToUpdate, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return tasks
	}

	for i, task := range tasks {
		if task.ID == idToUpdate {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			return tasks
		}
	}

	fmt.Println("Task not found.")
	return tasks
}

func ListTasks(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

func MarkInProgress(tasks []Task, idStr string) []Task {
	idForProgress, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return tasks
	}

	for i, task := range tasks {
		if task.ID == idForProgress {
			tasks[i].Status = "todo"
			return tasks
		}
	}

	fmt.Println("Task not found.")
	return tasks
}

func MarkAsDone(tasks []Task, idStr string) []Task {
	idForDone, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return tasks
	}

	for i, task := range tasks {
		if task.ID == idForDone {
			tasks[i].Status = "done"
			return tasks
		}
	}

	fmt.Println("Task not found.")
	return tasks
}

func ListDoneTasks(tasks []Task) {
	for _, task := range tasks {
		if task.Status == "done" {
			fmt.Printf("ID: %d, Description: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.ID, task.Description, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func ListTodoTasks(tasks []Task) {
	for _, task := range tasks {
		if task.Status == "todo" {
			fmt.Printf("ID: %d, Description: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.ID, task.Description, task.CreatedAt, task.UpdatedAt)
		}
	}
}
