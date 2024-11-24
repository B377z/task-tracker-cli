package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Task represents a single task with all required properties.
type Task struct {
	ID          int       `json:"id"`          // Unique identifier for the task
	Description string    `json:"description"` // Short description of the task
	Status      string    `json:"status"`      // Task status: "todo", "in-progress", "done"
	CreatedAt   time.Time `json:"createdAt"`   // Creation timestamp
	UpdatedAt   time.Time `json:"updatedAt"`   // Last update timestamp
}

// readTasks reads and parses tasks from tasks.json
func readTasks() []Task {
	// Read the file
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		// Handle missing file
		if os.IsNotExist(err) {
			fmt.Println("tasks.json not found. Creating a new file...")
			os.WriteFile("tasks.json", []byte("[]"), 0644)
			return []Task{}
		}
		// Handle other errors
		fmt.Println("Error reading tasks.json:", err)
		os.Exit(1)
	}

	// Parse JSON into a slice of Task
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks.json:", err)
		os.Exit(1)
	}
	return tasks
}

// saveTasks saves tasks to tasks.json
func saveTasks(tasks []Task) {
	// Convert tasks slice to JSON and write to file
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
		os.Exit(1)
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
}

// addTask creates and adds a new task
func addTask(description string) {
	tasks := readTasks()
	id := len(tasks) + 1 // Generate a unique ID
	now := time.Now()

	// Create a new task
	newTask := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Append and save the task
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

// updateTask updates the description of an existing task
func updateTask(id int, description string) {
	tasks := readTasks()
	for i, task := range tasks {
		if task.ID == id {
			// Update task
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task updated successfully (ID: %d)\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

// deleteTask removes a task by its ID
func deleteTask(id int) {
	tasks := readTasks()
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	saveTasks(newTasks)
	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
}

// markTask updates the status of a task
func markTask(id int, status string) {
	if status != "in-progress" && status != "done" {
		fmt.Println("Invalid status. Use 'in-progress' or 'done'.")
		return
	}

	tasks := readTasks()
	for i, task := range tasks {
		if task.ID == id {
			// Update status
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task marked as '%s' (ID: %d)\n", status, id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

// listTasks lists tasks based on the filter
func listTasks(filter string) {
	tasks := readTasks()
	for _, task := range tasks {
		if filter == "all" || filter == task.Status {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
				task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), task.UpdatedAt.Format(time.RFC1123))
		}
	}
}

func main() {
	// Parse command-line arguments
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	// Handle commands
	command := args[0]
	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		addTask(args[1])

	case "update":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli update <id> <description>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		updateTask(id, args[2])

	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		deleteTask(id)

	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		markTask(id, "in-progress")

	case "mark-done":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		markTask(id, "done")

	case "list":
		filter := "all"
		if len(args) > 1 {
			filter = args[1]
		}
		listTasks(filter)

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add, update, delete, mark-in-progress, mark-done, list")
	}
}
