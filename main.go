package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID                  int
	Description, Status string
	Date                time.Time
}

func exit() {
	return
}

func editing() {
	for num := 1; num != 0; {
		fmt.Println(
			"Выберите действие :\n",
			"0. Назад\n",
			"1. Новая задача\n",
			"2. Изменить задачу\n",
			"3. Удалить задачу")
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			// создать
		case 2:
			// изменение название-статус
		case 3:
			// удалить
		default:
			return
		}
	}
}

func listing() {
	for num := 1; num != 0; {
		fmt.Println(
			"Выберите действие :\n",
			"0. Назад\n",
			"1. Все задачи\n",
			"2. Ожидающие\n",
			"3. В процессе\n",
			"4. Завершённые",
		)
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			// all
		case 2:
			// todo
		case 3:
			// in progress
		case 4:
			// done
		default:
			return
		}
	}
}

func main() {
	//подгрузка файла
	//создание этого файла, если не найден
	for num := 1; num != 0; {
		fmt.Println(
			"Выберите действие :\n",
			"0. Выход из программы\n",
			"1. Работа с задачами\n",
			"2. Листинг задач",
		)
		fmt.Scan(&num)

		switch num {
		case 0:
			exit()
		case 1:
			editing()
		case 2:
			listing()
		default:
			continue
		}
	}
}
