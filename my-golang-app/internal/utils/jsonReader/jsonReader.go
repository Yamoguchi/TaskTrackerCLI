package jsonreader

import (
	"encoding/json"
	"fmt"
	"io"
	"my-golang-app/internal/task"
	"os"
)

const pathToData string = "./data/tasks.json"

func ReadJson() ([]task.Task, error) {
	file, err := os.Open(pathToData)

	// Checking that a file exists and can be opened
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}

	defer file.Close()

	// Checking that a file is not empty
	isEmpty, err := isJsonFileEmpty(file)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return nil, nil
	}

	var tasks []task.Task

	// Deserializing json file
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Println("Ошибка десериализации json файла")
		return nil, err
	}

	return tasks, nil
}

func isJsonFileEmpty(file *os.File) (bool, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return false, err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}

	return string(content) == "{}", nil
}
