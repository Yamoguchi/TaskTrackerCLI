package main

import (
	"fmt"
	Command "my-golang-app/internal/commandHandler"
	"os"
	"path/filepath"
	"strconv"
)

const pathToData string = "./data/tasks.json"

func main() {

	result := checkTasksJsonFileIsExist()

	if !result {
		err := createEmptyJsonFile()
		if err != nil {
			fmt.Println("Произошла ошибка создания tasks.json", err)
		}
	}

	// Getting command from arguments
	if len(os.Args) < 2 {
		fmt.Println("Вы не указали аргументы. Воспользуйтесь подсказкой")
	} else {
		command := os.Args[1]
		invokeCommand(command)
	}
}

func invokeCommand(command string) {
	switch command {
	case "add":
		Command.AddTask(os.Args[2])

	case "update":
		taskId, err := strconv.Atoi((os.Args[2]))

		if err == nil {
			Command.UpdateTask(taskId, os.Args[3])
		} else {
			fmt.Println("Ошибка ввода")
		}

	case "delete":
		taskId, err := strconv.Atoi((os.Args[2]))

		if err == nil {
			Command.DeleteTask(taskId)
		} else {
			fmt.Println("Ошибка ввода")
		}

		fmt.Println("delete command")

	case "list":
		if !(len(os.Args) < 3) {
			if os.Args[2] == "done" {
				Command.ListDoneTasks()
			} else if os.Args[2] == "todo" {
				Command.ListTodoTasks()
			} else if os.Args[2] == "in-progress" {
				Command.ListInProgressTasks()
			} else {
				fmt.Println("Wrong command. Check tutotial")
			}
		} else {
			Command.ListTasks(nil)
		}

	case "mark-in-progress":
		taskId, err := strconv.Atoi((os.Args[2]))

		if err == nil {
			Command.MarkTaskInProgress(taskId)
		} else {
			fmt.Println("Ошибка ввода")
		}

	case "mark-done":
		taskId, err := strconv.Atoi((os.Args[2]))

		if err == nil {
			Command.MarkTaskDone(taskId)
		} else {
			fmt.Println("Ошибка ввода")
		}

	default:
		fmt.Println("Wrong command")
	}
}

func checkTasksJsonFileIsExist() bool {
	_, err := os.Stat(pathToData)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println("Ошибка доступа")
		return false
	}
	return true
}

func createEmptyJsonFile() error {
	dir := filepath.Dir(pathToData)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(pathToData)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("{}")
	if err != nil {
		return err
	}

	return nil
}
