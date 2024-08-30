package main

import "fmt"

func Listing(ptr *[]Task) {
	for num := 1; num != 0; {
		fmt.Println(
			"Task listing :\n",
			"0. Back\n",
			"1. All tasks\n",
			"2. Todo\n",
			"3. In progress\n",
			"4. Done",
		)
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			for _, task := range *ptr {
				fmt.Println(task)
			}
		case 2:
			for _, task := range *ptr {
				if task.Status == "-todo-" {
					fmt.Println(task)
				}
			}
		case 3:
			for _, task := range *ptr {
				if task.Status == "..in_progress.." {
					fmt.Println(task)
				}
			}
		case 4:
			for _, task := range *ptr {
				if task.Status == "!!done!!" {
					fmt.Println(task)
				}
			}
		default:
			return
		}
	}
}
