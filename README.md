# todo-app

# To-Do менеджер на Go для командной строки.

## Установка

Сборка бинарника:

    go build -o todo ./cmd/todo

Запуск без сборки:

    go run ./cmd/todo <команда> [флаги]

## Команды

### add — добавить задачу

    ./todo add --desc="Купить книги"

Флаги:

- `--desc` — описание задачи (обязательно)

Вывод:

    добавлено: Купить книги

---

### list — показать список задач

    ./todo list
    ./todo list --filter=done
    ./todo list --filter=pending

Флаги:

- `--filter` — `all` (по умолчанию), `done`, `pending`

Вывод:
    [{1 Изучть GO false} {2 Сходить в спортзал false}]


---

### complete — отметить задачу выполненной

    ./todo complete --id=2

Флаги:

- `--id` — идентификатор задачи (обязательно, > 0)

Вывод:

    задача 2 отмечена выполненной

---

### delete — удалить задачу

    ./todo delete --id=3

Флаги:

- `--id` — идентификатор задачи (обязательно, > 0)

Вывод:

    задача 3 удалена

---

### export — экспортировать задачи в файл

    ./todo export --format=json --out=backup.json
    ./todo export --format=csv --out=backup.csv

Флаги:

- `--format` — `json` или `csv` (обязательно)
- `--out` — путь к файлу (обязательно, расширение должно совпадать с форматом)

Вывод:

    экспортировано в backup.json

Создаёт копию задач в указанном файле. `tasks.json` не меняется.

---

### load — импортировать задачи из файла

    ./todo load --file=backup.json
    ./todo load --file=tasks.csv

Флаги:

- `--file` — путь к файлу (обязательно, `.json` или `.csv`)

Вывод:

    загружено задач: 5

Перезаписывает `tasks.json` содержимым указанного файла.

---

## Хранение данных

Задачи хранятся в файле `tasks.json` в текущей директории.

Формат:

    [
      {
        "id": 1,
        "description": "Изучить Go",
        "done": false
      },
      {
        "id": 2,
        "description": "Купить продукты",
        "done": true
      }
    ]

При первом запуске файл создаётся автоматически (пустой массив `[]`).

## Структура проекта

    todo-app/
    ├── cmd/
    │   └── todo/
    │       └── main.go
    ├── internal/
    │   ├── todo/
    │   │   ├── task.go
    │   │   └── manager.go
    │   └── storage/
    │       ├── json_storage.go
    │       └── csv_storage.go
    ├── go.mod
    └── README.md

## Примеры использования

    ./todo add --desc="Купить книги"
    ./todo add --desc="Позвонить маме"
    ./todo add --desc="Сделать зарядку"

    ./todo list

    ./todo complete --id=2

    ./todo list --filter=done

    ./todo delete --id=3

    ./todo export --format=csv --out=tasks.csv

    ./todo load --file=backup.json