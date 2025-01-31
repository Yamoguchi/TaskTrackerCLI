package jsonreader

import "my-golang-app/internal/task"

func CheckLastTaskId(tasks []task.Task) int {
	if tasks == nil {
		return -1
	} else {
		return tasks[len(tasks)-1].Id
	}
}

func FindTaskById(taskId int, tasks []task.Task) int {
	for i := range tasks {
		if tasks[i].Id == taskId {
			return i
		}
	}
	return -1
}

func FindTasksByStatus(status string, tasks []task.Task) (result []int) {
	for i := range tasks {
		if tasks[i].Status == status {
			result = append(result, i)
		}
	}
	return result
}
