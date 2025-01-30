package jsonreader

import "my-golang-app/internal/task"

func CheckLastTaskId(tasks []task.Task) int {
	return tasks[len(tasks)-1].Id
}
