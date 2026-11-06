package todo

import (
	"fmt"
	"slices"
)


// Add добавляет задачу с описанием desc в список tasks
// и возвращает обновлённый список.
// ID новой задачи вычисляется как максимальный ID среди существующих + 1.
func Add(tasks []Task, desc string) []Task {

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	newID := maxID + 1

	task := Task{ID: newID, Description: desc, Done: false}
	tasks = append(tasks, task)
	return tasks

}

// List возвращает задачи, отфильтрованные по filter.
//
// Допустимые значения filter:
//   "all"     — все задачи
//   "done"    — только выполненные (Done == true)
//   "pending" — только невыполненные (Done == false)
//
// Неизвестный filter обрабатывается как "all".
func List(tasks []Task, filter string) []Task {

	filterd := []Task{}

	switch filter {
	case "all":
		return tasks

	case "done":
		for _, t := range tasks {
			if t.Done == true {
				filterd = append(filterd, t)
			}

		}
	case "pending":
		for _, t := range tasks {
			if t.Done == false {
				filterd = append(filterd, t)
			}
		}
	}
	return filterd

}
// Complete отмечает задачу с ID как выполненную (Done = true).
// Возвращает обновлённый список задач.
// Возвращает ошибку, если задача с таким ID не найдена
func Complete(tasks []Task, id int) ([]Task, error) {

	for index, t := range tasks {
		if t.ID == id {
			tasks[index].Done = true
			return tasks, nil

		}
	}
	return tasks, fmt.Errorf("нет такого ID %d", id)

}

// Delete удаляет задачу по ID .
// Возвращает обновлённый список задач.
// Возвращает ошибку, если задача с таким ID не найдена
func Delete(tasks []Task, id int) ([]Task, error) {

	taskCintains := slices.ContainsFunc(tasks, func(t Task) bool {
		return id == t.ID
	})
	if taskCintains != true {
		return tasks, fmt.Errorf("нет такого ID %d", id)
	}

	tasksDelete := slices.DeleteFunc(tasks, func(t Task) bool {
		return id == t.ID
	})
	return tasksDelete, nil

}
