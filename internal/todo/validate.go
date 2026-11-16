package todo

import (
	"fmt"
	"strings"
)
//  ValidateTasks проверяет срез задач:
//  -ID положительный и уникальный.
//  -Description не пустой после обрезки пробелов.
func ValidateTasks(tasks []Task) ([]Task, error) {
	seenIDs := map[int]int{}

	for index, task := range tasks {
		if task.ID <= 0 {
			return nil, fmt.Errorf("задача %d: некорректный ID %d  ожидается положительное число", index+1, task.ID)
		}
		if strings.TrimSpace(task.Description) == "" {
			return nil, fmt.Errorf("задача %d (ID %d): пустое описание  ожидается непустая строка", index+1, task.ID)
		}

		if prev, ok := seenIDs[task.ID]; ok {
			return nil, fmt.Errorf("задача %d (ID %d): дубликат ID уже встречался в задаче #%d", index+1, task.ID, prev)
		}
		seenIDs[task.ID] = index + 1
	}

	return tasks, nil
}
