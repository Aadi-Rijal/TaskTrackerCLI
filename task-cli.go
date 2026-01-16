package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [args]")
		return
	}

	command := os.Args[1]
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add \"description\"")
			return
		}
		tasks, task := AddTask(tasks, os.Args[2])
		SaveTasks(tasks)
		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> \"description\"")
			return
		}
		id, err := ParseID(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		tasks, ok := UpdateTask(tasks, id, os.Args[3])
		if !ok {
			fmt.Println("Task not found")
			return
		}
		SaveTasks(tasks)
		fmt.Println("Task updated successfully")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}

		id, err := ParseID(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		tasks, ok := DeleteTask(tasks, id)
		if !ok {
			fmt.Println("Task not found")
			return
		}
		SaveTasks(tasks)
		fmt.Println("Task deleted successfully")

	case "mark-in-progress":
		id, _ := ParseID(os.Args[2])
		tasks, ok := MarkTask(tasks, id, "inprogress")
		if !ok {
			fmt.Println("Task not found")
			return
		}
		SaveTasks(tasks)
		fmt.Println("Task marked as in-progress")

	case "mark-done":
		id, _ := ParseID(os.Args[2])
		tasks, ok := MarkTask(tasks, id, "done")
		if !ok {
			fmt.Println("Task not found")
			return
		}
		SaveTasks(tasks)
		fmt.Println("Task marked as done")

	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		ListTasks(tasks, filter)

	default:
		fmt.Println("Unknown command")
	}
}
