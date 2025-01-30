package jsonwriter

import (
	"encoding/json"
	"my-golang-app/internal/task"
	"os"
)

const pathToData string = "./data/tasks.json"

func WriteJson(tasks []task.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	file, err := os.Create(pathToData)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
