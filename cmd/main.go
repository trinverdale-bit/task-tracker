package main

import (
	"fmt"
	"os"
	"github.com/trinverdale-bit/task-tracker/internal/task"
	"github.com/trinverdale-bit/task-tracker/internal/file"
)

var command = os.Args[1]
var argument string

func help() {
	fmt.Println("Usage: ./task-tracker <command> [<argument>]")
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

	tasks, err := file.ReadTasksFromFile()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./main add <description>")
			return
		}
		task.AddTask(tasks, argument)
		file.WriteTasksToFile(tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./main delete <id>")
			return
		}
		task.DeleteTask(tasks, argument)
		file.WriteTasksToFile(tasks)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: ./main update <id> <new description>")
			return
		}
		task.UpdateTask(tasks, argument, os.Args[3])
		file.WriteTasksToFile(tasks)
	case "list":
		task.ListTasks(tasks)
	case "mark-in-progress":
		task.MarkInProgress(tasks, argument)
		file.WriteTasksToFile(tasks)
	case "mark-done":
		task.MarkAsDone(tasks, argument)
		file.WriteTasksToFile(tasks)
	case "list-done":
		task.ListDoneTasks(tasks)
	case "list-todo":
		task.ListTodoTasks(tasks)
	default:
		fmt.Println("Unknown command:", command)
	}
}
