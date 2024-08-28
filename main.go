package main

import (
	"fmt"
)

type Task struct {
	ID                  int
	Description, Status string
	Date                string
}

func listing() {
	for num := 1; num != 0; {
		fmt.Println(
			"Листинг задач :\n",
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
			continue
		case 1:
			Editing()
		case 2:
			listing()
		default:
			continue
		}
	}
}
