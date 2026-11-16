package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("использование: todo команда флаг. Пример ./todo add --desc=Купить книги")
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	// Вызов функции чтения файла
	tasksLoad, errLoad := storage.LoadJSON("tasks.json")
	if errLoad != nil {
		fmt.Println("ошибка", errLoad)
		os.Exit(1)
	}

	switch cmd {

	case "add":
		// разбор флагов команды add
		setFlags := flag.NewFlagSet("add", flag.ExitOnError)
		desc := setFlags.String("desc", "", "описание")
		setFlags.Parse(args)
		// проверка обязательного флага
		if strings.TrimSpace(*desc) == "" {
			fmt.Println("обязательный флаг --desc не указан")
			os.Exit(1)
		}
		// добавление задачи и сохранение в файл
		taskAdd := todo.Add(tasksLoad, *desc)
		if errSave := storage.SaveJSON("tasks.json", taskAdd); errSave != nil {
			fmt.Println("ошибка", errSave)
			os.Exit(1)
		}
		fmt.Println("добавлено:", *desc)

	case "list":
		// разбор флагов команды list
		setFlags := flag.NewFlagSet("list", flag.ExitOnError)
		filter := setFlags.String("filter", "all", "фильтрация")
		setFlags.Parse(args)

		// проверка, что значение filter входит в допустимые:
		// "all", "done", "pending"
		if strings.TrimSpace(*filter) != "all" && strings.TrimSpace(*filter) != "done" && strings.TrimSpace(*filter) != "pending" {
			fmt.Printf("неподдерживаемый фильтр: %q\n", *filter)
			fmt.Println("доступные значения: all, done, pending")
			os.Exit(1)
		}
		// фильтрация задач
		filtered := todo.List(tasksLoad, *filter)
		fmt.Println(filtered)

	case "complete":
		// разбор флагов команды complete
		setFlags := flag.NewFlagSet("complete", flag.ExitOnError)
		id := setFlags.Int("id", 0, "идентификатор")
		setFlags.Parse(args)
		// проверка на корректность ID
		if *id <= 0 {
			fmt.Println("некорректный --id: должен быть положительным числом")
			os.Exit(1)

		}
		// отметка задачи выполненной
		taskComplete, errComplete := todo.Complete(tasksLoad, *id)
		if errComplete != nil {
			fmt.Println("не удалось отметить задачу:", errComplete)
			os.Exit(1)
		}
		if errSave := storage.SaveJSON("tasks.json", taskComplete); errSave != nil {
			fmt.Println("ошибка", errSave)
			os.Exit(1)
		}
		fmt.Printf("задача %d отмечена выполненой\n", *id)

	case "delete":
		// разбор флагов команды delete
		setFlags := flag.NewFlagSet("delete", flag.ExitOnError)
		id := setFlags.Int("id", 0, "идентификатор")
		setFlags.Parse(args)
		// проверка на корректность ID
		if *id <= 0 {
			fmt.Println("некорректный --id: должен быть положительным числом")
			os.Exit(1)
		}
		// удаление задачи
		taskDelete, errDelete := todo.Delete(tasksLoad, *id)
		if errDelete != nil {
			fmt.Println("не удалось удалить задачу:", errDelete)
			os.Exit(1)
		}
		if errSave := storage.SaveJSON("tasks.json", taskDelete); errSave != nil {
			fmt.Println("ошибка", errSave)
			os.Exit(1)
		}
		fmt.Printf("задача %d удалена\n", *id)
	case "export":
		// разбор флагов команды export
		setFlags := flag.NewFlagSet("export", flag.ExitOnError)
		format := setFlags.String("format", "", "формат файла")
		out := setFlags.String("out", "", "путь к файлу")
		setFlags.Parse(args)
		// проверка, что обязательные флаги переданы
		if strings.TrimSpace(*format) == "" {
			fmt.Println("обязательный флаг --format не указан")
			os.Exit(1)
		}
		if strings.TrimSpace(*out) == "" {
			fmt.Println("обязательный флаг --out не указан")
			os.Exit(1)
		}
		// извлекает расширение файла и проверяет,что поддерживается
		extensionOut := filepath.Ext(*out)
		if extensionOut != ".csv" && extensionOut != ".json" {
			fmt.Printf("неподдерживаемое расширение %q: ожидается .json или .csv\n", extensionOut)
			os.Exit(1)
		}
		extensionFormat := filepath.Ext(*format)
		if extensionOut != extensionFormat {
			fmt.Printf("формат %s не соответствует расширению %s\n", *format, extensionOut)
			os.Exit(1)
		}
		switch *format {
		case ".csv":
			// экспортирует файл в CSV
			if errSave := storage.SaveCSV(*out, tasksLoad); errSave != nil {
				fmt.Println("ошибка", errSave)
				os.Exit(1)
			}
			// fmt.Printf("Файл задач успешно экспортирован в файл %s",*out)
		case ".json":
			// экспортирует файл в JSON
			if errSave := storage.SaveJSON(*out, tasksLoad); errSave != nil {
				fmt.Println("ошибка", errSave)
				os.Exit(1)
			}
			// fmt.Printf("Файл задач успешно экспортирован в файл %s",*out)
		default:
			fmt.Println("неизвестная команда:", extensionFormat)
			fmt.Println("доступные команды: .json или .csv")
			os.Exit(1)
		}
		fmt.Printf("экспортировано в %s\n", *out)
	case "load":
		// разбор флагов команды load
		setFlags := flag.NewFlagSet("load", flag.ExitOnError)
		file := setFlags.String("file", "", "путь к файлу")
		setFlags.Parse(args)
		// проверка, что обязательные флаги переданы
		if strings.TrimSpace(*file) == "" {
			fmt.Println("обязательный флаг --file не указан")
			os.Exit(1)
		}
		// извлекает расширение файла и проверяет,что поддерживается
		extensionFile := filepath.Ext(*file)
		if extensionFile != ".csv" && extensionFile != ".json" {
			fmt.Println("неподдерживаемый формат файла: ожидается .json или .csv")
			os.Exit(1)
		}
		switch extensionFile {
		case ".csv":
			// чтение файла CSV
			tasks, errLoad := storage.LoadCSV(*file)
			if errLoad != nil {
				fmt.Println("ошибка", errLoad)
				os.Exit(1)
			}
			//валидация задач
			tasksValidate, errValidate := todo.ValidateTasks(tasks)
			if errValidate != nil {
				fmt.Println("ошибка", errValidate)
				os.Exit(1)
			}
			if errSave := storage.SaveJSON("tasks.json", tasksValidate); errSave != nil {
				fmt.Println("ошибка", errSave)
				os.Exit(1)
			}
			fmt.Printf("Импортировано задач: %d (из %s)\n", len(tasksValidate), *file)
		case ".json":
			// импорт файла в JSON
			tasks, errLoad := storage.LoadJSON(*file)
			if errLoad != nil {
				fmt.Println("ошибка", errLoad)
				os.Exit(1)
			}
			// валидация импортируемых задач перед сохранением в основной файл
			tasksValidate, errValidate := todo.ValidateTasks(tasks)
			if errValidate != nil {
				fmt.Println("ошибка", errValidate)
				os.Exit(1)
			}
			if errSave := storage.SaveJSON("tasks.json", tasksValidate); errSave != nil {
				fmt.Println("ошибка", errSave)
				os.Exit(1)
			}
			fmt.Printf("Импортировано задач: %d (из %s)\n", len(tasksValidate), *file)
		}
	default:
		fmt.Println("неизвестная команда:", cmd)
		fmt.Println("доступные команды: add, list, complete, delete, export, load")
		os.Exit(1)
	}
}
