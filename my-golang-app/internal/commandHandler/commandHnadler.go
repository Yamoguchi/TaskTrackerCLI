package commandhandler

import (
	"fmt"
	"my-golang-app/internal/task"
	jsonreader "my-golang-app/internal/utils/jsonReader"
	jsonwriter "my-golang-app/internal/utils/jsonWriter"
)

func AddTask(description string) {
	if len(description) == 0 {
		fmt.Println("Не указана новая задача")
		return
	}

	tasks, err := jsonreader.ReadJson()

	if err != nil {
		fmt.Print(err)
		return
	}

	newTask := task.Task{
		Id:     jsonreader.CheckLastTaskId(tasks) + 1,
		Text:   description,
		Status: "todo",
	}

	tasks = append(tasks, newTask)

	jsonwriter.WriteJson(tasks)
}

func UpdateTask(taskId int, newTask string) {

}

func DeleteTask(taskId int) {

}

func ListTasks() {
	allTasks, err := jsonreader.ReadJson()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allTasks)
}

func ListDoneTasks() {

}

func ListTodoTasks() {

}

func ListInProgressTasks() {

}

func MarkTaskInProgress(taskId int) {

}

func MarkTaskDone(taskId int) {

}
