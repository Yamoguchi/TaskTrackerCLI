package jsonreader

import "my-golang-app/internal/task"

func CheckLastTaskId(tasks []task.Task) int {
	if tasks == nil {
		return 1
	} else {
		return tasks[len(tasks)-1].Id
	}
}
