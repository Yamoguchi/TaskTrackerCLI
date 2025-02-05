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
	if len(newTask) == 0 || taskId < 0 {
		fmt.Println("You have specified an incorrect parameter")
		return
	}

	tasks, err := jsonreader.ReadJson()
	if err != nil {
		fmt.Println(err)
		return
	}

	idOfSearchingTask := jsonreader.FindTaskById(taskId, tasks)
	if idOfSearchingTask == -1 {
		fmt.Println("Such a task does not exist")
		return
	}

	tasks[idOfSearchingTask].UpdateTaskText(newTask)

	jsonwriter.WriteJson(tasks)
}

func DeleteTask(taskId int) {

	if taskId < 0 {
		fmt.Println("You have specified an incorrect parameter")
		return
	}

	tasks, err := jsonreader.ReadJson()
	if err != nil {
		fmt.Println(err)
		return
	}

	idOfSearchingTask := jsonreader.FindTaskById(taskId, tasks)

	jsonwriter.WriteJson(append(tasks[:idOfSearchingTask], tasks[idOfSearchingTask+1]))
}

func ListTasks(tasks []task.Task) {
	if tasks == nil {
		var err error
		tasks, err = jsonreader.ReadJson()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("%-5s\t%-30s\t%-15s\n", "ID", "Task", "Status")
	fmt.Println("--------------------------------------------------------")

	// Вывод задач
	for _, task := range tasks {
		fmt.Printf("%-5d\t%-30s\t%-15s\n", task.Id, task.Text, task.Status)
	}
}

func MarkTaskInProgress(taskId int) {
	if taskId < 0 {
		fmt.Println("You have specified an incorrect parameter")
		return
	}

	tasks, err := jsonreader.ReadJson()
	if err != nil {
		fmt.Println(err)
		return
	}

	idOfSearchingTask := jsonreader.FindTaskById(taskId, tasks)
	tasks[idOfSearchingTask].UpdateTaskStatus("in-progress")
}

func MarkTaskDone(taskId int) {
	if taskId < 0 {
		fmt.Println("You have specified an incorrect parameter")
		return
	}

	tasks, err := jsonreader.ReadJson()
	if err != nil {
		fmt.Println(err)
		return
	}

	idOfSearchingTask := jsonreader.FindTaskById(taskId, tasks)
	tasks[idOfSearchingTask].UpdateTaskStatus("done")

}

func ListTasksByStatus(status string) {
	tasks, err := jsonreader.ReadJson()
	if err != nil {
		fmt.Println(err)
		return
	}

	searchIds := jsonreader.FindTasksByStatus(status, tasks)

	var searchTasks []task.Task

	for _, i := range searchIds {
		searchTasks = append(searchTasks, tasks[i])
	}

	ListTasks(searchTasks)
}
