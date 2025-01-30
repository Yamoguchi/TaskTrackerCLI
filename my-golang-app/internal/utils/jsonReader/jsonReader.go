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

	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}

	defer file.Close()

	isEmpty, err := isJsonFileEmpty(file)
	if err != nil {
		return nil, err
	}

	if isEmpty {
		return nil, fmt.Errorf("json файл пуст")
	}

	var tasks []task.Task

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
		return true, err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return true, err
	}

	return string(content) == "{}", nil
}
