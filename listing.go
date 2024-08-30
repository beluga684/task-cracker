package main

import "fmt"

func Listing(ptr *[]Task) {
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
			i := 1
			for _, task := range *ptr {
				fmt.Println(i, task)
				i++
			}
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
