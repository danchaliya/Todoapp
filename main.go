package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID     int
	Task   string
	Done   bool
}

var todos []Todo
var currentID int

func main() {
	fmt.Println("Todo App")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nPlease select an option:")
		fmt.Println("1. Add Todo")
		fmt.Println("2. List Todos")
		fmt.Println("3. Mark Todo as Done")
		fmt.Println("4. Exit")

		scanner.Scan()
		option, _ := strconv.Atoi(scanner.Text())

		switch option {
		case 1:
			fmt.Println("Enter the task description:")
			scanner.Scan()
			task := scanner.Text()

			addTodo(task)
			fmt.Println("Todo added successfully.")

		case 2:
			listTodos()

		case 3:
			fmt.Println("Enter the ID of the todo to mark as done:")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			err := markTodoAsDone(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Todo marked as done.")
			}

		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addTodo(task string) {
	currentID++
	todo := Todo{
		ID:     currentID,
		Task:   task,
		Done:   false,
	}
	todos = append(todos, todo)
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	fmt.Println("Todos:")
	for _, todo := range todos {
		status := "Not Done"
		if todo.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s (%s)\n", todo.ID, todo.Task, status)
	}
}

func markTodoAsDone(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
