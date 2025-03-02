package file

import (
	"encoding/json"
	"fmt"
	"github.com/trinverdale-bit/task-tracker/internal/task"
	"os"
)

const FileName = "tasks.json"

func ReadTasksFromFile() ([]task.Task, error) {
	var tasks []task.Task

	file, err := os.Open(FileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []task.Task) error {
	file, err := os.Create(FileName)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	formatted, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting JSON: %w", err)
	}

	_, err = file.Write(formatted)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
