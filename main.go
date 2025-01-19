package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const FileName = "tasks.json"

var command = os.Args[1]
var argument string

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var tasks = []Task{}

func writeTasksToFile() {
	file, err := os.Create(FileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	formatted, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	_, err = file.Write(formatted)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func readTasksFromFile() {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func addTask() {
	newTask := Task{
		ID:          int64(len(tasks) + 1),
		Description: argument,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	tasks = append(tasks, newTask)
	writeTasksToFile()

	fmt.Printf("Task added successfully (ID: %d).\n", newTask.ID)
}

func deleteTask() {
	idToDelete, err := strconv.ParseInt(argument, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == idToDelete {
			tasks = append(tasks[:i], tasks[i+1:]...)
			writeTasksToFile()
			fmt.Println("Task deleted successfully.")
			return
		}
	}

	fmt.Println("Task not found.")
}

func updateTask() {
	descriptionToUpdate := os.Args[3]
	idToUpdate, err := strconv.ParseInt(argument, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == idToUpdate {
			tasks[i].Description = descriptionToUpdate
			writeTasksToFile()
			fmt.Println("Task updated successfully.")
			return
		}
	}

	fmt.Println("Task not found.")
}

func listTasks() {
	formatted, _ := json.MarshalIndent(tasks, "", "  ")
	fmt.Println("Tasks: ")
	fmt.Println()
	fmt.Println(string(formatted))
    fmt.Println()
}

func markInProgress() {
	idForProgress, err := strconv.ParseInt(argument, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == idForProgress {
			tasks[i].Status = "todo"
			writeTasksToFile()
			fmt.Println("Task marked as todo.")
			return
		} else if task.Status == "todo" {
			fmt.Println("Task already in progress.")
		}
	}

	fmt.Println("Task not found.")
}

func markAsDone() {
	idForDone, err := strconv.ParseInt(argument, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == idForDone {
			tasks[i].Status = "done"
			writeTasksToFile()
			fmt.Println("Task marked as done.")
			return
		} else if task.Status == "done" {
			fmt.Println("Task already done.")
		}
	}

	fmt.Println("Task not found.")
}

func listTodoTasks() {
	for _, task := range tasks {
		if task.Status == "todo" {
			fmt.Println("ID:", task.ID)
			fmt.Println("Description:", task.Description)
			fmt.Println("Created at:", task.CreatedAt)
			fmt.Println("Updated at:", task.UpdatedAt)
		}
	}
}

func listDoneTasks() {
	for _, task := range tasks {
		if task.Status == "done" {
			fmt.Println("ID:", task.ID)
			fmt.Println("Description:", task.Description)
			fmt.Println("Created at:", task.CreatedAt)
			fmt.Println("Updated at:", task.UpdatedAt)
		}
	}
}

func help() {
	fmt.Println("Usage: ./main <command> [<argument>]")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  add <description>            Add a new task with a description")
	fmt.Println("  delete <id>                  Delete a task by ID")
	fmt.Println("  update <id> <new description> Update task description by ID")
	fmt.Println("  list                         List all tasks")
	fmt.Println("  mark-in-progress <id>        Mark task as 'todo' by ID")
	fmt.Println("  mark-done <id>               Mark task as 'done' by ID")
	fmt.Println("  list-done                    List all tasks that are done")
	fmt.Println("  list-todo                    List all tasks that are todo")
	fmt.Println("  --help                       Show this help message")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		help()
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <command> [<argument>]")
		return
	}

	command = os.Args[1]
	if len(os.Args) > 2 {
		argument = os.Args[2]
	}

	readTasksFromFile()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./main add <description>")
			return
		}
		addTask()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./main delete <id>")
			return
		}
		deleteTask()
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: ./main update <id> <new description>")
			return
		}
		updateTask()
	case "list":
		listTasks()
	case "mark-in-progress":
		markInProgress()
	case "mark-done":
		markAsDone()
	case "list-done":
		listDoneTasks()
	case "list-todo":
		listTodoTasks()
	default:
		fmt.Println("Unknown command:", command)
	}
}
