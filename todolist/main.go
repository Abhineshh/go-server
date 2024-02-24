package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int
	Name   string
	IsDone bool
}

var tasks []Task
var nextID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n Options : [list | add <task> | done <taskID> | delete <taskID> | exit]")
		fmt.Println("Enter Command:")
		cmdString, _ := reader.ReadString('\n')
		cmdString = strings.TrimSpace(cmdString)
		cmdParts := strings.SplitN(cmdString," ", 2)
		fmt.Println(cmdString,cmdParts)
		switch cmdParts[0] {
		case "list":
			listTasks()
		case "add":
			if len(cmdParts) < 2 {
				fmt.Println("Error : no task Provided")
				continue
			}
			addTask(cmdParts[1])
		case "done":
			if len(cmdParts) < 2 {
				fmt.Println("Error: no task ID Provided")
				continue
			}
			taskID, err := strconv.Atoi(cmdParts[1])
			if err != nil {
				fmt.Println("invalid task Id")
				continue
			}
			doneTask(taskID)
		case "delete":
			if len(cmdParts) < 2 {
				fmt.Println("Error: no task ID Provided")
				continue
			}
			taskID, err := strconv.Atoi(cmdParts[1])
			if err != nil {
				fmt.Println("invalid task Id")
				continue
			}
			deleteTask(taskID)
		case "exit":
			return
		default:
			fmt.Println("Invalid Command")
		}
	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to display")
		return
	}
	for _, task := range tasks {
		status := "Pending"
		if task.IsDone {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
	}
}

func addTask(name string) {
	tasks = append(tasks, Task{ID: nextID, Name: name, IsDone: false})
	nextID++
	fmt.Println("Added Task:", name)
}

func doneTask(taskId int) {
	for i, task := range tasks {
		if task.ID == taskId {
			tasks[i].IsDone = true
			fmt.Printf("Task %d marked as done \n", taskId)
			return
		}
	}
	fmt.Println("task not found")
}

func deleteTask(taskId int) {
	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("task %d deleted \n", taskId)
			return
		}
	}
	fmt.Println("task not found")
}
