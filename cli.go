package main

import (
	"fmt"
	"strconv"
	"time"
)

func ParseID(arg string) (int, error) {
	return strconv.Atoi(arg)
}

func NextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func AddTask(tasks []Task, desc string) ([]Task, Task) {
	now := time.Now().Format(time.RFC3339)
	task := Task{
		ID:          NextID(tasks),
		Description: desc,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return append(tasks, task), task
}

func UpdateTask(tasks []Task, id int, desc string) ([]Task, bool) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			return tasks, true
		}
	}
	return tasks, false
}

func DeleteTask(tasks []Task, id int) ([]Task, bool) {
	for i, t := range tasks {
		if t.ID == id {
			return append(tasks[:i], tasks[i+1:]...), true
		}
	}
	return tasks, false
}

func MarkTask(tasks []Task, id int, status string) ([]Task, bool) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			return tasks, true
		}
	}
	return tasks, false
}

func ListTasks(tasks []Task, filter string) {
	for _, t := range tasks {
		if filter == "" || t.Status == filter {
			fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
		}
	}
}
