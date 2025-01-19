package main

import (
	"encoding/json"
	"time"
	//"log"
	"fmt"
	"os"
)

const FileName = "tasks.json"

var command = os.Args[1]
var argument = os.Args[2]

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

	fmt.Println("Task added successfully.")
}

func deleteTask() {
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./main <command> \"<argument>\"")
		return
	}

    readTasksFromFile()


    switch command {
        case "add":
        addTask()
        case "delete":
        deleteTask()
    }
}
