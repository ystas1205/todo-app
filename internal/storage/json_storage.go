package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-app/internal/todo"
)

// SaveJSON сериализует список задач в JSON и записывает в файл path.
// Если файл существует, он перезаписывается.
func SaveJSON(path string, tasks []todo.Task) error {
    jsonBytes, errSerialization := json.MarshalIndent(tasks, "", "  ")
    if errSerialization != nil {
        return fmt.Errorf("ошибка сериализации: %w", errSerialization)
    }

    if errWrite := os.WriteFile(path, jsonBytes, 0644); errWrite != nil {
        return fmt.Errorf("ошибка записи в файл: %w", errWrite)
    }

    return nil
}

// LoadJSON читает задачи из JSON-файла.
// Если файла нет, создаёт пустой файл и возвращает пустой срез.
func LoadJSON(path string) ([]todo.Task, error) {

	tasks := []todo.Task{}
	
	_, errFileExists := os.Stat(path)
	if errFileExists != nil {

		if errSave := SaveJSON(path, tasks); errSave != nil {
			return nil, fmt.Errorf("ошибка создания файла %w", errSave)
		}
	}

	data, errRead := os.ReadFile(path)
	if errRead != nil {
		return nil, fmt.Errorf("ошибка чтения файла %w", errRead)
	}

	if errParse := json.Unmarshal(data, &tasks); errParse != nil {
		return nil, fmt.Errorf("ошибка десириализации файла %w", errParse)
	}
	return tasks, nil
}
