package main

import (
	"fmt"
	Command "my-golang-app/pkg/commandHandler"
	"os"
	"strconv"
)

func main() {
	command := os.Args[1]

	switch command {
	case "add":
		fmt.Println("Add command")
		Command.AddTask(os.Args[2])

	case "update":
		fmt.Println("Update command")
		taskId, _ := strconv.Atoi((os.Args[2]))
		Command.UpdateTask(taskId, os.Args[3])

	case "delete":
		taskId, _ := strconv.Atoi((os.Args[2]))
		Command.DeleteTask(taskId)
		fmt.Println("delete command")

	case "list":
		if !(len(os.Args) < 3) {
			if os.Args[2] == "done" {
				fmt.Println("list done")
				Command.ListDoneTasks()
			} else if os.Args[2] == "todo" {
				fmt.Println("List todo")
				Command.ListTodoTasks()
			} else if os.Args[2] == "in-progress" {
				fmt.Println("List in-progress")
				Command.ListInProgressTasks()
			} else {
				fmt.Println("Wrong command. Check tutotial")
			}
		} else {
			fmt.Println("List command")
			Command.ListTasks()
		}

	case "mark-in-progress":
		fmt.Println("Mark-in-progress command")
		taskId, _ := strconv.Atoi((os.Args[2]))
		Command.MarkTaskInProgress(taskId)

	case "mark-done":
		fmt.Println("Mark-done command")
		taskId, _ := strconv.Atoi((os.Args[2]))
		Command.MarkTaskDone(taskId)

	default:
		fmt.Println("Wrong command")
	}
}
