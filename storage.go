package main

import (
	"encoding/json"
	"os"
)

const dataFile = "tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if len(data) == 0 {
		return []Task{}, nil
	}

	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}
