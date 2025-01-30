package commandhandler

import (
	"fmt"
	jsonreader "my-golang-app/internal/utils/jsonReader"
)

func AddTask(description string) string {
	if len(description) == 0 {
		return "No task"
	}

	// newTask := task.Task{
	// 	Id:     1,
	// 	Text:   "My custom task",
	// 	Status: "todo",
	// }

	return ""
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
