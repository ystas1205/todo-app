package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"todo-app/internal/todo"
)

// SaveJSON сериализует список задач в JSON и записывает в файл path.
// Если файл существует, он перезаписывается.
func SaveJSON(path string, tasks []todo.Task) error {
	jsonBytes, errSerialization := json.MarshalIndent(tasks, "", "  ")
	if errSerialization != nil {
		return fmt.Errorf("сериализации: %w", errSerialization)
	}

	if errWrite := os.WriteFile(path, jsonBytes, 0644); errWrite != nil {
		return fmt.Errorf("записи в файл: %w", errWrite)
	}

	return nil
}
// LoadJSON читает задачи из JSON-файла.
// Если файла нет, создаёт пустой файл и возвращает пустой срез.
func LoadJSON(path string) ([]todo.Task, error) {
	_, errStat := os.Stat(path)
	if errStat != nil {
		if errors.Is(errStat, os.ErrNotExist) {
			if path == "tasks.json" {
				if errSave := SaveJSON(path, []todo.Task{}); errSave != nil {
					return nil, fmt.Errorf("создания файла %s: %w", path, errSave)
				}
			} else {
				return nil, fmt.Errorf("файл %s не найден: %w", path, errStat)
			}
		} else {
			return nil, fmt.Errorf("проверки файла %s: %w", path, errStat)
		}
	}

	data, errRead := os.ReadFile(path)
	if errRead != nil {
		return nil, fmt.Errorf("чтения файла %s: %w", path, errRead)
	}

	tasks := []todo.Task{}
	if len(data) == 0 {
		return tasks, nil
	}

	if errParse := json.Unmarshal(data, &tasks); errParse != nil {
		return nil, fmt.Errorf("десериализации файла %s: %w", path, errParse)
	}

	return tasks, nil
}
