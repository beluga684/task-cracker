package main

import (
	"fmt"
)

type Task struct {
	Description, Status string
	Date                string
}

func main() {
	tasks := []Task{}

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
			Editing(&tasks)
		case 2:
			Listing(&tasks)
		default:
			continue
		}
	}
}
