package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Todo struct {
	Id     uuid.UUID
	Title  string
	IsDone bool
}

func showWelcome() {
	fmt.Println("########  Welcome to Todo App!  ########")
}
func displayOptions() {
	fmt.Println("\n1) Show all Todos")
	fmt.Println("2) Add a Todo")
	fmt.Println("3) Delete a Todo")
	fmt.Println("4) Mark as Done")
	fmt.Println("5) Exit")
}

var allTodos = []Todo{}

func main() {
	showWelcome()
	displayOptions()

	for {
		reader := bufio.NewReader(os.Stdin)

		option, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error : Invalid Input, Enter correct number")
		}

		numberOption, err := strconv.ParseFloat(strings.TrimSpace(option), 64)

		if err != nil {
			fmt.Println("Error : Invalid Input, Enter correct number")
		}

		switch numberOption {
		case 1:
			showAll()
		case 2:
			fmt.Println("\nEnter Title :")
			handleCrud("add")
		case 3:
			fmt.Println("Enter Todo ID to delete :")
			handleCrud("delete")
		case 4:
			fmt.Println("Enter Todo ID to mark as done :")
			handleCrud("mark")
		case 5:
			fmt.Println("Closing the app")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option, Choose from below options")
			displayOptions()
		}
	}
}

func addTodo(title string) {
	todo := Todo{uuid.New(), title, false}
	allTodos = append(allTodos, todo)
	showAll()
}

func showAll() {
	fmt.Println("\n 🖊 All Todos")
	fmt.Println("\n ---------------------------------------------------------------")
	if len(allTodos) == 0 {
		fmt.Println("\n\tNo Todos available")
		fmt.Println("\n ---------------------------------------------------------------")
	}
	for _, todo := range allTodos {
		fmt.Printf("\tID : %s", todo.Id)
		fmt.Printf("\n\tTITLE : %s", todo.Title)
		fmt.Printf("\n\tCOMPLETED : %s", showCompletedStatus(todo.IsDone))
		fmt.Println("\n ---------------------------------------------------------------")
	}
	displayOptions()
}

func showCompletedStatus(isDone bool) string {
	if isDone {
		return "yes"
	} else {
		return "no"
	}
}

func deleteTodo(id string) {
	for i, todo := range allTodos {
		parsedId, err := uuid.Parse(id)
		if err != nil {
			fmt.Println("Error : Invalid todo Id")
			break
		}
		if todo.Id == parsedId {
			allTodos = append(allTodos[:i], allTodos[i+1:]...)
			break
		}
	}
	showAll()
}

func markAsDone(id string) {
	for i, todo := range allTodos {
		parsedId, err := uuid.Parse(id)
		if err != nil {
			fmt.Println("Error : Invalid todo Id")
			break
		}
		if todo.Id == parsedId {
			todo.IsDone = true
			allTodos[i] = todo
			break
		}
	}
	showAll()
}

func handleCrud(operationType string) {
	inputReader := bufio.NewReader(os.Stdin)

	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("Error : Invalid Input")
		os.Exit(0)
	}

	if operationType == "add" {
		addTodo(strings.TrimSpace(input))
	} else if operationType == "delete" {
		deleteTodo(strings.TrimSpace(input))
	} else if operationType == "mark" {
		markAsDone(strings.TrimSpace(input))
	}
}
