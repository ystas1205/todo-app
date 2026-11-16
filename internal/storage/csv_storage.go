package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"todo-app/internal/todo"
)

// SaveCSV преобразует список задач в String и записывает в файл.
func SaveCSV(path string, tasks []todo.Task) (err error) {

	file, errCreate := os.Create(path)
	if errCreate != nil {
		return fmt.Errorf("создания %s,%w", path, errCreate)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("ошибка закрытия файла %s: %w", path, cerr)
		}
	}()

	writer := csv.NewWriter(file)

	taskStr := [][]string{{"ID", "Description", "Done"}}
	for _, t := range tasks {
		taskStr = append(taskStr, []string{strconv.Itoa(t.ID), t.Description, strconv.FormatBool(t.Done)})
	}

	for _, record := range taskStr {
		errWriter := writer.Write(record)
		if errWriter != nil {
			return fmt.Errorf("записи %s,%w", path, errWriter)
		}

	}
	writer.Flush()
	if err := writer.Error(); err != nil {

		return fmt.Errorf("записи CSV в %s: %w", path, err)
	}
	

	return nil

}

// LoadCSV читает задачи из CSV-файла.
func LoadCSV(path string) ([]todo.Task, error) {

	tasks := []todo.Task{}

	file, errOpen := os.Open(path)

	if errOpen != nil {
		return nil, fmt.Errorf("чтение:%s %w", path, errOpen)

	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, errRead := reader.ReadAll()

	if errRead != nil {
		return nil, fmt.Errorf("чтение %s,%w", path, errRead)
	}
	for index, record := range records {
		if len(record) != 3 {
			return nil, fmt.Errorf("в строке %d (%s) ожидается 3 поля, получено %d", index+1, path, len(record))
		}

		if index == 0 {
			if record[0] != "ID" || record[1] != "Description" || record[2] != "Done" {
				return nil, fmt.Errorf("некорректный заголовок в %s: ожидается [ID Description Done], получено %v", path, record)
			}
			continue
		} else {
			id, errAtoi := strconv.Atoi(record[0])
			if errAtoi != nil {
				return nil, fmt.Errorf("некорректный ID в строке %d (%s): %w", index+1, path, errAtoi)

			}
			b, errParse := strconv.ParseBool(record[2])

			if errParse != nil {
				return nil, fmt.Errorf("некорректный Done в строке %d (%s): %w", index+1, path, errParse)

			}

			tasks = append(tasks, todo.Task{ID: id, Description: record[1], Done: b})
		}

	}

	return tasks, nil

}
