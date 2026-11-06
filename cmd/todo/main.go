package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("использование: todo команда флаг.Пример ./todo --desc add=Купить книги")
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	// Вызов функции чтения файла
	tasksLoad, err := storage.LoadJSON("tasks.json")
	if err != nil {
		fmt.Println("ошибка", err)
	}

	switch cmd {

	case "add":
		// разбор флагов команды add
		setFlags := flag.NewFlagSet("add", flag.ExitOnError)
		desc := setFlags.String("desc", "", "описание")
		setFlags.Parse(args)
		// проверка обязательного флага
		if *desc == "" {
			fmt.Println("не указан --desc")
			os.Exit(1)

		}
		// добавление задачи и сохранение в файл
		taskADd := todo.Add(tasksLoad, *desc)
		if err := storage.SaveJSON("tasks.json", taskADd); err != nil {
			fmt.Println("ошибка", err)
		}
		fmt.Println("добавлено:", *desc)

	case "list":
		// разбор флагов команды lsit
		setFlags := flag.NewFlagSet("list", flag.ExitOnError)
		filter := setFlags.String("filter", "all", "фильтрация")
		setFlags.Parse(args)
		// фильтрация задач
		filtered := todo.List(tasksLoad, *filter)
		fmt.Println(filtered)

	case "complete":
		// разбор флагов команды list
		setFalgs := flag.NewFlagSet("complete", flag.ExitOnError)
		id := setFalgs.Int("id", 0, "идентификатор")
		setFalgs.Parse(args)
		// проверка на коректность ID
		if *id == 0 {
			fmt.Println("некорректный --id: должен быть положительным числом")
			os.Exit(1)

		}
		// отметка задачи вылненой
		taskComplete, errComplete := todo.Complete(tasksLoad, *id)
		if errComplete != nil {
			fmt.Println("не удалось отметить задачу:", errComplete)
		}

		if err := storage.SaveJSON("tasks.json", taskComplete); err != nil {
			fmt.Println("ошибка", err)
		}
		fmt.Printf("задача %d отмечена выполненой\n", *id)

	case "delete":
		// разбор флагов команды delete
		setFlags := flag.NewFlagSet("delete", flag.ExitOnError)
		id := setFlags.Int("id", 0, "идентификатор")
		setFlags.Parse(args)
		// проверка на коректность ID
		if *id == 0 {
			fmt.Println("некорректный --id: должен быть положительным числом")
			os.Exit(1)

		}
		// удаление задачи
		taskDelete, errDelete := todo.Delete(tasksLoad, *id)
		if errDelete != nil {
			fmt.Println("не удалось удалить задчу:", errDelete)
		}

		if err := storage.SaveJSON("tasks.json", taskDelete); err != nil {
			fmt.Println("ошибка", err)

		}
		fmt.Printf("задача %d удалена\n", *id)

	case "export":
		// разбор флагов команды export
		setFlags := flag.NewFlagSet("export", flag.ExitOnError)
		format := setFlags.String("format", "", "формат файла")
		out := setFlags.String("out", "", "путь к файлу")
		setFlags.Parse(args)

		// извлекает расширение файла и проверяет, что оно поддерживается
		parseOut := strings.Split(*out, ".")
		if parseOut[1] != "csv" && parseOut[1] != "json" {
			fmt.Printf("недопустимое расширение файла %s должно быть csv или json\n", parseOut[1])
			os.Exit(1)
		}

		if parseOut[1] != *format {
			fmt.Printf("формат %s не соответствует расширению %s\n", *format, parseOut[1])
			os.Exit(1)
		}

		switch *format {
		case "csv":
			// экспортирует файл в CSV
			if err := storage.SaveCSV(*out, tasksLoad); err != nil {
				fmt.Println("ошибка", err)
			}

		case "json":
			// экспортирует файл в JSON
			if err := storage.SaveJSON(*out, tasksLoad); err != nil {
				fmt.Println("ошибка", err)
			}
		default:
			fmt.Println("неизвестная команда:", *format)
			fmt.Println("доступные команды: json или csv")
			os.Exit(1)
		}

	case "load":
		// разбор флагов команды load
		setFlags := flag.NewFlagSet("load", flag.ExitOnError)
		file := setFlags.String("file", "", "путь к файлу")
		setFlags.Parse(args)

		// извлекает расширение файла и проверяет, что оно поддерживается
		parseNameFile := strings.Split(*file, ".")
		if parseNameFile[1] != "csv" && parseNameFile[1] != "json" {
			fmt.Println("неподдерживаемый формат файла: ожидается .json или .csv")
			os.Exit(1)
		}
		switch parseNameFile[1] {
		case "csv":
			// ипорт файла в CSV
			tasks, err := storage.LoadCSV(*file)
			if err != nil {
				fmt.Println("ошибка", err)
			}

			if err := storage.SaveJSON("tasks.json", tasks); err != nil {
				fmt.Println("ошибка", err)
			}

		case "json":
			// импорт файла в JSON
			tasks, err := storage.LoadJSON(*file)
			if err != nil {
				fmt.Println("ошибка", err)
			}
			if err := storage.SaveJSON("tasks.json", tasks); err != nil {
				fmt.Println("ошибка", err)
			}
		}
	default:
		fmt.Println("неизвестная команда:", cmd)
		fmt.Println("доступные команды: add, list, complete, delete, export, load")
		os.Exit(1)
	}
}
