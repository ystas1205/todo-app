package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"todo-app/internal/todo"
)



// SaveCSV преобразует список задач в String и записывает в файл.
func SaveCSV(path string, tasks []todo.Task) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	taskStr := [][]string{{"ID", "Description", "Done"}}
	for _, t := range tasks {
		taskStr = append(taskStr, []string{strconv.Itoa(t.ID), t.Description, strconv.FormatBool(t.Done)})
	}

	for _, record := range taskStr {
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("записи %w", err)
		}

		writer.Flush()
	}
	return nil

}


// LoadCSV читает задачи из CSV-файла.
func LoadCSV(path string) ([]todo.Task, error) {

	tasks := []todo.Task{}

	file, errOpen := os.Open(path)

	if errOpen != nil {
		return tasks, fmt.Errorf("открытия: %w", errOpen)

	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, errRead := reader.ReadAll()
	if errRead != nil {
		return tasks, fmt.Errorf("чтения %w", errRead)
	}
	for index, record := range records {
		if index == 0 {
			continue
		} else {
			id, errAtoi := strconv.Atoi(record[0])
			if errAtoi != nil {
				fmt.Println("ошибка", errAtoi)

			}
			b, errParse := strconv.ParseBool(record[2])

			if errParse != nil {
				fmt.Println("ошибка", errParse)

			}

			tasks = append(tasks, todo.Task{ID: id, Description: record[1], Done: b})
		}

	}

	return tasks, nil

}
